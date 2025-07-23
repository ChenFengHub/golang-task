package middle

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ChenfengHub/golang-task/task04/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

// 全局db
var db *gorm.DB

func InitDB(inputDB *gorm.DB) {
	db = inputDB
}

// func CORSMiddleware() gin.HandlerFunc {
// 	return cors.New(cors.Config{
// 		AllowOrigins:     []string{"https://prod.com", "*"},
// 		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
// 		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
// 		ExposeHeaders:    []string{"Content-Length"},
// 		AllowCredentials: true,
// 		MaxAge:           12 * time.Hour,
// 	})
// }

func JWTAuth() gin.HandlerFunc {
	// # 1. 清理可能损坏的缓存(解决缓存问题)
	// go clean -modcache
	// # 2. 获取最新版本的 jwt 包
	// go get -u github.com/golang-jwt/jwt/v5
	// # 3. 同步依赖关系（解决依赖声明不一致问题）
	// go mod tidy
	return func(c *gin.Context) {
		// 过滤掉无需验证接口
		skipPaths := map[string]bool{
			"/v1/user/register": true,
			"/v1/user/login":    true,
		}
		// 检查当前路径是否需要跳过
		if _, exists := skipPaths[c.Request.URL.Path]; exists {
			c.Next() // 继续后续处理
			return
		}

		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "lack of token"})
			return
		}

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			secret := "b3f8d7a2e5c1f9b0a4d6c8e3b7f2a1d5e0c9b8a7d4f3e6c2a9b8d5f1e0a3c7b6d9"
			if os.Getenv("JWT_SECRET") != "" {
				secret = os.Getenv("JWT_SECRET")
			}
			return []byte(secret), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			hUserId := c.GetHeader("Userid")
			if hUserId == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Header lack of Userid"})
				return
			}
			if hUserId != claims["userId"] {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is invalid"})
				return
			}
			jwtTime, ok := claims["exp"].(float64)

			if !ok || time.Now().After(time.Unix(int64(jwtTime), 0)) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is expired"})
				return
			}

			c.Set("userId", claims["userId"])
			c.Set("roles", claims["roles"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		}
	}
}

func GenerateToken(userId string, roles []string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"roles":  roles,
		"exp":    time.Now().Add(8 * time.Hour).Unix(),
	})
	secret := "b3f8d7a2e5c1f9b0a4d6c8e3b7f2a1d5e0c9b8a7d4f3e6c2a9b8d5f1e0a3c7b6d9"
	if os.Getenv("JWT_SECRET") != "" {
		secret = os.Getenv("JWT_SECRET")
	}
	return token.SignedString([]byte(secret))
}

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *responseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// 错误处理中间件（带数据库存储）
func ErrorToDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a response writer wrapper
		w := &responseWriter{ResponseWriter: c.Writer, body: &bytes.Buffer{}}
		c.Writer = w

		// 继续处理请求链
		c.Next()
		statusCode := c.Writer.Status()
		if statusCode == 200 {
			return
		}

		// 记录错误到数据库
		if db != nil {
			logEntry := entity.Log{
				Method:     c.Request.Method,
				Path:       c.Request.URL.Path,
				StatusCode: statusCode,
				ClientIP:   c.ClientIP(),
				UserAgent:  c.Request.UserAgent(),
			}

			responseBody := w.body.String()
			// 限制错误消息长度
			if len(responseBody) > 1000 {
				logEntry.ErrMsg = string(responseBody[:1000]) + "..."
			} else {
				logEntry.ErrMsg = string(responseBody)
			}

			// 尝试获取请求ID（如果使用了请求ID中间件）
			if requestID, exists := c.Get("X-Request-ID"); exists {
				logEntry.RequestID = fmt.Sprintf("%v", requestID)
			}

			// 异步写入数据库
			go func(entry entity.Log) {
				if result := db.Create(&entry); result.Error != nil {
					//
				}
			}(logEntry)
		}
	}
}

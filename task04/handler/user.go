package hv1

import (
	"net/http"
	"strconv"

	"github.com/ChenfengHub/golang-task/task04/entity"
	"github.com/ChenfengHub/golang-task/task04/middle"
	"github.com/ChenfengHub/golang-task/task04/service"
	"github.com/ChenfengHub/golang-task/task04/store"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService *service.UserService
}

func newUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{userService: svc}
}

func SetupUserRoutes(r *gin.Engine, db *gorm.DB) {
	us := service.NewUserService(store.NewUserStore(db))
	uh := newUserHandler(us)

	api := r.Group("/v1")

	userGroup := api.Group("/user")
	// 代码块中批量操作
	{
		userGroup.POST("/register", uh.RegisterUser)
		userGroup.POST("/login", uh.Login)
	}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := h.userService.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	userId := strconv.FormatUint(uint64(user.ID), 10)
	authorization, _ := middle.GenerateToken(userId, []string{})
	c.Header("UserId", userId)
	c.Header("Authorization", authorization)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.userService.Login(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userId := strconv.FormatUint(uint64(user.ID), 10)
	authorization, _ := middle.GenerateToken(userId, []string{})
	c.Header("UserId", userId)
	c.Header("Authorization", authorization)
	c.JSON(http.StatusCreated, gin.H{"message": "User login successfully"})
}

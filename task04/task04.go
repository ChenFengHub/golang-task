package main

import (
	"fmt"
	"os"

	"github.com/ChenfengHub/golang-task/task04/entity"
	hv1 "github.com/ChenfengHub/golang-task/task04/handler"
	"github.com/ChenfengHub/golang-task/task04/middle"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 1.初始化DB
	username := "root"
	if os.Getenv("db_username") != "" {
		username = os.Getenv("db_username")
	}
	password := "123456"
	if os.Getenv("db_password") != "" {
		password = os.Getenv("db_password")
	}
	port := "3306"
	if os.Getenv("db_port") != "" {
		port = os.Getenv("db_port")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/gorm?charset=utf8mb4&parseTime=True&loc=Local", username, password, port)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entity.User{}, &entity.Post{}, &entity.Comment{}, &entity.Log{})

	// 2. 初始化路由
	router := gin.Default()
	middle.InitDB(db)
	router.Use(middle.ErrorToDB(), middle.JWTAuth())
	hv1.SetupUserRoutes(router, db)
	hv1.SetupPostRoutes(router, db)
	hv1.SetupCommentRoutes(router, db)

	// 3. 启动web服务
	router.Run(":8888")
}

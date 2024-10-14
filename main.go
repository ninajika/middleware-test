package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ninajika/middleware-test/handler"
)

func main() {
	route := gin.Default()
	// route.Use(middleware.AuthMiddleware())
	route.POST("/login", handler.LoginHandler)
	route.Run(":8080")
}

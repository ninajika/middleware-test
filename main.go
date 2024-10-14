package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ninajika/middleware-test/handler"
	"github.com/ninajika/middleware-test/middleware"
)

func main() {
	route := gin.Default()
	route.POST("/login", handler.LoginHandler)
	// Grup rute yang dilindungi
	protected := route.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/protected-route", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Selamat anda telah login di rute protected!"})
	})
	route.Run(":8080")
}

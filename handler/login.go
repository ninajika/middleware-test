package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var hardcodedUsername = "admin"
var hardcodedPassword = "password"

func LoginHandler(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if credentials.Username == hardcodedUsername && credentials.Password == hardcodedPassword {
		c.JSON(http.StatusOK, gin.H{"token": "Jadis1ngA"})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	}
}

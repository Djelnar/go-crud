package main

import "github.com/gin-gonic/gin"

// Logout controller
func Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		`message`: `Logout`,
	})
}

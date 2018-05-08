package main

import "github.com/gin-gonic/gin"

// Login controller
func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		`message`: `Login`,
	})
}

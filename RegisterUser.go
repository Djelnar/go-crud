package main

import (
	"github.com/gin-gonic/gin"
)

type tregisterC struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterUser controller
func RegisterUser(c *gin.Context) {
	json := tregisterC{}

	err := c.ShouldBindJSON(&json)

	if err == nil {
		usernameExist := Client.Cmd(`GET`, string(json.Username))
		uStr, _ := usernameExist.Str()
		if len(uStr) == 0 {
			Client.Cmd(`SET`, string(json.Username), string(json.Password))
			Client.Cmd(`SADD`, string(json.Username)+`_roles`, `user`)
			c.Status(200)
		} else {
			c.JSON(400, gin.H{
				`error`: `username ` + json.Username + ` already taken`,
			})
		}
	}

}

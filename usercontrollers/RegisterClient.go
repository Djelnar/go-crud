package usercontrollers

import (
	"github.com/gin-gonic/gin"
	"github.com/mediocregopher/radix.v2/redis"
)

type tregisterC struct {
	Username string `json:"username" binding:"required,alphanum,min=2"`
	Password string `json:"password" binding:"required,min=8"`
}

// RegisterClient controller
func RegisterClient(Client *redis.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		json := tregisterC{}

		err := c.ShouldBindJSON(&json)

		if err != nil {
			c.JSON(400, gin.H{
				`error`: `u dun goofed`,
			})
		} else {
			usernameExist := Client.Cmd(`GET`, string(json.Username))
			uStr, _ := usernameExist.Str()
			if len(uStr) != 0 {
				c.JSON(400, gin.H{
					`error`: `username ` + json.Username + ` already taken`,
				})
			} else {
				Client.Cmd(`SET`, string(json.Username), string(json.Password))
				Client.Cmd(`SADD`, string(json.Username)+`_roles`, `user`)
				c.JSON(200, `OK`)
			}
		}

	}
}

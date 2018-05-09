package usercontrollers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mediocregopher/radix.v2/redis"
)

type tregisterA struct {
	Username string `json:"username" binding:"required,alphanum,min=2"`
	Password string `json:"password" binding:"required,min=8"`
	Secret   string `json:"secret" binding:"required"`
}

// RegisterAdmin controller
func RegisterAdmin(Client *redis.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		json := tregisterA{}

		err := c.ShouldBindJSON(&json)

		secret := os.Getenv(`ADMINSECRET`)

		if err == nil && json.Secret == secret {
			usernameExist := Client.Cmd(`GET`, string(json.Username))
			uStr, _ := usernameExist.Str()
			if len(uStr) == 0 {
				Client.Cmd(`SET`, string(json.Username), string(json.Password))
				Client.Cmd(`SADD`, string(json.Username)+`_roles`, `user`)
				Client.Cmd(`SADD`, string(json.Username)+`_roles`, `admin`)
				c.JSON(200, `OK`)
			} else {
				c.JSON(400, gin.H{
					`error`: `username ` + json.Username + ` already taken`,
				})
			}
		} else if err == nil && json.Secret != secret {
			c.JSON(403, gin.H{
				`error`: `ur not allowd`,
			})
		} else {
			c.JSON(400, gin.H{
				`error`: `u dun goofed`,
			})
		}

	}
}

package main

import (
	"github.com/gin-gonic/gin"
)

type tregisterA struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
}

// RegisterAdmin controller
func RegisterAdmin(c *gin.Context) {
	json := tregisterA{}

	err := c.ShouldBindJSON(&json)

	if err == nil && json.Secret == `cyka` {
		usernameExist := Client.Cmd(`GET`, string(json.Username))
		uStr, _ := usernameExist.Str()
		if len(uStr) == 0 {
			Client.Cmd(`SET`, string(json.Username), string(json.Password))
			Client.Cmd(`SADD`, string(json.Username)+`_roles`, `user`)
			Client.Cmd(`SADD`, string(json.Username)+`_roles`, `admin`)
			c.Status(200)
		} else {
			c.JSON(400, gin.H{
				`error`: `username ` + json.Username + ` already taken`,
			})
		}
	} else if json.Secret != `cyka` {
		c.JSON(400, gin.H{
			`error`: `ur not allowed`,
		})
	}

}

package usercontrollers

import (
	"crypto/rand"
	"encoding/hex"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mediocregopher/radix.v2/redis"
)

type tlogin struct {
	Username string `json:"username" binding:"required,alphanum,min=2"`
	Password string `json:"password" binding:"required,min=8"`
}

// Login controller
func Login(Client *redis.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		json := tlogin{}

		err := c.ShouldBindJSON(&json)

		if err == nil {
			usernameExist := Client.Cmd(`GET`, string(json.Username))
			uStr, _ := usernameExist.Str()
			log.Println(uStr)
			if len(uStr) == 0 || uStr != json.Password {
				c.JSON(404, gin.H{
					`message`: `Unknown boi`,
				})
			} else {
				b := make([]byte, 16)
				rand.Read(b)
				token := hex.EncodeToString(b)
				Client.Cmd(`SADD`, string(json.Username)+`_tokens`, token)
				c.JSON(200, gin.H{
					`token`: token,
				})
			}
		} else {
			c.JSON(400, gin.H{
				`error`: `u dun goofed`,
			})
		}
	}
}

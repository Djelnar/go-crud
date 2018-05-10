package usercontrollers

import (
	"github.com/gin-gonic/gin"
	"github.com/mediocregopher/radix.v2/redis"
)

type tlogout struct {
	Username string `json:"username" binding:"required,alphanum,min=2"`
	Token    string `json:"token" binding:"required,len=32"`
}

// Logout controller
func Logout(Client *redis.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		json := tlogout{}
		err := c.ShouldBindJSON(&json)

		if err != nil {
			c.JSON(400, gin.H{
				`error`: `u dun goofed`,
			})
		} else {
			tokenExists := Client.Cmd(`SREM`, json.Username+`_tokens`, json.Token)
			v, e := tokenExists.Int64()
			if e != nil || v == 0 {
				c.JSON(403, gin.H{
					`error`: `u dun goofed`,
				})
			} else {
				c.JSON(200, gin.H{
					`message`: v,
				})
			}
		}
	}
}

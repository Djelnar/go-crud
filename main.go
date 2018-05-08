package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mediocregopher/radix.v2/redis"
)

// Client redis client
var Client *redis.Client

func main() {
	Client, _ = redis.Dial(`tcp`, `localhost:6379`)

	r := gin.Default()

	user := r.Group(`/user`)
	{
		user.POST(`/register-client`, RegisterUser)
		user.POST(`/register-admin`, RegisterAdmin)
		user.POST(`/login`, Login)
		user.POST(`/logout`, Logout)
	}

	r.Run(`:3000`)
}

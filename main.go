package main

import (
	"go-crud/usercontrollers"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mediocregopher/radix.v2/redis"
)

// Client redis client
var Client *redis.Client

func main() {
	host := os.Getenv(`REDISHOST`)

	port := os.Getenv(`REDISPORT`)

	Client, _ = redis.Dial(`tcp`, host+`:`+port)

	r := gin.Default()

	user := r.Group(`/user`)
	{
		user.POST(`/register-client`, usercontrollers.RegisterClient(Client))
		user.POST(`/register-admin`, usercontrollers.RegisterAdmin(Client))
		user.POST(`/login`, usercontrollers.Login(Client))
		user.POST(`/logout`, usercontrollers.Logout(Client))
	}

	r.Run(`:3000`)
}

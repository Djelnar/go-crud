package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mediocregopher/radix.v2/redis"
)

var redisClient *redis.Client

type registerData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == `POST` {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var t registerData
		err = json.Unmarshal(body, &t)
		if err != nil {
			panic(err)
		}

		usernameExist, _ := redisClient.Cmd(`GET`, string(t.Username)).Str()

		if len(usernameExist) == 0 {
			w.Header().Set(`Content-Type`, `application/json`)

			redisClient.Cmd(`SET`, string(t.Username), string(t.Password))

			w.WriteHeader(200)
			fmt.Fprintln(w, `OK`)
		} else {
			w.Header().Set(`Content-Type`, `application/json`)
			w.WriteHeader(404)
			fmt.Fprintln(w, `no lol`)
		}
	}

}

func main() {
	log.Printf(`Listening on %s`, `3333`)

	client, _ := redis.Dial("tcp", "localhost:6379")

	redisClient = client

	http.HandleFunc(`/register`, register)
	http.ListenAndServe(`:3333`, nil)
}

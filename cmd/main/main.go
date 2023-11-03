package main

import (
	"go-chat-app/api"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY")) //見つからないときは空を返すので注意
)

func main() {
	if jwtSecret == nil {
		log.Printf("JWT_SECRET_KEY is not set")
		return
	}

	r := gin.Default()
	api.SetupRoutes(r)
	r.Run(":8000")
}

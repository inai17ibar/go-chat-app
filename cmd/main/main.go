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
	// 実際のデプロイ時にはHTTPSを使用してサーバーを起動することをおすすめします。
	// r.RunTLS(":8000", "path_to_certfile", "path_to_keyfile")
	r.Run(":8000")
}

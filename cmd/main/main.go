package main

import (
	"go-chat-app/api"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var (
	jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY")) //見つからないときは空を返すので注意
	store     = cookie.NewStore([]byte(os.Getenv("JWT_SECRET_KEY")))
)

// カスタムエラーハンドラーミドルウェア
func customErrorHandler(c *gin.Context) {
	// エラーハンドリング
	defer func() {
		if r := recover(); r != nil {
			// エラーログを出力
			log.Println("Recovered from panic:", r)

			// エラーレスポンスを返す
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
			})
		}
	}()

	// リクエストの処理を続ける
	c.Next()
}

func main() {
	if jwtSecret == nil {
		log.Printf("JWT_SECRET_KEY is not set")
		return
	}

	r := gin.Default()

	// Panic Recovery Middlewareの使用
	r.Use(gin.Recovery())

	// カスタムエラーハンドラーミドルウェアの設定
	r.Use(customErrorHandler)

	r.Use(sessions.Sessions("my-session", store))

	api.SetupRoutes(r)
	// 実際のデプロイ時にはHTTPSを使用してサーバーを起動することをおすすめします。
	// r.RunTLS(":8000", "path_to_certfile", "path_to_keyfile")
	r.Run(":8000")
}

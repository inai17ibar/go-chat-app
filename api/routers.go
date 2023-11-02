package api

import (
	"go-chat-app/internal/auth"
	"go-chat-app/internal/chat"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST"}
	r.Use(cors.New(config))

	// Chat routes
	r.GET("/ws", func(c *gin.Context) {
		chat.HandleConnections(c.Writer, c.Request)
	})

	// Auth routes
	r.POST("/login", auth.Login)
	r.POST("/welcome", auth.Welcome)

	// Auth routes
	// ... You can define routes for login, registration, etc. here.
}

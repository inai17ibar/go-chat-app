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
	config.AllowMethods = []string{"GET", "POST", "DELETE"}
	r.Use(cors.New(config))

	// Chat routes
	r.GET("/ws", func(c *gin.Context) {
		chat.HandleConnections(c.Writer, c.Request)
	})

	// Auth routes
	r.POST("/auth/login", auth.Login)
	r.POST("/auth/register", auth.Register)
	r.DELETE("/auth/delete", auth.DeleteAccount)
}

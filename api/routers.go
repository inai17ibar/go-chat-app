package api

import (
	"fmt"
	"go-chat-app/internal/auth"
	"go-chat-app/internal/chat"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No authorization header provided"})
			return
		}

		if !strings.HasPrefix(header, BearerSchema) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			return
		}

		tokenString := header[len(BearerSchema):]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if jwt.SigningMethodHS256 != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
			return jwtSecret, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("username", claims["name"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}
	}
}

func SetupRoutes(r *gin.Engine) {
	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "Upgrade", "Connection", "Sec-WebSocket-Key", "Sec-WebSocket-Version", "Sec-WebSocket-Extensions"} // WebSocket関連のヘッダーを追加
	//config.AllowCredentials = true                                                                                                                                                         // クレデンシャルを許可
	r.Use(cors.New(config))
	//r.Use(Authenticate()) // このミドルウェアを適用するエンドポイントにのみ追加

	chat.Init()

	// Chat routes
	r.GET("/ws", func(c *gin.Context) {
		chat.HandleConnections(c.Writer, c.Request)
	})

	// Auth routes
	r.POST("/auth/login", auth.Login)
	r.POST("/auth/logout", auth.Logout)
	r.POST("/auth/register", auth.Register)
	r.DELETE("/auth/delete", auth.AuthMiddleware(), auth.DeleteAccount)
}

package middleware

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitCORS(r *gin.Engine) {
	FRONTEND_URL := os.Getenv("FRONTEND_URL")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{FRONTEND_URL},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))
}
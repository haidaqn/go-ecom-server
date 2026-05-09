package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func ApiKeyMiddleware() gin.HandlerFunc {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing X-API-KEY"})
		}
	}

	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("X-API-KEY")
		if apiKey == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing X-API-KEY"})
			return
		}

		if apiKey != os.Getenv("API_KEY") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API KEY"})
			return
		}

		ctx.Next()
	}
}

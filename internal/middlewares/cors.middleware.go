package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/haidaqn/go-ecommerce-backend-api/pkg/response"
)

func CorsMiddleware() gin.HandlerFunc {

	allowedOrigins := []string{"https://haidaqn.com", "http://localhost:3000"}

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin == "" {
			response.ErrorResponse(c, response.CodeUnauthorized, "Origin not allowed")
			c.Abort()
			return
		}
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
				c.Next()
				return
			}
		}
		response.ErrorResponse(c, response.CodeUnauthorized, "Origin not allowed")
		c.Abort()
		return
	}
}

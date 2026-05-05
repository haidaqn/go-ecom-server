package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/haidaqn/go-ecommerce-backend-api/pkg/response"
)

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.ErrorResponse(c, response.CodeUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

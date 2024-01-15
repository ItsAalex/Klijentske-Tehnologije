package middlewares

import (
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

// Add CORS headers to the HTTP response.
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Status(http.StatusNoContent)

			return
		}

		// Add CORS headers for non-preflight requests
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Continue with the next middleware/handler in the chain
		c.Next()
	}
}

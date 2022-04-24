package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	apiKeyHeader = "X-API-Key"
)

func AuthAPIKey(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.Header.Get(apiKeyHeader)

		if secret != key {
			log.Println("Key doesnt match!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		log.Println("No error during check")
		c.Next()
	}
}

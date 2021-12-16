package middleware

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/matias-ziliotto/test-golang/pkg/web"
)

func TokenAuth() gin.HandlerFunc {
	requiredToken := os.Getenv("API_TOKEN")

	if requiredToken == "" {
		log.Fatal("please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			web.Error(c, 400, "missing token")
			return
		}

		if token != os.Getenv("API_TOKEN") {
			web.Error(c, 401, "invalid token")
			return
		}

		c.Next()
	}
}

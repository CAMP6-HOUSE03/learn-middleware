package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// fungsi middleware gin
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		username := c.GetHeader("username")

		// pengecekan 1
		if username == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "for authentication, please provide username",
			})
			return
		}

		if username != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid username",
			})
			return
		}

		// melanjutkan ke handler / ke middleware selanjutnya
		c.Next()

		since := time.Since(t)

		fmt.Println("latency:", since.String())

		// nggk ada apaapa
		// logic setelah next ???

	}
}

func Auth2() gin.HandlerFunc {
	return func(c *gin.Context) {
		password := c.GetHeader("password")

		if password == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "for authentication, please provide password",
			})
			return
		}

		if password != "1234" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid password",
			})
			return
		}

		c.Next()
	}
}

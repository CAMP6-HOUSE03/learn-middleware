package middleware

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
)

// fungsi middleware gin
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1 detik

		t := time.Now()

		username := c.GetHeader("username")
		password := c.GetHeader("password")

		// pengecekan 1
		if username == "" || password == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "for authentication, please provide username and password",
			})
			return
		}

		if username != "admin" || password != "1234" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid username and password",
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

var roles = []string{"admin", "user", "guest"}

func CheckingRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetHeader("role")

		if role == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "for authorization, please provide role",
			})
			return
		}

		if !slices.Contains(roles, role) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid role",
			})
			return
		}

		// gin Keys
		c.Set("role", role)

		// mendapatkan data role dari header, setelah di check
		// nanti hasilnya dikirim ke handler

		c.Next()

	}
}

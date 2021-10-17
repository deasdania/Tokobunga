package middlewares

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities/token"
	"github.com/gin-gonic/gin"

	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

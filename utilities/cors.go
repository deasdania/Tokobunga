package utilities

import "github.com/gin-gonic/gin"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Writer.Header()
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Methods", "*")
		header.Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, Origin, Authorization")
		if c.Request.Method == "OPTIONS" {
			header.Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

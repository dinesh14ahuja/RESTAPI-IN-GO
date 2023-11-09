package middleware

import "github.com/gin-gonic/gin"

func Authenticate(c *gin.Context) {

	token := c.Request.Header.Get("token")

	if len(token) == 0 || token != "auth" {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Kindly provide token in headers",
		})
		return

	} else {
		c.Next()
	}
}

func ResponseMiddleware(c *gin.Context) {
	c.Writer.Header().Add("App", "Rest api in Go")
	c.Next()
}

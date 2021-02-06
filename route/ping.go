package route

import "github.com/gin-gonic/gin"

// Ping 心跳
func Ping(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

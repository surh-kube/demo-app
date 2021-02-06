package route

import "github.com/gin-gonic/gin"

// Index 主页
func Index(r *gin.Engine) {
	r.LoadHTMLFiles("template/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
}

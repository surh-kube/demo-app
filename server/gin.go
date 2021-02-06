package server

import (
	"bitbucket.it.ittreasury.com/gitops/demo/route"
	"github.com/gin-gonic/gin"
)

// NewGinServer 创建gin服务
func NewGinServer() *gin.Engine {
	r := gin.Default()
	route.Index(r)
	route.Ping(r)
	return r
}

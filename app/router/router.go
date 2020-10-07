package router

import (
	"github.com/gin-gonic/gin"
	"multihost-vuepress/app/pkg/mus"
)

func InitRouter() *gin.Engine {
	r := mus.Gin
	r.GET("/", func(context *gin.Context) {
		context.String(200, "%s", "webhook server")
	})
	r.NoRoute(Info)
	return r
}

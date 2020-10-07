package mus

import (
	"github.com/gin-gonic/gin"
	"github.com/mygomod/muses/pkg/logger"
	musgin "github.com/mygomod/muses/pkg/server/gin"
)

var (
	Logger *logger.Client
	Gin    *gin.Engine
)

// Init 初始化muses相关容器
func Init() error {
	Gin = musgin.Caller()
	if Gin == nil {
		panic("gin is nil")
	}
	Logger = logger.Caller("system")
	if Logger == nil {
		panic("logger is nil")
	}
	return nil
}

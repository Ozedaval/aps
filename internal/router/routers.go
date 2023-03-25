package router

import (
	"github.com/gin-gonic/gin"
	"github/hanzhang2566/asp/global/constant"
	"github/hanzhang2566/asp/internal/service"
)

func NewRouter() *gin.Engine {
	engine := gin.Default()

	group := engine.Group(constant.ApiV1)
	{
		group.POST("/:repo/pr", service.Pr)
		group.POST("/:repo/comment", service.Comment)
		group.POST("/:repo/push", service.Push)
		group.GET("ping", service.Ping)
	}

	return engine
}

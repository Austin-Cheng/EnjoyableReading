package driver

import (
	"github.com/Austin-Cheng/EnjoyableReading/adapter/driver/controllers/tags"
	"github.com/gin-gonic/gin"
)

var _ IRouter = (*Router)(nil)

type IRouter interface {
	Register(r *gin.Engine)
}

type Router struct {
	TagController *tags.Service
}

func (r *Router) Register(engine *gin.Engine) {
	router := engine.Group("/api/v1")
	{
		//标签
		tag := router.Group("/tags")
		{
			tag.POST("", r.TagController.Create) //新建
			tag.GET("", r.TagController.List)    //列表
		}
	}

}

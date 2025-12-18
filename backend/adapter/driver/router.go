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
	tagController tags.Service
}

func (r *Router) Register(engine *gin.Engine) {
	router := engine.Group("/api/v1")
	{
		//标签
		tag := router.Group("/tags")
		{
			tag.GET("", r.tagController.Create) //create
		}
	}

}

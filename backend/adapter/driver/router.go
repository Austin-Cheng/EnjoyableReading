package driver

import (
	"github.com/Austin-Cheng/EnjoyableReading/adapter/driver/controllers/papers"
	"github.com/Austin-Cheng/EnjoyableReading/adapter/driver/controllers/tags"
	"github.com/gin-gonic/gin"
)

var _ IRouter = (*Router)(nil)

type IRouter interface {
	Register(r *gin.Engine)
}

type Router struct {
	TagController  *tags.Service
	PaperController *papers.Service
}

func (r *Router) Register(engine *gin.Engine) {
	router := engine.Group("/api/v1")
	{
		//标签
		tag := router.Group("/tags")
		{
			tag.POST("", r.TagController.Create)       //新建
			tag.PUT("/:id", r.TagController.Update)    //更新
			tag.GET("/:id", r.TagController.Get)       //详情
			tag.DELETE("/:id", r.TagController.Delete) //删除
			tag.GET("", r.TagController.List)          //列表
		}

		//文章
		paper := router.Group("/papers")
		{
			paper.POST("", r.PaperController.Create)       //新建
			paper.PUT("/:id", r.PaperController.Update)    //更新
			paper.GET("/:id", r.PaperController.Get)       //详情
			paper.DELETE("/:id", r.PaperController.Delete) //删除
			paper.GET("", r.PaperController.List)          //列表
		}
	}

}

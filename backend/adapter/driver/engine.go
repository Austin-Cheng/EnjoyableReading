package driver

import (
	"github.com/Austin-Cheng/EnjoyableReading/common/settings"
	"github.com/zeromicro/go-zero/core/logx"
	"time"

	log "github.com/dulisoft/spirit/core/log/zapx"
	"github.com/dulisoft/spirit/core/middleware/ginMiddleWare"
	"github.com/dulisoft/spirit/core/transport/rest"
	"github.com/gin-gonic/gin"
)

// NewHttpServer 创建了一个绑定了路由的Web引擎
func NewHttpServer(r Router) *rest.Server {
	log.Info("start NewHttpServer")

	// step1: 默认启动一个Web引擎
	app := gin.New()

	// 默认注册recovery中间件
	app.Use(gin.Recovery())

	writer := log.NewZapWriter("basic_bigdata_service_request")
	logx.SetWriter(writer)

	app.Use(ginMiddleWare.GinZap(writer, time.RFC3339, false))
	app.Use(ginMiddleWare.RecoveryWithZap(writer, true))

	// 业务绑定路由操作
	r.Register(app)

	// step2: 返回绑定路由后的Web引擎
	httpSrv := rest.NewServer(app, rest.Address(settings.GetConfig().Server.HttpConf.Host))
	log.Info("end NewHttpServer")
	return httpSrv
}

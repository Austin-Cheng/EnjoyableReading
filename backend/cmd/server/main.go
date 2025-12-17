package main

import (
	"flag"
	"github.com/Austin-Cheng/EnjoyableReading/common/settings"
	"github.com/dulisoft/spirit"
	"github.com/dulisoft/spirit/core/validator"
	"runtime"
)

var appInfo = &spirit.Application{
	Name:     "af_basic_bigdata_service",
	Version:  "1.0",
	Addr:     ":8287",
	ConfPath: "cmd/server/config/",
}

func init() {
	flag.StringVar(&appInfo.ConfPath, "confPath", "cmd/server/config/", "config path, eg: -conf config.yaml")
	flag.StringVar(&appInfo.Addr, "addr", ":8287", "config path, eg: -addr 0.0.0.0:8287")
	flag.Parse()
	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum)
}

// @title       basic-bigdata-service
// @version     1.0.0.0
// @description AnyFabric标签和数据血缘，大数据基础服务
// @BasePath /api/basic-bigdata-service/v1
func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//初始化配置
	config := settings.InitConfig(appInfo)

	//初始化各种模块
	// 初始化验证器
	validator.SetupValidator()
	//app, cleanup, err := mock.InitApp(&bc)
	app, cleanup, err := InitApp(config, &config.Database)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err = app.Run(); err != nil {
		panic(err)
	}
}

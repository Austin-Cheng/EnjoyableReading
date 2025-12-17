package settings

import (
	"github.com/dulisoft/spirit"
	configuration "github.com/dulisoft/spirit/core/config"
	"github.com/dulisoft/spirit/core/config/sources"
	"github.com/dulisoft/spirit/core/config/sources/env"
	"github.com/dulisoft/spirit/core/config/sources/file"
)

func InitConfig(app *spirit.Application) *Config {
	paths := []string{app.ConfPath}
	sources := make([]sources.Source, len(paths)+1)
	sources[0] = env.NewSource()
	for i, path := range paths {
		sources[i+1] = file.NewSource(path)
	}
	configuration.Init(sources...)

	conf := configuration.Scan[Config]()

	//更新端口信息
	if app.Addr != "" {
		conf.Server.HttpConf.Host = app.Addr
	}

	ResetConfig(&conf)
	return &conf
}

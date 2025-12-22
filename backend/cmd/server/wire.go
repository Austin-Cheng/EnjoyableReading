//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Austin-Cheng/EnjoyableReading/adapter/driven"
	"github.com/Austin-Cheng/EnjoyableReading/adapter/driver"
	"github.com/Austin-Cheng/EnjoyableReading/common/settings"
	"github.com/Austin-Cheng/EnjoyableReading/domain"

	"github.com/google/wire"
)

var appSet = wire.NewSet(
	newApp,
	wire.NewSet(wire.Struct(new(AppRunner), "*")),
)

func InitApp(conf *settings.Config) (*AppRunner, func(), error) {
	panic(wire.Build(
		appSet,
		driven.Set,
		driver.Set,
		domain.Set,
	))
}

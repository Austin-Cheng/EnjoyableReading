package main

import (
	"github.com/Austin-Cheng/EnjoyableReading/common/settings"

	"github.com/dulisoft/spirit/core/store/database"
	"github.com/google/wire"
)

func InitApp(conf *settings.Config, dbOptions *database.Options) (*AppRunner, func(), error) {
	panic(wire.Build())
}

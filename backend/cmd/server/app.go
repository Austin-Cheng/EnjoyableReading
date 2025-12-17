package main

import (
	"github.com/dulisoft/spirit"
	"github.com/dulisoft/spirit/core/transport/rest"
)

type AppRunner struct {
	App *spirit.App
}

func newApp(hs *rest.Server) *spirit.App {
	return spirit.New(
		spirit.Name(appInfo.Name),
		spirit.Server(hs),
	)
}

func (a *AppRunner) Run() error {
	return a.App.Run()
}

package driver

import (
	"github.com/Austin-Cheng/EnjoyableReading/adapter/driver/controllers/tags"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	controllerSet,
	routerSet,
)

var routerSet = wire.NewSet(
	wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router))),
	NewHttpServer,
)

var controllerSet = wire.NewSet(
	tags.NewService,
)

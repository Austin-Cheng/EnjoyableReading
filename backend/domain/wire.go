package domain

import (
	domainTag "github.com/Austin-Cheng/EnjoyableReading/domain/tag/impl"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	domainTag.NewTag,
)

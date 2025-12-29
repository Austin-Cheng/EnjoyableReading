package domain

import (
	"github.com/Austin-Cheng/EnjoyableReading/domain/paper/impl"
	tagImpl "github.com/Austin-Cheng/EnjoyableReading/domain/tag/impl"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	tagImpl.NewTag,
	impl.NewPaper,
)

package driven

import (
	tagImpl "github.com/Austin-Cheng/EnjoyableReading/adapter/driven/repo/impl"
	"github.com/Austin-Cheng/EnjoyableReading/infrastructure/repository/db"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	db.NewClient,
	db.NewDB,
	repoSet,
)

var repoSet = wire.NewSet(
	tagImpl.NewTag,
)

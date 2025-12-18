package repo

import (
	"context"
	"github.com/Austin-Cheng/EnjoyableReading/infrastructure/repository/db/model"
)

type Tag interface {
	Create(ctx context.Context, tag *model.Tag) error
}

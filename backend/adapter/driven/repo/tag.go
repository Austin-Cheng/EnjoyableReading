package repo

import (
	"context"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/dto"
	"github.com/Austin-Cheng/EnjoyableReading/infrastructure/repository/db/model"
)

type Tag interface {
	Create(ctx context.Context, tag *model.TTag) (int64, error)
	Update(ctx context.Context, tag *model.TTag) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*dto.Tag, error)
	List(ctx context.Context, req *dto.ListTagReq) (int64, []*model.TTag, error)
	tools
}
type tools interface {
	CheckNameExists(ctx context.Context, id int64, name string) (bool, error)
	TagExists(ctx context.Context, id int64) error
}

package repo

import (
	"context"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/dto"
	"github.com/Austin-Cheng/EnjoyableReading/infrastructure/repository/db/model"
)

type Tag interface {
	Create(ctx context.Context, tag *model.TTag) error
	Update(ctx context.Context, tag *model.TTag) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*model.TTag, error)
	List(ctx context.Context, req *dto.ListTagReq) (int64, []*model.TTag, error)
}

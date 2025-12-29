package repo

import (
	"context"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/dto"
	"github.com/Austin-Cheng/EnjoyableReading/infrastructure/repository/db/model"
)

type Paper interface {
	Create(ctx context.Context, paper *model.TPaper) (int64, error)
	Update(ctx context.Context, paper *model.TPaper) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*model.TPaper, error)
	List(ctx context.Context, req *dto.ListPaperReq) (int64, []*model.TPaper, error)
}

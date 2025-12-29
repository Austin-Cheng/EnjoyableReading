package paper

import (
	"context"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/dto"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/response"
)

type UseCase interface {
	Create(ctx context.Context, req *dto.CreatePaperReq) (int64, error)
	Update(ctx context.Context, req *dto.UpdatePaperReq) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*dto.Paper, error)
	List(ctx context.Context, req *dto.ListPaperReq) (*response.PageResult[dto.Paper], error)
}

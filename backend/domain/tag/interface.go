package tag

import (
	"context"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/dto"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/response"
)

type UseCase interface {
	Create(ctx context.Context, tag *dto.CreateTagReq) (int64, error)
	Update(ctx context.Context, tag *dto.UpdateTagReq) error
	Get(ctx context.Context, id int64) (*dto.Tag, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, req *dto.ListTagReq) (*response.PageResult[dto.Tag], error)
}

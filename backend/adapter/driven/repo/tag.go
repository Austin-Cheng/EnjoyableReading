package repo

import (
	"context"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/response"
	"github.com/Austin-Cheng/EnjoyableReading/domain/tag"
	"github.com/Austin-Cheng/EnjoyableReading/infrastructure/repository/db/model"
)

type Tag interface {
	Create(ctx context.Context, tag *model.TTag) error
	List(ctx context.Context, req *tag.ListTagReq) (*response.PageResult[model.TTag], error)
}

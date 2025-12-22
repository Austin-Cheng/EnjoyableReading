package tag

import (
	"context"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/request"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/response"
	"github.com/Austin-Cheng/EnjoyableReading/infrastructure/repository/db/model"
)

type UseCase interface {
	Create(ctx context.Context, tag *CreateTagReq) error
	List(ctx context.Context, req *ListTagReq) (*response.PageResult[model.TTag], error)
}

type CreateTagReq struct {
	Name string
}

type ListTagReq struct {
	request.PageInfoWithKeyword
}

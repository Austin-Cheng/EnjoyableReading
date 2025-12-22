package impl

import (
	"context"
	"github.com/Austin-Cheng/EnjoyableReading/adapter/driven/repo"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/response"
	"github.com/Austin-Cheng/EnjoyableReading/domain/tag"
	"github.com/Austin-Cheng/EnjoyableReading/infrastructure/repository/db/model"
)

type tagImpl struct {
	tagRepo repo.Tag
}

func NewTag(tagRepo repo.Tag) tag.UseCase {
	return &tagImpl{
		tagRepo: tagRepo,
	}
}

func (t tagImpl) Create(ctx context.Context, req *tag.CreateTagReq) error {
	obj := &model.TTag{
		Name: req.Name,
	}
	return t.tagRepo.Create(ctx, obj)
}

func (t tagImpl) List(ctx context.Context, req *tag.ListTagReq) (*response.PageResult[model.TTag], error) {
	return t.tagRepo.List(ctx, req)
}

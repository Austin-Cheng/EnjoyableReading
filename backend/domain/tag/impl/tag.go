package impl

import (
	"context"
	"github.com/Austin-Cheng/EnjoyableReading/adapter/driven/repo"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/dto"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/response"
	"github.com/Austin-Cheng/EnjoyableReading/domain/tag"
	"github.com/Austin-Cheng/EnjoyableReading/infrastructure/repository/db/model"
	"github.com/samber/lo"
)

type tagImpl struct {
	tagRepo repo.Tag
}

func NewTag(tagRepo repo.Tag) tag.UseCase {
	return &tagImpl{
		tagRepo: tagRepo,
	}
}

func (t tagImpl) Create(ctx context.Context, req *dto.CreateTagReq) error {
	obj := &model.TTag{
		Name: req.Name,
	}
	return t.tagRepo.Create(ctx, obj)
}

func (t tagImpl) List(ctx context.Context, req *dto.ListTagReq) (*response.PageResult[dto.Tag], error) {
	total, ds, err := t.tagRepo.List(ctx, req)
	if err != nil {
		return nil, err
	}
	objs := lo.Times(len(ds), func(i int) *dto.Tag {
		return dto.ToTag(ds[i])
	})
	return &response.PageResult[dto.Tag]{
		TotalCount: total,
		Entries:    objs,
	}, nil
}

func (t tagImpl) Update(ctx context.Context, tag *dto.UpdateTagReq) error {
	return t.tagRepo.Update(ctx, tag.ToModel())
}

func (t tagImpl) Get(ctx context.Context, id int64) (*dto.Tag, error) {
	ds, err := t.tagRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToTag(ds), nil
}

func (t tagImpl) Delete(ctx context.Context, id int64) error {
	return t.tagRepo.Delete(ctx, id)
}

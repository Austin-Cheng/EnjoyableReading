package impl

import (
	"context"

	"github.com/Austin-Cheng/EnjoyableReading/adapter/driven/repo"
	"github.com/Austin-Cheng/EnjoyableReading/common/errorcode"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/dto"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/response"
	"github.com/Austin-Cheng/EnjoyableReading/domain/tag"
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

// Create  新建标签
// step1: 检查标签名称是否已存在，存在则返回错误
// step2: 如果有父级标签ID，检查父标签是否存在，不存在则返回错误
func (t tagImpl) Create(ctx context.Context, req *dto.CreateTagReq) (int64, error) {
	if req.ParentID > 0 {
		// step1: 检查标签是否存在
		existingTag, err := t.tagRepo.Get(ctx, req.ParentID)
		if err != nil {
			return 0, errorcode.PublicInvalidParameter.Detail("标签不存在")
		}
		if existingTag == nil {
			return 0, errorcode.PublicInvalidParameter.Detail("标签不存在")
		}
	}

	id, err := t.tagRepo.Create(ctx, req.ToModel())
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Update 更新标签，支持父标签更新
// step1: 检查标签是否存在，不存在则返回错误
// step2: 如果有父级标签ID，检查父标签是否存在，不存在则返回错误
// step3: 更新标签信息
func (t tagImpl) Update(ctx context.Context, tag *dto.UpdateTagReq) error {
	// step1: 检查标签是否存在
	if err := t.tagRepo.TagExists(ctx, tag.ID); err != nil {
		return err
	}
	return t.tagRepo.Update(ctx, tag.ToModel())
}

// List 获取标签列表
// step1: 构建查询条件
// step2: 执行查询
// step3: 转换结果
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

// Get 获取标签详情，会返回子标签列表
func (t tagImpl) Get(ctx context.Context, id int64) (*dto.Tag, error) {
	// step1: 检查标签是否存在
	if err := t.tagRepo.TagExists(ctx, id); err != nil {
		return nil, err
	}
	ds, err := t.tagRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return ds, nil
}

// Delete 删除标签，
// step1: 检查标签是否存在，不存在则返回错误
// step2: 删除标签，同时在一个事务中，删除所有的子标签
func (t tagImpl) Delete(ctx context.Context, id int64) error {
	// step1: 检查标签是否存在
	if err := t.tagRepo.TagExists(ctx, id); err != nil {
		return err
	}
	// step2: 删除标签，同时删除所有的子标签
	if err := t.tagRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

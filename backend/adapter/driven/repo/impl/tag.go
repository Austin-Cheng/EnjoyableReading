package impl

import (
	"context"
	"github.com/Austin-Cheng/EnjoyableReading/adapter/driven/repo"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/response"
	"github.com/Austin-Cheng/EnjoyableReading/domain/tag"
	"github.com/Austin-Cheng/EnjoyableReading/infrastructure/repository/db/model"
	"gorm.io/gorm"
)

type tagImpl struct {
	db *gorm.DB
}

func NewTag(db *gorm.DB) repo.Tag {
	return &tagImpl{
		db: db,
	}
}

func (t tagImpl) Create(ctx context.Context, tag *model.TTag) error {
	return t.db.Create(tag).Error
}

func (t tagImpl) List(ctx context.Context, req *tag.ListTagReq) (*response.PageResult[model.TTag], error) {
	db := t.db.WithContext(ctx).Model(&model.TTag{})
	if req.Keyword != "" {
		db = db.Where("name like ?", "%"+req.Keyword+"%")
	}
	resp := &response.PageResult[model.TTag]{}
	//总数
	if err := db.Count(&resp.TotalCount).Error; err != nil {
		return nil, err
	}
	// 排序
	db = db.Order(req.Sort + " " + req.Direction)
	// 分页
	db = db.Offset((req.Offset - 1) * req.Limit).Limit(req.Limit)
	// 查询结果
	if err := db.Find(&resp.Entries).Error; err != nil {
		return nil, err
	}
	return resp, nil
}

package impl

import (
	"context"
	"github.com/Austin-Cheng/EnjoyableReading/adapter/driven/repo"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/dto"
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

func (t tagImpl) Update(ctx context.Context, tag *model.TTag) error {
	return t.db.Updates(tag).Error
}

func (t tagImpl) Delete(ctx context.Context, id int64) error {
	return t.db.Delete(&model.TTag{ID: id}).Error
}

func (t tagImpl) Get(ctx context.Context, id int64) (*model.TTag, error) {
	data := &model.TTag{}
	err := t.db.WithContext(ctx).Where("id = ?", id).First(data).Error
	return data, err
}

func (t tagImpl) List(ctx context.Context, req *dto.ListTagReq) (int64, []*model.TTag, error) {
	db := t.db.WithContext(ctx).Model(&model.TTag{})
	if req.Keyword != "" {
		db = db.Where("name like ?", "%"+req.Keyword+"%")
	}
	//总数
	total := int64(0)
	if err := db.Count(&total).Error; err != nil {
		return 0, nil, err
	}
	// 排序
	db = db.Order(req.Sort + " " + req.Direction)
	// 分页
	db = db.Offset((req.Offset - 1) * req.Limit).Limit(req.Limit)
	// 查询结果
	datas := make([]*model.TTag, 0)
	if err := db.Find(&datas).Error; err != nil {
		return 0, nil, err
	}
	return total, datas, nil
}

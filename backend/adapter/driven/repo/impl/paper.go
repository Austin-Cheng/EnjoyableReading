package impl

import (
	"context"
	"github.com/Austin-Cheng/EnjoyableReading/adapter/driven/repo"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/dto"
	"github.com/Austin-Cheng/EnjoyableReading/infrastructure/repository/db/model"
	"gorm.io/gorm"
)

type paperImpl struct {
	db *gorm.DB
}

func NewPaper(db *gorm.DB) repo.Paper {
	return &paperImpl{
		db: db,
	}
}

// Create 创建文章
func (p paperImpl) Create(ctx context.Context, paper *model.TPaper) (int64, error) {
	err := p.db.WithContext(ctx).Create(paper).Error
	if err != nil {
		return 0, err
	}
	return int64(paper.ID), nil
}

// Update 更新文章
func (p paperImpl) Update(ctx context.Context, paper *model.TPaper) error {
	return p.db.WithContext(ctx).Updates(paper).Error
}

// Delete 删除文章
func (p paperImpl) Delete(ctx context.Context, id int64) error {
	return p.db.WithContext(ctx).Delete(&model.TPaper{ID: int32(id)}).Error
}

// Get 获取文章详情
func (p paperImpl) Get(ctx context.Context, id int64) (*model.TPaper, error) {
	data := &model.TPaper{}
	err := p.db.WithContext(ctx).Where("id = ?", id).First(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

// List 获取文章列表
func (p paperImpl) List(ctx context.Context, req *dto.ListPaperReq) (int64, []*model.TPaper, error) {
	db := p.db.WithContext(ctx).Model(&model.TPaper{})
	if req.Keyword != "" {
		db = db.Where("title like ? OR title_ch like ? OR summary like ? OR summary_ch like ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
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
	datas := make([]*model.TPaper, 0)
	if err := db.Find(&datas).Error; err != nil {
		return 0, nil, err
	}
	return total, datas, nil
}

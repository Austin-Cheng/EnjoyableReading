package impl

import (
	"context"
	"github.com/Austin-Cheng/EnjoyableReading/adapter/driven/repo"
	"github.com/Austin-Cheng/EnjoyableReading/common/errorcode"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/dto"
	"github.com/Austin-Cheng/EnjoyableReading/infrastructure/repository/db/model"
	"github.com/samber/lo"
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

// TagExists 检查标签名称是否重复
func (t tagImpl) TagExists(ctx context.Context, id int64) error {
	// 检查标签名称是否重复
	db := t.db.WithContext(ctx).Model(&model.TTag{})
	var count int64
	err := db.Where("id = ?", id).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		return errorcode.PublicTagNotExists.Err()
	}
	return nil
}

// CheckNameExists 检查标签名称是否重复
func (t tagImpl) CheckNameExists(ctx context.Context, id int64, name string) (bool, error) {
	// 检查标签名称是否重复
	db := t.db.WithContext(ctx).Model(&model.TTag{})
	var count int64
	err := db.Where("name = ? AND id <> ?", name, id).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return false, errorcode.PublicTagExists.Err()
	}
	return true, nil
}

// Create 创建标签
func (t tagImpl) Create(ctx context.Context, tag *model.TTag) (int64, error) {
	// 检查标签名称是否重复
	exists, err := t.CheckNameExists(ctx, tag.ID, tag.Name)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, errorcode.PublicTagExists.Err()
	}
	err = t.db.Create(tag).Error
	if err != nil {
		return 0, err
	}
	return tag.ID, nil
}

// Update 更新标签
func (t tagImpl) Update(ctx context.Context, tag *model.TTag) error {
	// 检查标签名称是否重复
	exists, err := t.CheckNameExists(ctx, tag.ID, tag.Name)
	if err != nil {
		return err
	}
	if !exists {
		return errorcode.PublicTagExists.Err()
	}
	return t.db.Updates(tag).Error
}

func (t tagImpl) deleteSelfAndChildren(tx *gorm.DB, id int64) error {
	// 查询所有子标签
	childTags := make([]model.TTag, 0)
	if err := tx.Model(&model.TTag{}).Where("parent_id = ?", id).Find(&childTags).Error; err != nil {
		return err
	}
	if err := tx.Delete(&model.TTag{ID: id}).Error; err != nil {
		return err
	}
	// 删除所有子标签
	for _, childTag := range childTags {
		return t.deleteSelfAndChildren(tx, childTag.ID)
	}
	return nil
}

func (t tagImpl) Delete(ctx context.Context, id int64) error {
	err := t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return t.deleteSelfAndChildren(tx, id)
	})
	return err
}

func (t tagImpl) getChildren(tx *gorm.DB, tag *dto.Tag) error {
	// 查询所有子标签
	childTags := make([]*model.TTag, 0)
	if err := tx.Model(&model.TTag{}).Where("parent_id = ?", tag.ID).Find(&childTags).Error; err != nil {
		return err
	}
	//结构给塞上
	tag.Children = lo.Times(len(childTags), func(i int) *dto.Tag {
		return dto.ToTag(childTags[i])
	})
	// 查询所有孙子标签
	for _, childTag := range tag.Children {
		if err := t.getChildren(tx, childTag); err != nil {
			return err
		}
	}
	return nil
}

func (t tagImpl) Get(ctx context.Context, id int64) (*dto.Tag, error) {
	data := &model.TTag{}
	err := t.db.WithContext(ctx).Where("id = ?", id).First(data).Error
	if err != nil {
		return nil, err
	}
	tag := dto.ToTag(data)
	if err = t.getChildren(t.db.WithContext(ctx), tag); err != nil {
		return nil, err
	}
	return tag, nil
}

func (t tagImpl) List(ctx context.Context, req *dto.ListTagReq) (int64, []*model.TTag, error) {
	db := t.db.WithContext(ctx).Model(&model.TTag{})
	if req.Keyword != "" {
		db = db.Where("name like ?", "%"+req.Keyword+"%")
	}
	if req.ID > 0 {
		db = db.Where("parent_id = ?", req.ID)
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

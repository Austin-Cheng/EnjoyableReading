package impl

import (
	"context"
	"github.com/Austin-Cheng/EnjoyableReading/adapter/driven/repo"
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

func (t tagImpl) Create(ctx context.Context, tag *model.Tag) error {
	return t.db.Create(tag).Error
}

package dto

import (
	"github.com/Austin-Cheng/EnjoyableReading/common/models/request"
	"github.com/Austin-Cheng/EnjoyableReading/infrastructure/repository/db/model"
)

type CreateTagReq struct {
	//ParentID int64  `json:"parent_id" binding:"omitempty,gt=0"`
	Name string `json:"name" binding:"required,max=255"`
}

type UpdateTagReq struct {
	request.IDReq
	Name string `json:"name" binding:"required,max=255"`
}

func (u *UpdateTagReq) ToModel() *model.TTag {
	return &model.TTag{
		ID:   u.ID,
		Name: u.Name,
	}
}

type ListTagReq struct {
	request.PageInfoWithKeyword
}

type Tag struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ParentID int64  `json:"parent_id"`
}

func ToTag(t *model.TTag) *Tag {
	return &Tag{
		ID:       t.ID,
		Name:     t.Name,
		ParentID: t.ParentID,
	}
}

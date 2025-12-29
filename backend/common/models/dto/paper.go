package dto

import (
	"time"

	"github.com/Austin-Cheng/EnjoyableReading/common/models/request"
	"github.com/Austin-Cheng/EnjoyableReading/infrastructure/repository/db/model"
)

type CreatePaperReq struct {
	Address       *string   `json:"address" binding:"omitempty,url"`
	Title         string    `json:"title" binding:"required,max=255"`
	TitleCh       string    `json:"title_ch" binding:"required,max=255"`
	Authors       *string   `json:"authors" binding:"omitempty,max=255"`
	PublishedTime time.Time `json:"published_time" binding:"required"`
	Summary       *string   `json:"summary" binding:"omitempty"`
	SummaryCh     *string   `json:"summary_ch" binding:"omitempty"`
	Categories    *string   `json:"categories" binding:"omitempty"`
	Filepath      *string   `json:"filepath" binding:"omitempty"`
	Read          int32     `json:"read" binding:"omitempty,min=0"`
	Favorite      int32     `json:"favorite" binding:"omitempty,min=0"`
	FulltextCh    *string   `json:"fulltext_ch" binding:"omitempty"`
}

func (c *CreatePaperReq) ToModel() *model.TPaper {
	return &model.TPaper{
		Address:       c.Address,
		Title:         c.Title,
		TitleCh:       c.TitleCh,
		Authors:       c.Authors,
		PublishedTime: c.PublishedTime,
		Summary:       c.Summary,
		SummaryCh:     c.SummaryCh,
		Categories:    c.Categories,
		Filepath:      c.Filepath,
		Read:          c.Read,
		Favorite:      c.Favorite,
		FulltextCh:    c.FulltextCh,
	}
}

type UpdatePaperReq struct {
	request.IDReq
	Address       *string   `json:"address" binding:"omitempty,url"`
	Title         string    `json:"title" binding:"required,max=255"`
	TitleCh       string    `json:"title_ch" binding:"required,max=255"`
	Authors       *string   `json:"authors" binding:"omitempty,max=255"`
	PublishedTime time.Time `json:"published_time" binding:"required"`
	Summary       *string   `json:"summary" binding:"omitempty"`
	SummaryCh     *string   `json:"summary_ch" binding:"omitempty"`
	Categories    *string   `json:"categories" binding:"omitempty"`
	Filepath      *string   `json:"filepath" binding:"omitempty"`
	Read          int32     `json:"read" binding:"omitempty,min=0"`
	Favorite      int32     `json:"favorite" binding:"omitempty,min=0"`
	FulltextCh    *string   `json:"fulltext_ch" binding:"omitempty"`
}

func (u *UpdatePaperReq) ToModel() *model.TPaper {
	return &model.TPaper{
		ID:            int32(u.ID),
		Address:       u.Address,
		Title:         u.Title,
		TitleCh:       u.TitleCh,
		Authors:       u.Authors,
		PublishedTime: u.PublishedTime,
		Summary:       u.Summary,
		SummaryCh:     u.SummaryCh,
		Categories:    u.Categories,
		Filepath:      u.Filepath,
		Read:          u.Read,
		Favorite:      u.Favorite,
		FulltextCh:    u.FulltextCh,
	}
}

type ListPaperReq struct {
	request.PageInfoWithKeyword
}

type Paper struct {
	ID            int32      `json:"id"`
	Address       *string    `json:"address"`
	Title         string     `json:"title"`
	TitleCh       string     `json:"title_ch"`
	Authors       *string    `json:"authors"`
	PublishedTime time.Time  `json:"published_time"`
	Summary       *string    `json:"summary"`
	SummaryCh     *string    `json:"summary_ch"`
	Categories    *string    `json:"categories"`
	Filepath      *string    `json:"filepath"`
	Read          int32      `json:"read"`
	Favorite      int32      `json:"favorite"`
	FulltextCh    *string    `json:"fulltext_ch"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}

func ToPaper(p *model.TPaper) *Paper {
	return &Paper{
		ID:            p.ID,
		Address:       p.Address,
		Title:         p.Title,
		TitleCh:       p.TitleCh,
		Authors:       p.Authors,
		PublishedTime: p.PublishedTime,
		Summary:       p.Summary,
		SummaryCh:     p.SummaryCh,
		Categories:    p.Categories,
		Filepath:      p.Filepath,
		Read:          p.Read,
		Favorite:      p.Favorite,
		FulltextCh:    p.FulltextCh,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

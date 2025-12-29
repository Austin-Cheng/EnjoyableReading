package impl

import (
	"context"
	"github.com/Austin-Cheng/EnjoyableReading/adapter/driven/repo"
	"github.com/Austin-Cheng/EnjoyableReading/common/errorcode"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/dto"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/response"
	"github.com/Austin-Cheng/EnjoyableReading/domain/paper"
	"github.com/samber/lo"
)

type paperImpl struct {
	paperRepo repo.Paper
}

func NewPaper(paperRepo repo.Paper) paper.UseCase {
	return &paperImpl{
		paperRepo: paperRepo,
	}
}

// Create 创建文章
func (p paperImpl) Create(ctx context.Context, req *dto.CreatePaperReq) (int64, error) {
	id, err := p.paperRepo.Create(ctx, req.ToModel())
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Update 更新文章
func (p paperImpl) Update(ctx context.Context, req *dto.UpdatePaperReq) error {
	// 检查文章是否存在
	existingPaper, err := p.paperRepo.Get(ctx, req.ID)
	if err != nil {
		return errorcode.PublicInvalidParameter.Detail("文章不存在")
	}
	if existingPaper == nil {
		return errorcode.PublicInvalidParameter.Detail("文章不存在")
	}

	return p.paperRepo.Update(ctx, req.ToModel())
}

// Delete 删除文章
func (p paperImpl) Delete(ctx context.Context, id int64) error {
	// 检查文章是否存在
	existingPaper, err := p.paperRepo.Get(ctx, id)
	if err != nil {
		return errorcode.PublicInvalidParameter.Detail("文章不存在")
	}
	if existingPaper == nil {
		return errorcode.PublicInvalidParameter.Detail("文章不存在")
	}

	return p.paperRepo.Delete(ctx, id)
}

// Get 获取文章详情
func (p paperImpl) Get(ctx context.Context, id int64) (*dto.Paper, error) {
	// 检查文章是否存在
	existingPaper, err := p.paperRepo.Get(ctx, id)
	if err != nil {
		return nil, errorcode.PublicInvalidParameter.Detail("文章不存在")
	}
	if existingPaper == nil {
		return nil, errorcode.PublicInvalidParameter.Detail("文章不存在")
	}

	return dto.ToPaper(existingPaper), nil
}

// List 获取文章列表
func (p paperImpl) List(ctx context.Context, req *dto.ListPaperReq) (*response.PageResult[dto.Paper], error) {
	total, papers, err := p.paperRepo.List(ctx, req)
	if err != nil {
		return nil, err
	}

	objs := lo.Times(len(papers), func(i int) *dto.Paper {
		return dto.ToPaper(papers[i])
	})

	return &response.PageResult[dto.Paper]{
		TotalCount: total,
		Entries:    objs,
	}, nil
}

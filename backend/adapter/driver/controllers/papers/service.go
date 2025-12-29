package papers

import (
	"net/http"

	"github.com/Austin-Cheng/EnjoyableReading/common/errorcode"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/dto"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/request"
	"github.com/Austin-Cheng/EnjoyableReading/domain/paper"
	"github.com/dulisoft/spirit/core/transport/rest/ginx"
	"github.com/dulisoft/spirit/core/validator"
	"github.com/gin-gonic/gin"
)

type Service struct {
	uc paper.UseCase
}

func NewService(uc paper.UseCase) *Service {
	return &Service{uc: uc}
}

// Create
// @Summary     新建文章
// @Description 新建文章
// @Tags 文章管理
// @Accept      json
// @Produce     json
// @Param       _  body     dto.CreatePaperReq  true "新建文章请求参数"
// @Success     200 {object} map[string]int64    "成功响应"
// @Failure     400 {object} ginx.HttpError     "失败响应参数"
// @Router      /paper [post]
func (s *Service) Create(c *gin.Context) {
	req := &dto.CreatePaperReq{}
	if _, err := validator.BindJsonAndValid(c, req); err != nil {
		ginx.ResErrJson(c, errorcode.PublicInvalidParameter.Detail(err))
		return
	}
	id, err := s.uc.Create(c, req)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		ginx.ResErrJson(c, err)
		return
	}
	ginx.ResOKJson(c, gin.H{"id": id})
}

// Update
// @Summary     更新文章
// @Description 更新文章信息
// @Tags 文章管理
// @Accept      json
// @Produce     json
// @Param       _  body     dto.UpdatePaperReq  true "更新文章请求参数"
// @Success     200 {object} nil               "成功响应"
// @Failure     400 {object} ginx.HttpError    "失败响应参数"
// @Router      /paper [put]
func (s *Service) Update(c *gin.Context) {
	req := &dto.UpdatePaperReq{}
	if _, err := validator.BindJsonAndValid(c, req); err != nil {
		ginx.ResErrJson(c, errorcode.PublicInvalidParameter.Detail(err))
		return
	}
	if err := s.uc.Update(c, req); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		ginx.ResErrJson(c, err)
		return
	}
	ginx.ResOKJson(c, nil)
}

// Get
// @Summary     获取文章详情
// @Description 根据ID获取文章详情
// @Tags 文章管理
// @Accept      json
// @Produce     json
// @Param       id path     int64  true "文章ID"
// @Success     200 {object} dto.Paper "文章详情"
// @Failure     400 {object} ginx.HttpError "失败响应参数"
// @Router      /paper/{id} [get]
func (s *Service) Get(c *gin.Context) {
	req := &request.IDReq{}
	if _, err := validator.BindUriAndValid(c, req); err != nil {
		ginx.ResErrJson(c, errorcode.PublicInvalidParameter.Detail(err))
		return
	}
	data, err := s.uc.Get(c, req.ID)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		ginx.ResErrJson(c, err)
		return
	}
	ginx.ResOKJson(c, data)
}

// Delete
// @Summary     删除文章
// @Description 根据ID删除文章
// @Tags 文章管理
// @Accept      json
// @Produce     json
// @Param       id path     int64  true "文章ID"
// @Success     200 {object} nil    "成功响应"
// @Failure     400 {object} ginx.HttpError "失败响应参数"
// @Router      /paper/{id} [delete]
func (s *Service) Delete(c *gin.Context) {
	req := &request.IDReq{}
	if _, err := validator.BindUriAndValid(c, req); err != nil {
		ginx.ResErrJson(c, errorcode.PublicInvalidParameter.Detail(err))
		return
	}
	if err := s.uc.Delete(c, req.ID); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		ginx.ResErrJson(c, err)
		return
	}
	ginx.ResOKJson(c, nil)
}

// List
// @Summary     获取文章列表
// @Description 获取文章列表
// @Tags 文章管理
// @Accept      json
// @Produce     json
// @Param       _  query    dto.ListPaperReq  false "文章列表请求参数"
// @Success     200 {object} []dto.Paper       "文章列表"
// @Failure     400 {object} ginx.HttpError  "失败响应参数"
// @Router      /paper [get]
func (s *Service) List(c *gin.Context) {
	req := &dto.ListPaperReq{}
	if _, err := validator.BindQueryAndValid(c, req); err != nil {
		ginx.ResErrJson(c, errorcode.PublicInvalidParameter.Detail(err))
		return
	}
	resp, err := s.uc.List(c, req)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		ginx.ResErrJson(c, err)
		return
	}
	ginx.ResOKJson(c, resp)
}

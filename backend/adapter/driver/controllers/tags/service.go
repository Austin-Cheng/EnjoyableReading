package tags

import (
	"github.com/Austin-Cheng/EnjoyableReading/common/models/response"
	"net/http"

	"github.com/Austin-Cheng/EnjoyableReading/common/errorcode"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/dto"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/request"
	"github.com/Austin-Cheng/EnjoyableReading/domain/tag"
	"github.com/dulisoft/spirit/core/transport/rest/ginx"
	"github.com/dulisoft/spirit/core/validator"
	"github.com/gin-gonic/gin"
)

type Service struct {
	uc tag.UseCase
}

func NewService(uc tag.UseCase) *Service {
	return &Service{uc: uc}
}

// Create
// @Summary     新建标签
// @Description 新建业务标签
// @Tags 业务标签管理
// @Accept      json
// @Produce     json
// @Param       _  body     dto.CreateTagReq  true "新建标签请求参数"
// @Success     200 {object} nil                "成功响应"
// @Failure     400 {object} ginx.HttpError     "失败响应参数"
// @Router      /tag [post]
func (s *Service) Create(c *gin.Context) {
	req := &dto.CreateTagReq{}
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
	ginx.ResOKJson(c, response.ToIDResp(id))
}

// Update
// @Summary     更新标签
// @Description 更新业务标签信息
// @Tags 业务标签管理
// @Accept      json
// @Produce     json
// @Param       _  body     dto.UpdateTagReq  true "更新标签请求参数"
// @Success     200 {object} nil               "成功响应"
// @Failure     400 {object} ginx.HttpError    "失败响应参数"
// @Router      /tag [put]
func (s *Service) Update(c *gin.Context) {
	req := &dto.UpdateTagReq{}
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
// @Summary     获取标签
// @Description 根据ID获取业务标签详情
// @Tags 业务标签管理
// @Accept      json
// @Produce     json
// @Param       id path     int64  true "标签ID"
// @Success     200 {object} dto.Tag "标签详情"
// @Failure     400 {object} ginx.HttpError "失败响应参数"
// @Router      /tag/{id} [get]
func (s *Service) Get(c *gin.Context) {
	req := &request.IDReq{}
	if _, err := validator.BindJsonAndValid(c, req); err != nil {
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
// @Summary     删除标签
// @Description 根据ID删除业务标签
// @Tags 业务标签管理
// @Accept      json
// @Produce     json
// @Param       id path     int64  true "标签ID"
// @Success     200 {object} nil    "成功响应"
// @Failure     400 {object} ginx.HttpError "失败响应参数"
// @Router      /tag/{id} [delete]
func (s *Service) Delete(c *gin.Context) {
	req := &request.IDReq{}
	if _, err := validator.BindJsonAndValid(c, req); err != nil {
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
// @Summary     获取标签列表
// @Description 获取业务标签列表
// @Tags 业务标签管理
// @Accept      json
// @Produce     json
// @Param       _  query    dto.ListTagReq  false "标签列表请求参数"
// @Success     200 {object} []dto.Tag       "标签列表"
// @Failure     400 {object} ginx.HttpError  "失败响应参数"
// @Router      /tag [get]
func (s *Service) List(c *gin.Context) {
	req := &dto.ListTagReq{}
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

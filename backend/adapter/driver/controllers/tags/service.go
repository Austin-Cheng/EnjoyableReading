package tags

import (
	"github.com/Austin-Cheng/EnjoyableReading/common/errorcode"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/dto"
	"github.com/Austin-Cheng/EnjoyableReading/common/models/request"
	"github.com/Austin-Cheng/EnjoyableReading/domain/tag"
	"github.com/dulisoft/spirit/core/transport/rest/ginx"
	"github.com/dulisoft/spirit/core/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Service struct {
	uc tag.UseCase
}

func NewService(uc tag.UseCase) *Service {
	return &Service{uc: uc}
}

// Create
// @Summary     内部系统使用标签页面分页标签分类
// @Description 内部系统使用标签页面分页标签分类
// @Tags 业务标签管理
// @Accept      x-www-form-urlencoded
// @Produce     json
// @Param       _  query    domain.QueryCategoryPageReq  false "查询参数"
// @Success     200 {object} domain.QueryCategoryPageResp "成功响应参数"
// @Failure     400 {object} rest.HttpError     "失败响应参数"
// @Router      /label/category/page [get]
func (s *Service) Create(c *gin.Context) {
	req := &dto.CreateTagReq{}
	if _, err := validator.BindJsonAndValid(c, req); err != nil {
		ginx.ResErrJson(c, errorcode.PublicInvalidParameter.Detail(err))
		return
	}
	if err := s.uc.Create(c, req); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		ginx.ResErrJson(c, err)
		return
	}
	ginx.ResOKJson(c, nil)
}

// Update
// @Summary     内部系统使用标签页面分页标签分类
// @Description 内部系统使用标签页面分页标签分类
// @Tags 业务标签管理
// @Accept      x-www-form-urlencoded
// @Produce     json
// @Param       _  query    domain.QueryCategoryPageReq  false "查询参数"
// @Success     200 {object} domain.QueryCategoryPageResp "成功响应参数"
// @Failure     400 {object} rest.HttpError     "失败响应参数"
// @Router      /label/category/page [get]
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
// @Summary     内部系统使用标签页面分页标签分类
// @Description 内部系统使用标签页面分页标签分类
// @Tags 业务标签管理
// @Accept      x-www-form-urlencoded
// @Produce     json
// @Param       _  query    domain.QueryCategoryPageReq  false "查询参数"
// @Success     200 {object} domain.QueryCategoryPageResp "成功响应参数"
// @Failure     400 {object} rest.HttpError     "失败响应参数"
// @Router      /label/category/page [get]
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
// @Summary     内部系统使用标签页面分页标签分类
// @Description 内部系统使用标签页面分页标签分类
// @Tags 业务标签管理
// @Accept      x-www-form-urlencoded
// @Produce     json
// @Param       _  query    domain.QueryCategoryPageReq  false "查询参数"
// @Success     200 {object} domain.QueryCategoryPageResp "成功响应参数"
// @Failure     400 {object} rest.HttpError     "失败响应参数"
// @Router      /label/category/page [get]
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
// @Summary     内部系统使用标签页面分页标签分类
// @Description 内部系统使用标签页面分页标签分类
// @Tags 业务标签管理
// @Accept      x-www-form-urlencoded
// @Produce     json
// @Param       _  query    domain.QueryCategoryPageReq  false "查询参数"
// @Success     200 {object} domain.QueryCategoryPageResp "成功响应参数"
// @Failure     400 {object} rest.HttpError     "失败响应参数"
// @Router      /label/category/page [get]
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

package blog

import (
	"elian-blog/internal/service"
	"elian-blog/pkg/response"

	"github.com/gin-gonic/gin"
)

type PageController struct {
	svc *service.PageService
}

func NewPageController(svc *service.PageService) *PageController {
	return &PageController{svc: svc}
}

func (ctrl *PageController) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")
	page, err := ctrl.svc.GetBySlug(slug)
	if err != nil {
		response.Fail(c, 404, "页面不存在")
		return
	}
	response.Ok(c, page)
}

func (ctrl *PageController) List(c *gin.Context) {
	pages, err := ctrl.svc.List()
	if err != nil {
		response.InternalError(c, "获取页面列表失败")
		return
	}
	response.Ok(c, pages)
}

type SiteController struct {
	svc *service.SiteConfigService
}

func NewSiteController(svc *service.SiteConfigService) *SiteController {
	return &SiteController{svc: svc}
}

func (ctrl *SiteController) GetConfig(c *gin.Context) {
	configs, err := ctrl.svc.List()
	if err != nil {
		response.InternalError(c, "获取配置失败")
		return
	}
	response.Ok(c, configs)
}

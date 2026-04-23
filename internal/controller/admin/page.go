package admin

import (
	"elian-blog/internal/model"
	"elian-blog/internal/service"
	"elian-blog/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PageController struct {
	svc *service.PageService
}

func NewPageController(svc *service.PageService) *PageController {
	return &PageController{svc: svc}
}

func (ctrl *PageController) Create(c *gin.Context) {
	page := &model.Page{}
	if err := c.ShouldBindJSON(page); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}
	if err := ctrl.svc.Create(page); err != nil {
		response.InternalError(c, "创建页面失败")
		return
	}
	response.Ok(c, page)
}

func (ctrl *PageController) Update(c *gin.Context) {
	page := &model.Page{}
	if err := c.ShouldBindJSON(page); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}
	if err := ctrl.svc.Update(page); err != nil {
		response.InternalError(c, "更新页面失败")
		return
	}
	response.Ok(c, page)
}

func (ctrl *PageController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.svc.Delete(uint(id)); err != nil {
		response.InternalError(c, "删除失败")
		return
	}
	response.Ok(c, nil)
}

func (ctrl *PageController) List(c *gin.Context) {
	pages, err := ctrl.svc.ListAdmin()
	if err != nil {
		response.InternalError(c, "获取页面失败")
		return
	}
	response.Ok(c, pages)
}

type SiteConfigController struct {
	svc *service.SiteConfigService
}

func NewSiteConfigController(svc *service.SiteConfigService) *SiteConfigController {
	return &SiteConfigController{svc: svc}
}

func (ctrl *SiteConfigController) Get(c *gin.Context) {
	configs, err := ctrl.svc.List()
	if err != nil {
		response.InternalError(c, "获取配置失败")
		return
	}
	response.Ok(c, configs)
}

func (ctrl *SiteConfigController) Set(c *gin.Context) {
	var req struct {
		Key   string `json:"key" binding:"required"`
		Value string `json:"value" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}
	if err := ctrl.svc.Set(req.Key, req.Value); err != nil {
		response.InternalError(c, "保存配置失败")
		return
	}
	response.Ok(c, nil)
}

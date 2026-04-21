package admin

import (
	"elian-blog/internal/model"
	"elian-blog/internal/service"
	"elian-blog/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	svc *service.CategoryService
}

func NewCategoryController(svc *service.CategoryService) *CategoryController {
	return &CategoryController{svc: svc}
}

func (ctrl *CategoryController) Create(c *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Sort        int    `json:"sort"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	cat, err := ctrl.svc.Create(req.Name, req.Description, req.Sort)
	if err != nil {
		response.InternalError(c, "创建分类失败")
		return
	}
	response.Ok(c, cat)
}

func (ctrl *CategoryController) Update(c *gin.Context) {
	cat := &model.Category{}
	if err := c.ShouldBindJSON(cat); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}
	if err := ctrl.svc.Update(cat); err != nil {
		response.InternalError(c, "更新分类失败")
		return
	}
	response.Ok(c, cat)
}

func (ctrl *CategoryController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.svc.Delete(uint(id)); err != nil {
		response.InternalError(c, "删除失败")
		return
	}
	response.Ok(c, nil)
}

func (ctrl *CategoryController) List(c *gin.Context) {
	cats, err := ctrl.svc.List()
	if err != nil {
		response.InternalError(c, "获取分类失败")
		return
	}
	response.Ok(c, cats)
}

type TagController struct {
	svc *service.TagService
}

func NewTagController(svc *service.TagService) *TagController {
	return &TagController{svc: svc}
}

func (ctrl *TagController) Create(c *gin.Context) {
	var req struct {
		Name  string `json:"name" binding:"required"`
		Color string `json:"color"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	tag, err := ctrl.svc.Create(req.Name, req.Color)
	if err != nil {
		response.InternalError(c, "创建标签失败")
		return
	}
	response.Ok(c, tag)
}

func (ctrl *TagController) Update(c *gin.Context) {
	tag := &model.Tag{}
	if err := c.ShouldBindJSON(tag); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}
	if err := ctrl.svc.Update(tag); err != nil {
		response.InternalError(c, "更新标签失败")
		return
	}
	response.Ok(c, tag)
}

func (ctrl *TagController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.svc.Delete(uint(id)); err != nil {
		response.InternalError(c, "删除失败")
		return
	}
	response.Ok(c, nil)
}

func (ctrl *TagController) List(c *gin.Context) {
	tags, err := ctrl.svc.List()
	if err != nil {
		response.InternalError(c, "获取标签失败")
		return
	}
	response.Ok(c, tags)
}
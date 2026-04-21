package admin

import (
	"elian-blog/internal/model"
	"elian-blog/internal/service"
	"elian-blog/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	svc     *service.RoleService
	menuSvc *service.MenuService
}

func NewRoleController(svc *service.RoleService, menuSvc *service.MenuService) *RoleController {
	return &RoleController{svc: svc, menuSvc: menuSvc}
}

func (ctrl *RoleController) Create(c *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		Label       string `json:"label" binding:"required"`
		Description string `json:"description"`
		Sort        int    `json:"sort"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	role, err := ctrl.svc.Create(req.Name, req.Label, req.Description, req.Sort)
	if err != nil {
		response.InternalError(c, "创建角色失败")
		return
	}
	response.Ok(c, role)
}

func (ctrl *RoleController) Update(c *gin.Context) {
	role := &model.Role{}
	if err := c.ShouldBindJSON(role); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}
	if err := ctrl.svc.Update(role); err != nil {
		response.InternalError(c, "更新角色失败")
		return
	}
	response.Ok(c, role)
}

func (ctrl *RoleController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.svc.Delete(uint(id)); err != nil {
		response.InternalError(c, "删除失败")
		return
	}
	response.Ok(c, nil)
}

func (ctrl *RoleController) List(c *gin.Context) {
	roles, err := ctrl.svc.List()
	if err != nil {
		response.InternalError(c, "获取角色失败")
		return
	}
	response.Ok(c, roles)
}

func (ctrl *RoleController) UpdateMenus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		MenuIDs []uint `json:"menu_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}
	if err := ctrl.svc.UpdateMenus(uint(id), req.MenuIDs); err != nil {
		response.InternalError(c, "更新角色菜单失败")
		return
	}
	response.Ok(c, nil)
}

type MenuController struct {
	svc *service.MenuService
}

func NewMenuController(svc *service.MenuService) *MenuController {
	return &MenuController{svc: svc}
}

func (ctrl *MenuController) Create(c *gin.Context) {
	menu := &model.Menu{}
	if err := c.ShouldBindJSON(menu); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}
	if err := ctrl.svc.Create(menu); err != nil {
		response.InternalError(c, "创建菜单失败")
		return
	}
	response.Ok(c, menu)
}

func (ctrl *MenuController) Update(c *gin.Context) {
	menu := &model.Menu{}
	if err := c.ShouldBindJSON(menu); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}
	if err := ctrl.svc.Update(menu); err != nil {
		response.InternalError(c, "更新菜单失败")
		return
	}
	response.Ok(c, menu)
}

func (ctrl *MenuController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.svc.Delete(uint(id)); err != nil {
		response.InternalError(c, "删除失败")
		return
	}
	response.Ok(c, nil)
}

func (ctrl *MenuController) List(c *gin.Context) {
	menus, err := ctrl.svc.List()
	if err != nil {
		response.InternalError(c, "获取菜单失败")
		return
	}
	tree := ctrl.svc.BuildTree(menus)
	response.Ok(c, tree)
}
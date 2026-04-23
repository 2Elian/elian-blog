package admin

import (
	"elian-blog/internal/service"
	"elian-blog/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	svc *service.UserService
}

func NewUserController(svc *service.UserService) *UserController {
	return &UserController{svc: svc}
}

func (ctrl *UserController) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	users, total, err := ctrl.svc.List(page, pageSize)
	if err != nil {
		response.InternalError(c, "获取用户列表失败")
		return
	}
	response.OkPage(c, users, total, page, pageSize)
}

func (ctrl *UserController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := ctrl.svc.GetByID(uint(id))
	if err != nil {
		response.Fail(c, 404, "用户不存在")
		return
	}
	response.Ok(c, user)
}

func (ctrl *UserController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.svc.Delete(uint(id)); err != nil {
		response.InternalError(c, "删除失败")
		return
	}
	response.Ok(c, nil)
}

type CommentController struct {
	svc *service.CommentService
}

func NewCommentController(svc *service.CommentService) *CommentController {
	return &CommentController{svc: svc}
}

func (ctrl *CommentController) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	comments, total, err := ctrl.svc.ListAdmin(page, pageSize)
	if err != nil {
		response.InternalError(c, "获取评论列表失败")
		return
	}
	response.OkPage(c, comments, total, page, pageSize)
}

func (ctrl *CommentController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.svc.Delete(uint(id)); err != nil {
		response.InternalError(c, "删除失败")
		return
	}
	response.Ok(c, nil)
}

func (ctrl *CommentController) UpdateStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		Status int `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}
	if err := ctrl.svc.UpdateStatus(uint(id), req.Status); err != nil {
		response.InternalError(c, "更新状态失败")
		return
	}
	response.Ok(c, nil)
}

package admin

import (
	"elian-blog/internal/model"
	"elian-blog/internal/service"
	"elian-blog/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FriendLinkController struct {
	svc *service.FriendLinkService
}

func NewFriendLinkController(svc *service.FriendLinkService) *FriendLinkController {
	return &FriendLinkController{svc: svc}
}

func (ctrl *FriendLinkController) Create(c *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		URL         string `json:"url" binding:"required"`
		Logo        string `json:"logo"`
		Description string `json:"description"`
		Sort        int    `json:"sort"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	link, err := ctrl.svc.Create(req.Name, req.URL, req.Logo, req.Description, req.Sort)
	if err != nil {
		response.InternalError(c, "创建友链失败")
		return
	}
	response.Ok(c, link)
}

func (ctrl *FriendLinkController) Update(c *gin.Context) {
	link := &model.FriendLink{}
	if err := c.ShouldBindJSON(link); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}
	if err := ctrl.svc.Update(link); err != nil {
		response.InternalError(c, "更新友链失败")
		return
	}
	response.Ok(c, link)
}

func (ctrl *FriendLinkController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.svc.Delete(uint(id)); err != nil {
		response.InternalError(c, "删除失败")
		return
	}
	response.Ok(c, nil)
}

func (ctrl *FriendLinkController) List(c *gin.Context) {
	links, err := ctrl.svc.ListAdmin()
	if err != nil {
		response.InternalError(c, "获取友链失败")
		return
	}
	response.Ok(c, links)
}

type MessageController struct {
	svc *service.MessageService
}

func NewMessageController(svc *service.MessageService) *MessageController {
	return &MessageController{svc: svc}
}

func (ctrl *MessageController) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	messages, total, err := ctrl.svc.ListAdmin(page, pageSize)
	if err != nil {
		response.InternalError(c, "获取留言失败")
		return
	}
	response.OkPage(c, messages, total, page, pageSize)
}

func (ctrl *MessageController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.svc.Delete(uint(id)); err != nil {
		response.InternalError(c, "删除失败")
		return
	}
	response.Ok(c, nil)
}

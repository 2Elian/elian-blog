package blog

import (
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

func (ctrl *FriendLinkController) List(c *gin.Context) {
	links, err := ctrl.svc.List()
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

	messages, total, err := ctrl.svc.List(page, pageSize)
	if err != nil {
		response.InternalError(c, "获取留言失败")
		return
	}
	response.OkPage(c, messages, total, page, pageSize)
}

func (ctrl *MessageController) Create(c *gin.Context) {
	req := &service.MessageCreateReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	userID, _ := c.Get("user_id")
	uid, _ := userID.(uint)

	if err := ctrl.svc.Create(uid, req); err != nil {
		response.InternalError(c, "留言失败")
		return
	}
	response.Ok(c, nil)
}
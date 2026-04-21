package blog

import (
	"elian-blog/internal/service"
	"elian-blog/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	svc *service.AuthService
}

func NewAuthController(svc *service.AuthService) *AuthController {
	return &AuthController{svc: svc}
}

func (ctrl *AuthController) Login(c *gin.Context) {
	req := &service.LoginReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	resp, err := ctrl.svc.Login(req)
	if err != nil {
		response.Fail(c, 400, err.Error())
		return
	}
	response.Ok(c, resp)
}

func (ctrl *AuthController) Register(c *gin.Context) {
	req := &service.RegisterReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	resp, err := ctrl.svc.Register(req)
	if err != nil {
		response.Fail(c, 400, err.Error())
		return
	}
	response.Ok(c, resp)
}

func (ctrl *AuthController) GetUserInfo(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := userID.(uint)

	user, err := ctrl.svc.GetUserInfo(uid)
	if err != nil {
		response.Fail(c, 404, "用户不存在")
		return
	}
	response.Ok(c, user)
}
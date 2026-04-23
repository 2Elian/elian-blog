package admin

import (
	"elian-blog/internal/service"
	"elian-blog/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	svc *service.ArticleService
}

func NewArticleController(svc *service.ArticleService) *ArticleController {
	return &ArticleController{svc: svc}
}

func (ctrl *ArticleController) Create(c *gin.Context) {
	req := &service.ArticleCreateReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	userID, _ := c.Get("user_id")
	req.AuthorID = userID.(uint)

	article, err := ctrl.svc.Create(req)
	if err != nil {
		response.InternalError(c, "创建文章失败")
		return
	}
	response.Ok(c, article)
}

func (ctrl *ArticleController) Update(c *gin.Context) {
	req := &service.ArticleUpdateReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	article, err := ctrl.svc.Update(req)
	if err != nil {
		response.InternalError(c, "更新文章失败")
		return
	}
	response.Ok(c, article)
}

func (ctrl *ArticleController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.svc.Delete(uint(id)); err != nil {
		response.InternalError(c, "删除失败")
		return
	}
	response.Ok(c, nil)
}

func (ctrl *ArticleController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := ctrl.svc.GetByID(uint(id))
	if err != nil {
		response.Fail(c, 404, "文章不存在")
		return
	}
	response.Ok(c, article)
}

func (ctrl *ArticleController) List(c *gin.Context) {
	req := &service.ArticleQueryReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	articles, total, err := ctrl.svc.List(req)
	if err != nil {
		response.InternalError(c, "获取文章列表失败")
		return
	}
	response.OkPage(c, articles, total, req.Page, req.PageSize)
}

package blog

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

func (ctrl *ArticleController) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	articles, total, err := ctrl.svc.ListPublished(page, pageSize)
	if err != nil {
		response.InternalError(c, "获取文章列表失败")
		return
	}

	response.OkPage(c, articles, total, page, pageSize)
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

func (ctrl *ArticleController) Search(c *gin.Context) {
	req := &service.ArticleQueryReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	articles, total, err := ctrl.svc.List(req)
	if err != nil {
		response.InternalError(c, "搜索失败")
		return
	}

	response.OkPage(c, articles, total, req.Page, req.PageSize)
}
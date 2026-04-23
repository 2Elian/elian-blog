package blog

import (
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

func (ctrl *CategoryController) List(c *gin.Context) {
	categories, err := ctrl.svc.List()
	if err != nil {
		response.InternalError(c, "获取分类失败")
		return
	}
	response.Ok(c, categories)
}

type TagController struct {
	svc *service.TagService
}

func NewTagController(svc *service.TagService) *TagController {
	return &TagController{svc: svc}
}

func (ctrl *TagController) List(c *gin.Context) {
	tags, err := ctrl.svc.List()
	if err != nil {
		response.InternalError(c, "获取标签失败")
		return
	}
	response.Ok(c, tags)
}

type CommentController struct {
	svc *service.CommentService
}

func NewCommentController(svc *service.CommentService) *CommentController {
	return &CommentController{svc: svc}
}

func (ctrl *CommentController) ListByArticle(c *gin.Context) {
	articleID, _ := strconv.Atoi(c.Param("id"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	comments, total, err := ctrl.svc.ListByArticle(uint(articleID), page, pageSize)
	if err != nil {
		response.InternalError(c, "获取评论失败")
		return
	}
	response.OkPage(c, comments, total, page, pageSize)
}

func (ctrl *CommentController) Create(c *gin.Context) {
	req := &service.CommentCreateReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	userID, _ := c.Get("user_id")
	uid, _ := userID.(uint)

	if err := ctrl.svc.Create(uid, req); err != nil {
		response.InternalError(c, "评论失败")
		return
	}
	response.Ok(c, nil)
}

func (ctrl *CommentController) Recent(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	comments, err := ctrl.svc.ListRecent(limit)
	if err != nil {
		response.InternalError(c, "获取评论失败")
		return
	}
	response.Ok(c, comments)
}

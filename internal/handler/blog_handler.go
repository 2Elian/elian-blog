package handler

import (
	"net/http"

	"elian-blog/internal/logic/blog"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// --- 文章 ---

func ListArticlesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryArticleHomeReq
		if err := httpx.Parse(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		list, total, err := blog.NewArticleLogic(svcCtx).ListArticles(r.Context(), &req)
		if err != nil {
			fail(w, 500, "获取文章失败")
			return
		}
		okPage(w, list, total, req.Page, req.PageSize)
	}
}

func GetArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.ParsePath(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := blog.NewArticleLogic(svcCtx).GetArticle(r.Context(), req.ID)
		if err != nil {
			fail(w, 404, "文章不存在")
			return
		}
		ok(w, data)
	}
}

func SearchArticlesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryArticleHomeReq
		if err := httpx.Parse(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		list, total, err := blog.NewArticleLogic(svcCtx).ListArticles(r.Context(), &req)
		if err != nil {
			fail(w, 500, "搜索失败")
			return
		}
		okPage(w, list, total, req.Page, req.PageSize)
	}
}

// --- 认证 ---

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := blog.NewAuthLogic(svcCtx).Login(r.Context(), &req)
		if err != nil {
			fail(w, 401, err.Error())
			return
		}
		ok(w, data)
	}
}

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := blog.NewAuthLogic(svcCtx).Register(r.Context(), &req)
		if err != nil {
			fail(w, 400, err.Error())
			return
		}
		ok(w, data)
	}
}

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := blog.NewAuthLogic(svcCtx).GetUserInfo(r.Context())
		if err != nil {
			fail(w, 401, "未登录")
			return
		}
		ok(w, data)
	}
}

// --- 分类 ---

func ListCategoriesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := blog.NewCategoryLogic(svcCtx).List(r.Context())
		if err != nil {
			fail(w, 500, "获取分类失败")
			return
		}
		ok(w, data)
	}
}

// --- 标签 ---

func ListTagsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := blog.NewTagLogic(svcCtx).List(r.Context())
		if err != nil {
			fail(w, 500, "获取标签失败")
			return
		}
		ok(w, data)
	}
}

// --- 评论 ---

func ListCommentsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := blog.NewCommentLogic(svcCtx).ListByArticle(r.Context(), &req)
		if err != nil {
			fail(w, 500, "获取评论失败")
			return
		}
		ok(w, data)
	}
}

func RecentCommentsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := blog.NewCommentLogic(svcCtx).Recent(r.Context())
		if err != nil {
			fail(w, 500, "获取评论失败")
			return
		}
		ok(w, data)
	}
}

func CreateCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateCommentReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := blog.NewCommentLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			fail(w, 500, "评论失败")
			return
		}
		ok(w, data)
	}
}

// --- 友链 ---

func ListFriendLinksHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := blog.NewFriendLinkLogic(svcCtx).List(r.Context())
		if err != nil {
			fail(w, 500, "获取友链失败")
			return
		}
		ok(w, data)
	}
}

// --- 留言 ---

func ListMessagesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageQuery
		if err := httpx.Parse(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := blog.NewMessageLogic(svcCtx).List(r.Context(), &req)
		if err != nil {
			fail(w, 500, "获取留言失败")
			return
		}
		ok(w, data)
	}
}

func CreateMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateMessageReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := blog.NewMessageLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			fail(w, 500, "留言失败")
			return
		}
		ok(w, data)
	}
}

// --- 页面 ---

func ListPagesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := blog.NewPageLogic(svcCtx).List(r.Context())
		if err != nil {
			fail(w, 500, "获取页面失败")
			return
		}
		ok(w, data)
	}
}

func GetPageBySlugHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Slug string `path:"slug"`
		}
		if err := httpx.ParsePath(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := blog.NewPageLogic(svcCtx).GetBySlug(r.Context(), req.Slug)
		if err != nil {
			fail(w, 404, "页面不存在")
			return
		}
		ok(w, data)
	}
}

// --- 站点配置 ---

func GetSiteConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := blog.NewSiteLogic(svcCtx).GetConfig(r.Context())
		if err != nil {
			fail(w, 500, "获取配置失败")
			return
		}
		ok(w, data)
	}
}

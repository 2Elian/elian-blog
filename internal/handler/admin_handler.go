package handler

import (
	"net/http"

	"elian-blog/internal/logic/admin"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// --- 文章管理 ---

func AdminCreateArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateArticleReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := admin.NewArticleLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			fail(w, 500, "创建文章失败")
			return
		}
		ok(w, data)
	}
}

func AdminUpdateArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateArticleReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewArticleLogic(svcCtx).Update(r.Context(), &req)
		if err != nil {
			fail(w, 500, "更新文章失败")
			return
		}
		ok(w, nil)
	}
}

func AdminDeleteArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.ParsePath(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewArticleLogic(svcCtx).Delete(r.Context(), req.ID)
		if err != nil {
			fail(w, 500, "删除失败")
			return
		}
		ok(w, nil)
	}
}

func AdminListArticlesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryArticleHomeReq
		if err := httpx.Parse(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		list, total, err := admin.NewArticleLogic(svcCtx).List(r.Context(), &req)
		if err != nil {
			fail(w, 500, "获取文章失败")
			return
		}
		okPage(w, list, total, req.Page, req.PageSize)
	}
}

func AdminGetArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.ParsePath(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := admin.NewArticleLogic(svcCtx).Get(r.Context(), req.ID)
		if err != nil {
			fail(w, 404, "文章不存在")
			return
		}
		ok(w, data)
	}
}

// --- 分类管理 ---

func AdminCreateCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateCategoryReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := admin.NewCategoryLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			fail(w, 500, "创建分类失败")
			return
		}
		ok(w, data)
	}
}

func AdminUpdateCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCategoryReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewCategoryLogic(svcCtx).Update(r.Context(), &req)
		if err != nil {
			fail(w, 500, "更新分类失败")
			return
		}
		ok(w, nil)
	}
}

func AdminDeleteCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.ParsePath(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewCategoryLogic(svcCtx).Delete(r.Context(), req.ID)
		if err != nil {
			fail(w, 500, "删除失败")
			return
		}
		ok(w, nil)
	}
}

func AdminListCategoriesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewCategoryLogic(svcCtx).List(r.Context())
		if err != nil {
			fail(w, 500, "获取分类失败")
			return
		}
		ok(w, data)
	}
}

// --- 标签管理 ---

func AdminCreateTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateTagReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := admin.NewTagLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			fail(w, 500, "创建标签失败")
			return
		}
		ok(w, data)
	}
}

func AdminUpdateTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateTagReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewTagLogic(svcCtx).Update(r.Context(), &req)
		if err != nil {
			fail(w, 500, "更新标签失败")
			return
		}
		ok(w, nil)
	}
}

func AdminDeleteTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.ParsePath(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewTagLogic(svcCtx).Delete(r.Context(), req.ID)
		if err != nil {
			fail(w, 500, "删除失败")
			return
		}
		ok(w, nil)
	}
}

func AdminListTagsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewTagLogic(svcCtx).List(r.Context())
		if err != nil {
			fail(w, 500, "获取标签失败")
			return
		}
		ok(w, data)
	}
}

// --- 用户管理 ---

func AdminListUsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageQuery
		if err := httpx.Parse(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := admin.NewUserLogic(svcCtx).List(r.Context(), &req)
		if err != nil {
			fail(w, 500, "获取用户失败")
			return
		}
		ok(w, data)
	}
}

func AdminDeleteUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.ParsePath(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewUserLogic(svcCtx).Delete(r.Context(), req.ID)
		if err != nil {
			fail(w, 500, "删除失败")
			return
		}
		ok(w, nil)
	}
}

// --- 评论管理 ---

func AdminListCommentsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := admin.NewCommentLogic(svcCtx).List(r.Context(), &req)
		if err != nil {
			fail(w, 500, "获取评论失败")
			return
		}
		ok(w, data)
	}
}

func AdminDeleteCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.ParsePath(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewCommentLogic(svcCtx).Delete(r.Context(), req.ID)
		if err != nil {
			fail(w, 500, "删除失败")
			return
		}
		ok(w, nil)
	}
}

func AdminUpdateCommentStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCommentStatusReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewCommentLogic(svcCtx).UpdateStatus(r.Context(), &req)
		if err != nil {
			fail(w, 500, "更新状态失败")
			return
		}
		ok(w, nil)
	}
}

// --- 友链管理 ---

func AdminListFriendLinksHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewFriendLogic(svcCtx).List(r.Context())
		if err != nil {
			fail(w, 500, "获取友链失败")
			return
		}
		ok(w, data)
	}
}

func AdminCreateFriendLinkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateFriendLinkReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := admin.NewFriendLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			fail(w, 500, "创建友链失败")
			return
		}
		ok(w, data)
	}
}

func AdminUpdateFriendLinkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateFriendLinkReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewFriendLogic(svcCtx).Update(r.Context(), &req)
		if err != nil {
			fail(w, 500, "更新友链失败")
			return
		}
		ok(w, nil)
	}
}

func AdminDeleteFriendLinkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.ParsePath(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewFriendLogic(svcCtx).Delete(r.Context(), req.ID)
		if err != nil {
			fail(w, 500, "删除失败")
			return
		}
		ok(w, nil)
	}
}

// --- 留言管理 ---

func AdminListMessagesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageQuery
		if err := httpx.Parse(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := admin.NewMessageLogic(svcCtx).List(r.Context(), &req)
		if err != nil {
			fail(w, 500, "获取留言失败")
			return
		}
		ok(w, data)
	}
}

func AdminDeleteMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.ParsePath(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewMessageLogic(svcCtx).Delete(r.Context(), req.ID)
		if err != nil {
			fail(w, 500, "删除失败")
			return
		}
		ok(w, nil)
	}
}

// --- 页面管理 ---

func AdminListPagesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewPageLogic(svcCtx).List(r.Context())
		if err != nil {
			fail(w, 500, "获取页面失败")
			return
		}
		ok(w, data)
	}
}

func AdminCreatePageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreatePageReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := admin.NewPageLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			fail(w, 500, "创建页面失败")
			return
		}
		ok(w, data)
	}
}

func AdminUpdatePageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePageReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewPageLogic(svcCtx).Update(r.Context(), &req)
		if err != nil {
			fail(w, 500, "更新页面失败")
			return
		}
		ok(w, nil)
	}
}

func AdminDeletePageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.ParsePath(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewPageLogic(svcCtx).Delete(r.Context(), req.ID)
		if err != nil {
			fail(w, 500, "删除失败")
			return
		}
		ok(w, nil)
	}
}

// --- 站点配置管理 ---

func AdminGetSiteConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewSiteLogic(svcCtx).GetConfig(r.Context())
		if err != nil {
			fail(w, 500, "获取配置失败")
			return
		}
		ok(w, data)
	}
}

func AdminSetSiteConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetSiteConfigReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewSiteLogic(svcCtx).SetConfig(r.Context(), &req)
		if err != nil {
			fail(w, 500, "保存配置失败")
			return
		}
		ok(w, nil)
	}
}

// --- 角色管理 ---

func AdminListRolesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewRoleLogic(svcCtx).List(r.Context())
		if err != nil {
			fail(w, 500, "获取角色失败")
			return
		}
		ok(w, data)
	}
}

func AdminCreateRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRoleReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := admin.NewRoleLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			fail(w, 500, "创建角色失败")
			return
		}
		ok(w, data)
	}
}

func AdminUpdateRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRoleReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewRoleLogic(svcCtx).Update(r.Context(), &req)
		if err != nil {
			fail(w, 500, "更新角色失败")
			return
		}
		ok(w, nil)
	}
}

func AdminDeleteRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.ParsePath(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewRoleLogic(svcCtx).Delete(r.Context(), req.ID)
		if err != nil {
			fail(w, 500, "删除失败")
			return
		}
		ok(w, nil)
	}
}

func AdminUpdateRoleMenusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRoleMenusReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewRoleLogic(svcCtx).UpdateMenus(r.Context(), &req)
		if err != nil {
			fail(w, 500, "更新角色菜单失败")
			return
		}
		ok(w, nil)
	}
}

// --- 菜单管理 ---

func AdminListMenusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewMenuLogic(svcCtx).List(r.Context())
		if err != nil {
			fail(w, 500, "获取菜单失败")
			return
		}
		ok(w, data)
	}
}

func AdminCreateMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateMenuReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		data, err := admin.NewMenuLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			fail(w, 500, "创建菜单失败")
			return
		}
		ok(w, data)
	}
}

func AdminUpdateMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateMenuReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewMenuLogic(svcCtx).Update(r.Context(), &req)
		if err != nil {
			fail(w, 500, "更新菜单失败")
			return
		}
		ok(w, nil)
	}
}

func AdminDeleteMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.ParsePath(r, &req); err != nil {
			fail(w, 400, "参数错误")
			return
		}
		err := admin.NewMenuLogic(svcCtx).Delete(r.Context(), req.ID)
		if err != nil {
			fail(w, 500, "删除失败")
			return
		}
		ok(w, nil)
	}
}

// --- 仪表盘 ---

func AdminDashboardStatsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewDashboardLogic(svcCtx).GetStats(r.Context())
		if err != nil {
			fail(w, 500, "获取统计数据失败")
			return
		}
		ok(w, data)
	}
}

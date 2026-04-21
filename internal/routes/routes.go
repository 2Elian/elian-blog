package routes

import (
	"net/http"

	"elian-blog/internal/handler"
	"elian-blog/internal/middleware"
	"elian-blog/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, svcCtx *svc.ServiceContext) {
	// 前台 API
	registerBlogHandlers(server, svcCtx)
	// 后台 API
	registerAdminHandlers(server, svcCtx)
	// 静态文件
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/ping",
		Handler: pingHandler,
	})
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func registerBlogHandlers(server *rest.Server, svcCtx *svc.ServiceContext) {
	// 公开接口
	server.AddRoutes([]rest.Route{
		{Method: http.MethodPost, Path: "/blog-api/v1/login", Handler: handler.LoginHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/blog-api/v1/register", Handler: handler.RegisterHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/blog-api/v1/articles", Handler: handler.ListArticlesHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/blog-api/v1/articles/:id", Handler: handler.GetArticleHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/blog-api/v1/articles/search", Handler: handler.SearchArticlesHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/blog-api/v1/categories", Handler: handler.ListCategoriesHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/blog-api/v1/tags", Handler: handler.ListTagsHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/blog-api/v1/articles/:id/comments", Handler: handler.ListCommentsHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/blog-api/v1/comments/recent", Handler: handler.RecentCommentsHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/blog-api/v1/friend-links", Handler: handler.ListFriendLinksHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/blog-api/v1/messages", Handler: handler.ListMessagesHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/blog-api/v1/pages", Handler: handler.ListPagesHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/blog-api/v1/pages/:slug", Handler: handler.GetPageBySlugHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/blog-api/v1/site/config", Handler: handler.GetSiteConfigHandler(svcCtx)},
	})

	// 需要登录的接口
	authRoutes := []rest.Route{
		{Method: http.MethodGet, Path: "/blog-api/v1/user/info", Handler: handler.GetUserInfoHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/blog-api/v1/comments", Handler: handler.CreateCommentHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/blog-api/v1/messages", Handler: handler.CreateMessageHandler(svcCtx)},
	}
	server.AddRoutes(
		rest.WithMiddlewares([]rest.Middleware{middleware.JWTAuthMiddleware(svcCtx)}, authRoutes...),
	)
}

func registerAdminHandlers(server *rest.Server, svcCtx *svc.ServiceContext) {
	// 公开接口 - 登录
	server.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/admin-api/v1/login",
		Handler: handler.LoginHandler(svcCtx),
	})

	// 需要 JWT + RBAC 的接口
	adminRoutes := []rest.Route{
		{Method: http.MethodGet, Path: "/admin-api/v1/articles", Handler: handler.AdminListArticlesHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/admin-api/v1/articles/:id", Handler: handler.AdminGetArticleHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/articles", Handler: handler.AdminCreateArticleHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/articles/:id", Handler: handler.AdminUpdateArticleHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/articles/:id", Handler: handler.AdminDeleteArticleHandler(svcCtx)},

		{Method: http.MethodGet, Path: "/admin-api/v1/categories", Handler: handler.AdminListCategoriesHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/categories", Handler: handler.AdminCreateCategoryHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/categories/:id", Handler: handler.AdminUpdateCategoryHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/categories/:id", Handler: handler.AdminDeleteCategoryHandler(svcCtx)},

		{Method: http.MethodGet, Path: "/admin-api/v1/tags", Handler: handler.AdminListTagsHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/tags", Handler: handler.AdminCreateTagHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/tags/:id", Handler: handler.AdminUpdateTagHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/tags/:id", Handler: handler.AdminDeleteTagHandler(svcCtx)},

		{Method: http.MethodGet, Path: "/admin-api/v1/users", Handler: handler.AdminListUsersHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/users/:id", Handler: handler.AdminDeleteUserHandler(svcCtx)},

		{Method: http.MethodGet, Path: "/admin-api/v1/comments", Handler: handler.AdminListCommentsHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/comments/:id/status", Handler: handler.AdminUpdateCommentStatusHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/comments/:id", Handler: handler.AdminDeleteCommentHandler(svcCtx)},

		{Method: http.MethodGet, Path: "/admin-api/v1/friend-links", Handler: handler.AdminListFriendLinksHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/friend-links", Handler: handler.AdminCreateFriendLinkHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/friend-links/:id", Handler: handler.AdminUpdateFriendLinkHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/friend-links/:id", Handler: handler.AdminDeleteFriendLinkHandler(svcCtx)},

		{Method: http.MethodGet, Path: "/admin-api/v1/messages", Handler: handler.AdminListMessagesHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/messages/:id", Handler: handler.AdminDeleteMessageHandler(svcCtx)},

		{Method: http.MethodGet, Path: "/admin-api/v1/pages", Handler: handler.AdminListPagesHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/pages", Handler: handler.AdminCreatePageHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/pages/:id", Handler: handler.AdminUpdatePageHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/pages/:id", Handler: handler.AdminDeletePageHandler(svcCtx)},

		{Method: http.MethodGet, Path: "/admin-api/v1/site/config", Handler: handler.AdminGetSiteConfigHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/site/config", Handler: handler.AdminSetSiteConfigHandler(svcCtx)},

		{Method: http.MethodGet, Path: "/admin-api/v1/roles", Handler: handler.AdminListRolesHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/roles", Handler: handler.AdminCreateRoleHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/roles/:id", Handler: handler.AdminUpdateRoleHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/roles/:id", Handler: handler.AdminDeleteRoleHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/roles/:id/menus", Handler: handler.AdminUpdateRoleMenusHandler(svcCtx)},

		{Method: http.MethodGet, Path: "/admin-api/v1/menus", Handler: handler.AdminListMenusHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/menus", Handler: handler.AdminCreateMenuHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/menus/:id", Handler: handler.AdminUpdateMenuHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/menus/:id", Handler: handler.AdminDeleteMenuHandler(svcCtx)},
	}

	server.AddRoutes(
		rest.WithMiddlewares([]rest.Middleware{
			middleware.JWTAuthMiddleware(svcCtx),
			middleware.RBACAuthMiddleware(svcCtx),
		}, adminRoutes...),
	)
}
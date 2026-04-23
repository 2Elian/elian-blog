package routes

import (
	"net/http"

	"elian-blog/internal/handler"
	"elian-blog/internal/middleware"
	"elian-blog/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, svcCtx *svc.ServiceContext) {
	registerBlogHandlers(server, svcCtx)
	registerVeAdminHandlers(server, svcCtx)
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

	authRoutes := []rest.Route{
		{Method: http.MethodGet, Path: "/blog-api/v1/user/info", Handler: handler.GetUserInfoHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/blog-api/v1/comments", Handler: handler.CreateCommentHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/blog-api/v1/messages", Handler: handler.CreateMessageHandler(svcCtx)},
	}
	server.AddRoutes(
		rest.WithMiddlewares([]rest.Middleware{middleware.JWTAuthMiddleware(svcCtx)}, authRoutes...),
	)
}

// ve-admin-element 路由注册
func registerVeAdminHandlers(server *rest.Server, svcCtx *svc.ServiceContext) {
	// 公开接口
	server.AddRoutes([]rest.Route{
		{Method: http.MethodPost, Path: "/admin-api/v1/login", Handler: handler.VeLoginHandler(svcCtx)},
	})

	// 需要 JWT 认证的接口
	protectedRoutes := []rest.Route{
		// 认证
		{Method: http.MethodGet, Path: "/admin-api/v1/logout", Handler: handler.VeLogoutHandler(svcCtx)},

		// 用户信息
		{Method: http.MethodGet, Path: "/admin-api/v1/user/get_user_info", Handler: handler.VeGetUserInfoHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/admin-api/v1/user/get_user_menus", Handler: handler.VeGetUserMenusHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/admin-api/v1/user/get_user_roles", Handler: handler.VeGetUserRolesHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/admin-api/v1/user/get_user_apis", Handler: handler.VeGetUserApisHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/user/update_user_info", Handler: handler.VeUpdateUserInfoHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/user/update_user_avatar", Handler: handler.VeUpdateUserAvatarHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/user/update_user_password", Handler: handler.VeUpdateUserPasswordHandler(svcCtx)},

		// 文章
		{Method: http.MethodPost, Path: "/admin-api/v1/article/find_article_list", Handler: handler.VeFindArticleListHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/article/get_article", Handler: handler.VeGetArticleHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/article/add_article", Handler: handler.VeAddArticleHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/article/update_article", Handler: handler.VeUpdateArticleHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/article/delete_article", Handler: handler.VeDeleteArticleHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/article/update_article_top", Handler: handler.VeUpdateArticleTopHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/article/update_article_delete", Handler: handler.VeUpdateArticleDeleteHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/article/export_article_list", Handler: handler.VeExportArticleListHandler(svcCtx)},

		// 分类
		{Method: http.MethodPost, Path: "/admin-api/v1/category/find_category_list", Handler: handler.VeFindCategoryListHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/category/add_category", Handler: handler.VeAddCategoryHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/category/update_category", Handler: handler.VeUpdateCategoryHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/category/deletes_category", Handler: handler.VeDeletesCategoryHandler(svcCtx)},

		// 标签
		{Method: http.MethodPost, Path: "/admin-api/v1/tag/find_tag_list", Handler: handler.VeFindTagListHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/tag/add_tag", Handler: handler.VeAddTagHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/tag/update_tag", Handler: handler.VeUpdateTagHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/tag/deletes_tag", Handler: handler.VeDeletesTagHandler(svcCtx)},

		// 评论
		{Method: http.MethodPost, Path: "/admin-api/v1/comment/find_comment_back_list", Handler: handler.VeFindCommentBackListHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/comment/update_comment_status", Handler: handler.VeUpdateCommentStatusHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/comment/deletes_comment", Handler: handler.VeDeletesCommentHandler(svcCtx)},

		// 友链
		{Method: http.MethodPost, Path: "/admin-api/v1/friend/find_friend_list", Handler: handler.VeFindFriendListHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/friend/add_friend", Handler: handler.VeAddFriendHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/friend/update_friend", Handler: handler.VeUpdateFriendHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/friend/deletes_friend", Handler: handler.VeDeletesFriendHandler(svcCtx)},

		// 留言
		{Method: http.MethodPost, Path: "/admin-api/v1/message/find_message_list", Handler: handler.VeFindMessageListHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/message/update_message_status", Handler: handler.VeUpdateMessageStatusHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/message/deletes_message", Handler: handler.VeDeletesMessageHandler(svcCtx)},

		// 页面
		{Method: http.MethodPost, Path: "/admin-api/v1/page/find_page_list", Handler: handler.VeFindPageListHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/page/add_page", Handler: handler.VeAddPageHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/page/update_page", Handler: handler.VeUpdatePageHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/page/delete_page", Handler: handler.VeDeletePageHandler(svcCtx)},

		// 角色
		{Method: http.MethodPost, Path: "/admin-api/v1/role/find_role_list", Handler: handler.VeFindRoleListHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/role/add_role", Handler: handler.VeAddRoleHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/role/update_role", Handler: handler.VeUpdateRoleHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/role/deletes_role", Handler: handler.VeDeletesRoleHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/role/find_role_resources", Handler: handler.VeFindRoleResourcesHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/role/update_role_menus", Handler: handler.VeUpdateRoleMenusHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/role/update_role_apis", Handler: handler.VeUpdateRoleApisHandler(svcCtx)},

		// 菜单
		{Method: http.MethodPost, Path: "/admin-api/v1/menu/find_menu_list", Handler: handler.VeFindMenuListHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/menu/add_menu", Handler: handler.VeAddMenuHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/menu/update_menu", Handler: handler.VeUpdateMenuHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/menu/deletes_menu", Handler: handler.VeDeletesMenuHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/menu/sync_menu_list", Handler: handler.VeSyncMenuListHandler(svcCtx)},

		// 账号管理
		{Method: http.MethodPost, Path: "/admin-api/v1/account/find_account_list", Handler: handler.VeFindAccountListHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/account/find_account_online_list", Handler: handler.VeFindAccountOnlineListHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/account/update_account_status", Handler: handler.VeUpdateAccountStatusHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/account/update_account_roles", Handler: handler.VeUpdateAccountRolesHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/account/update_account_password", Handler: handler.VeUpdateAccountPasswordHandler(svcCtx)},

		// 首页/网站
		{Method: http.MethodGet, Path: "/admin-api/v1/admin", Handler: handler.VeGetAdminHomeInfoHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/admin-api/v1/admin/get_website_config", Handler: handler.VeGetWebsiteConfigHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/admin/update_website_config", Handler: handler.VeUpdateWebsiteConfigHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/admin-api/v1/admin/get_visit_stats", Handler: handler.VeGetVisitStatsHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/admin/get_visit_trend", Handler: handler.VeGetVisitTrendHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/admin-api/v1/admin/get_about_me", Handler: handler.VeGetAboutMeHandler(svcCtx)},
		{Method: http.MethodPut, Path: "/admin-api/v1/admin/update_about_me", Handler: handler.VeUpdateAboutMeHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/admin-api/v1/admin/get_system_state", Handler: handler.VeGetSystemStateHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/admin/get_user_area_stats", Handler: handler.VeGetUserAreaStatsHandler(svcCtx)},

		// 通知
			{Method: http.MethodPost, Path: "/admin-api/v1/notice/find_user_notice_list", Handler: handler.VeFindUserNoticeListHandler(svcCtx)},

			// 文件上传
		{Method: http.MethodPost, Path: "/admin-api/v1/upload/upload_file", Handler: handler.VeUploadFileHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/upload/multi_upload_file", Handler: handler.VeMultiUploadFileHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/admin-api/v1/upload/list_upload_file", Handler: handler.VeListUploadFileHandler(svcCtx)},
		{Method: http.MethodDelete, Path: "/admin-api/v1/upload/deletes_upload_file", Handler: handler.VeDeletesUploadFileHandler(svcCtx)},
	}

	server.AddRoutes(
		rest.WithMiddlewares([]rest.Middleware{
			middleware.JWTAuthMiddleware(svcCtx),
		}, protectedRoutes...),
	)
}

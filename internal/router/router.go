package router

import (
	"elian-blog/internal/controller/admin"
	blogctrl "elian-blog/internal/controller/blog"
	"elian-blog/internal/dao"
	"elian-blog/internal/middleware"
	"elian-blog/internal/model"
	"elian-blog/internal/service"
	"elian-blog/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, rdb *redis.Client, cfg *config.Config, log *zap.Logger) *gin.Engine {
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	_ = model.AutoMigrate(db)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())

	r.Static("/uploads", cfg.Upload.Path)

	userDao := dao.NewUserDao(db)
	articleDao := dao.NewArticleDao(db)
	categoryDao := dao.NewCategoryDao(db)
	tagDao := dao.NewTagDao(db)
	commentDao := dao.NewCommentDao(db)
	friendDao := dao.NewFriendLinkDao(db)
	messageDao := dao.NewMessageDao(db)
	pageDao := dao.NewPageDao(db)
	siteDao := dao.NewSiteConfigDao(db)
	roleDao := dao.NewRoleDao(db)
	menuDao := dao.NewMenuDao(db)
	opLogDao := dao.NewOperationLogDao(db)

	authSvc := service.NewAuthService(userDao, cfg.JWT.Secret, cfg.JWT.ExpireHours, log)
	articleSvc := service.NewArticleService(articleDao, categoryDao, tagDao)
	userSvc := service.NewUserService(userDao)
	categorySvc := service.NewCategoryService(categoryDao)
	tagSvc := service.NewTagService(tagDao)
	commentSvc := service.NewCommentService(commentDao)
	friendSvc := service.NewFriendLinkService(friendDao)
	messageSvc := service.NewMessageService(messageDao)
	pageSvc := service.NewPageService(pageDao)
	siteSvc := service.NewSiteConfigService(siteDao)
	roleSvc := service.NewRoleService(roleDao, menuDao)
	menuSvc := service.NewMenuService(menuDao)

	setupBlogRoutes(r, authSvc, articleSvc, categorySvc, tagSvc, commentSvc, friendSvc, messageSvc, pageSvc, siteSvc, cfg)
	setupAdminRoutes(r, authSvc, articleSvc, userSvc, categorySvc, tagSvc, commentSvc, friendSvc, messageSvc, pageSvc, siteSvc, roleSvc, menuSvc, opLogDao, cfg)

	return r
}

func setupBlogRoutes(r *gin.Engine, authSvc *service.AuthService, articleSvc *service.ArticleService,
	categorySvc *service.CategoryService, tagSvc *service.TagService, commentSvc *service.CommentService,
	friendSvc *service.FriendLinkService, messageSvc *service.MessageService, pageSvc *service.PageService,
	siteSvc *service.SiteConfigService, cfg *config.Config) {

	api := r.Group("/blog-api/v1")

	authCtrl := blogctrl.NewAuthController(authSvc)
	articleCtrl := blogctrl.NewArticleController(articleSvc)
	categoryCtrl := blogctrl.NewCategoryController(categorySvc)
	tagCtrl := blogctrl.NewTagController(tagSvc)
	commentCtrl := blogctrl.NewCommentController(commentSvc)
	friendCtrl := blogctrl.NewFriendLinkController(friendSvc)
	messageCtrl := blogctrl.NewMessageController(messageSvc)
	pageCtrl := blogctrl.NewPageController(pageSvc)
	siteCtrl := blogctrl.NewSiteController(siteSvc)

	api.POST("/login", authCtrl.Login)
	api.POST("/register", authCtrl.Register)

	api.GET("/articles", articleCtrl.List)
	api.GET("/articles/:id", articleCtrl.Get)
	api.GET("/articles/search", articleCtrl.Search)

	api.GET("/categories", categoryCtrl.List)
	api.GET("/tags", tagCtrl.List)

	api.GET("/articles/:id/comments", commentCtrl.ListByArticle)
	api.GET("/comments/recent", commentCtrl.Recent)

	api.GET("/friend-links", friendCtrl.List)
	api.GET("/messages", messageCtrl.List)
	api.GET("/pages", pageCtrl.List)
	api.GET("/pages/:slug", pageCtrl.GetBySlug)
	api.GET("/site/config", siteCtrl.GetConfig)

	authGroup := api.Group("")
	authGroup.Use(middleware.JWTAuth(cfg.JWT.Secret))
	{
		authGroup.GET("/user/info", authCtrl.GetUserInfo)
		authGroup.POST("/comments", commentCtrl.Create)
		authGroup.POST("/messages", messageCtrl.Create)
	}
}

func setupAdminRoutes(r *gin.Engine, authSvc *service.AuthService, articleSvc *service.ArticleService,
	userSvc *service.UserService, categorySvc *service.CategoryService, tagSvc *service.TagService,
	commentSvc *service.CommentService, friendSvc *service.FriendLinkService, messageSvc *service.MessageService,
	pageSvc *service.PageService, siteSvc *service.SiteConfigService, roleSvc *service.RoleService,
	menuSvc *service.MenuService, opLogDao *dao.OperationLogDao, cfg *config.Config) {

	authCtrl := blogctrl.NewAuthController(authSvc)
	r.POST("/admin-api/v1/login", authCtrl.Login)

	adminGroup := r.Group("/admin-api/v1")
	adminGroup.Use(middleware.JWTAuth(cfg.JWT.Secret))
	adminGroup.Use(middleware.RBACAuth())
	adminGroup.Use(middleware.OperationLog(opLogDao))

	articleCtrl := admin.NewArticleController(articleSvc)
	adminGroup.GET("/articles", articleCtrl.List)
	adminGroup.GET("/articles/:id", articleCtrl.Get)
	adminGroup.POST("/articles", articleCtrl.Create)
	adminGroup.PUT("/articles/:id", articleCtrl.Update)
	adminGroup.DELETE("/articles/:id", articleCtrl.Delete)

	categoryCtrl := admin.NewCategoryController(categorySvc)
	adminGroup.GET("/categories", categoryCtrl.List)
	adminGroup.POST("/categories", categoryCtrl.Create)
	adminGroup.PUT("/categories/:id", categoryCtrl.Update)
	adminGroup.DELETE("/categories/:id", categoryCtrl.Delete)

	tagCtrl := admin.NewTagController(tagSvc)
	adminGroup.GET("/tags", tagCtrl.List)
	adminGroup.POST("/tags", tagCtrl.Create)
	adminGroup.PUT("/tags/:id", tagCtrl.Update)
	adminGroup.DELETE("/tags/:id", tagCtrl.Delete)

	userCtrl := admin.NewUserController(userSvc)
	adminGroup.GET("/users", userCtrl.List)
	adminGroup.GET("/users/:id", userCtrl.Get)
	adminGroup.DELETE("/users/:id", userCtrl.Delete)

	commentCtrl := admin.NewCommentController(commentSvc)
	adminGroup.GET("/comments", commentCtrl.List)
	adminGroup.PUT("/comments/:id/status", commentCtrl.UpdateStatus)
	adminGroup.DELETE("/comments/:id", commentCtrl.Delete)

	friendCtrl := admin.NewFriendLinkController(friendSvc)
	adminGroup.GET("/friend-links", friendCtrl.List)
	adminGroup.POST("/friend-links", friendCtrl.Create)
	adminGroup.PUT("/friend-links/:id", friendCtrl.Update)
	adminGroup.DELETE("/friend-links/:id", friendCtrl.Delete)

	messageCtrl := admin.NewMessageController(messageSvc)
	adminGroup.GET("/messages", messageCtrl.List)
	adminGroup.DELETE("/messages/:id", messageCtrl.Delete)

	pageCtrl := admin.NewPageController(pageSvc)
	adminGroup.GET("/pages", pageCtrl.List)
	adminGroup.POST("/pages", pageCtrl.Create)
	adminGroup.PUT("/pages/:id", pageCtrl.Update)
	adminGroup.DELETE("/pages/:id", pageCtrl.Delete)

	siteCtrl := admin.NewSiteConfigController(siteSvc)
	adminGroup.GET("/site/config", siteCtrl.Get)
	adminGroup.PUT("/site/config", siteCtrl.Set)

	roleCtrl := admin.NewRoleController(roleSvc, menuSvc)
	adminGroup.GET("/roles", roleCtrl.List)
	adminGroup.POST("/roles", roleCtrl.Create)
	adminGroup.PUT("/roles/:id", roleCtrl.Update)
	adminGroup.DELETE("/roles/:id", roleCtrl.Delete)
	adminGroup.PUT("/roles/:id/menus", roleCtrl.UpdateMenus)

	menuCtrl := admin.NewMenuController(menuSvc)
	adminGroup.GET("/menus", menuCtrl.List)
	adminGroup.POST("/menus", menuCtrl.Create)
	adminGroup.PUT("/menus/:id", menuCtrl.Update)
	adminGroup.DELETE("/menus/:id", menuCtrl.Delete)
}

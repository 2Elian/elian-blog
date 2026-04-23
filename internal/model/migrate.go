package model

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Role{},
		&Menu{},
		&Article{},
		&Category{},
		&Tag{},
		&Comment{},
		&FriendLink{},
		&Message{},
		&Page{},
		&SiteConfig{},
		&OperationLog{},
	)
}

// SeedMenus inserts default menu data if the menu table is empty.
func SeedMenus(db *gorm.DB) error {
	var count int64
	db.Model(&Menu{}).Count(&count)
	if count > 0 {
		return nil
	}

	menus := []Menu{
		// 仪表盘
		{Name: "Dashboard", Title: "仪表盘", Path: "/dashboard", Component: "views/admin/home/Home.vue", Icon: "ep:home-filled", ParentID: 0, Sort: 1, Type: 1, KeepAlive: 1},

		// 文章管理（目录）
		{Name: "ArticleMgmt", Title: "文章管理", Path: "/article", Icon: "ep:document", ParentID: 0, Sort: 2, Type: 0, AlwaysShow: 1},
		{Name: "ArticleList", Title: "文章列表", Path: "/article/list", Component: "views/admin/article/Article.vue", ParentID: 0, Sort: 1, Type: 1, KeepAlive: 1},
		{Name: "ArticleWrite", Title: "写文章", Path: "/article/write", Component: "views/admin/article/ArticleEditor.vue", ParentID: 0, Sort: 2, Type: 1, KeepAlive: 1},

		// 分类管理
		{Name: "CategoryMgmt", Title: "分类管理", Path: "/category", Component: "views/admin/category/Category.vue", Icon: "ep:menu", ParentID: 0, Sort: 3, Type: 1, KeepAlive: 1},

		// 标签管理
		{Name: "TagMgmt", Title: "标签管理", Path: "/tag", Component: "views/admin/tag/Tag.vue", Icon: "ep:price-tag", ParentID: 0, Sort: 4, Type: 1, KeepAlive: 1},

		// 评论管理
		{Name: "CommentMgmt", Title: "评论管理", Path: "/comment", Component: "views/admin/comment/Comment.vue", Icon: "ep:chat-dot-square", ParentID: 0, Sort: 5, Type: 1, KeepAlive: 1},

		// 友链管理
		{Name: "FriendMgmt", Title: "友链管理", Path: "/friend", Component: "views/admin/friend/Friend.vue", Icon: "ep:link", ParentID: 0, Sort: 6, Type: 1, KeepAlive: 1},

		// 留言管理
		{Name: "MessageMgmt", Title: "留言管理", Path: "/message", Component: "views/admin/message/Message.vue", Icon: "ep:message", ParentID: 0, Sort: 7, Type: 1, KeepAlive: 1},

		// 页面管理
		{Name: "PageMgmt", Title: "页面管理", Path: "/page", Component: "views/admin/page/Page.vue", Icon: "ep:document-copy", ParentID: 0, Sort: 8, Type: 1, KeepAlive: 1},

		// 系统管理（目录）
		{Name: "System", Title: "系统管理", Path: "/system", Icon: "ep:setting", ParentID: 0, Sort: 9, Type: 0, AlwaysShow: 1},
		{Name: "UserMgmt", Title: "用户管理", Path: "/system/user", Component: "views/admin/system/user/User.vue", ParentID: 0, Sort: 1, Type: 1, KeepAlive: 1},
		{Name: "RoleMgmt", Title: "角色管理", Path: "/system/role", Component: "views/admin/system/role/Role.vue", ParentID: 0, Sort: 2, Type: 1, KeepAlive: 1},
		{Name: "MenuMgmt", Title: "菜单管理", Path: "/system/menu", Component: "views/admin/system/menu/Menu.vue", ParentID: 0, Sort: 3, Type: 1, KeepAlive: 1},

		// 站点设置
		{Name: "SiteConfig", Title: "站点设置", Path: "/site", Component: "views/admin/site/SiteConfig.vue", Icon: "ep:tools", ParentID: 0, Sort: 10, Type: 1, KeepAlive: 1},

		// 个人中心
		{Name: "Profile", Title: "个人中心", Path: "/profile", Component: "views/admin/profile/Profile.vue", Icon: "ep:user", ParentID: 0, Sort: 11, Type: 1, IsHidden: 1},
	}

	return db.Create(&menus).Error
}

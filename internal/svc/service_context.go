package svc

import (
	"elian-blog/internal/config"
	"elian-blog/internal/dao"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
	Log    *zap.Logger

	// DAO
	UserDao         *dao.UserDao
	ArticleDao      *dao.ArticleDao
	CategoryDao     *dao.CategoryDao
	TagDao          *dao.TagDao
	CommentDao      *dao.CommentDao
	FriendDao       *dao.FriendLinkDao
	MessageDao      *dao.MessageDao
	PageDao         *dao.PageDao
	SiteDao         *dao.SiteConfigDao
	RoleDao         *dao.RoleDao
	MenuDao         *dao.MenuDao
	OperationLogDao *dao.OperationLogDao
}

func NewServiceContext(c config.Config, db *gorm.DB, rdb *redis.Client, log *zap.Logger) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		DB:              db,
		RDB:             rdb,
		Log:             log,
		UserDao:         dao.NewUserDao(db),
		ArticleDao:      dao.NewArticleDao(db),
		CategoryDao:     dao.NewCategoryDao(db),
		TagDao:          dao.NewTagDao(db),
		CommentDao:      dao.NewCommentDao(db),
		FriendDao:       dao.NewFriendLinkDao(db),
		MessageDao:      dao.NewMessageDao(db),
		PageDao:         dao.NewPageDao(db),
		SiteDao:         dao.NewSiteConfigDao(db),
		RoleDao:         dao.NewRoleDao(db),
		MenuDao:         dao.NewMenuDao(db),
		OperationLogDao: dao.NewOperationLogDao(db),
	}
}

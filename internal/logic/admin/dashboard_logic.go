package admin

import (
	"context"

	"elian-blog/internal/model"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type DashboardLogic struct {
	svcCtx *svc.ServiceContext
}

func NewDashboardLogic(svcCtx *svc.ServiceContext) *DashboardLogic {
	return &DashboardLogic{svcCtx: svcCtx}
}

func (l *DashboardLogic) GetStats(ctx context.Context) (interface{}, error) {
	var articleCount, userCount, commentCount, viewCount int64

	// 文章总数
	l.svcCtx.DB.Model(&model.Article{}).Count(&articleCount)

	// 用户总数
	l.svcCtx.DB.Model(&model.User{}).Count(&userCount)

	// 评论总数
	l.svcCtx.DB.Model(&model.Comment{}).Count(&commentCount)

	// 访问总量（文章浏览数总和）
	l.svcCtx.DB.Model(&model.Article{}).Select("SUM(views)").Scan(&viewCount)

	return types.DashboardStats{
		ArticleCount: int(articleCount),
		UserCount:    int(userCount),
		CommentCount: int(commentCount),
		ViewCount:    int(viewCount),
	}, nil
}

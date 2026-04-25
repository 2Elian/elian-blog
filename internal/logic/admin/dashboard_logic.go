package admin

import (
	"context"

	"elian-blog/internal/model"
	"elian-blog/internal/svc"
)


type DashboardLogic struct {
	svcCtx *svc.ServiceContext
}

func NewDashboardLogic(svcCtx *svc.ServiceContext) *DashboardLogic {
	return &DashboardLogic{svcCtx: svcCtx}
}

type CategoryVO struct {
	ID            uint   `json:"id" gorm:"column:id"`
	CategoryName  string `json:"category_name" gorm:"column:name"`
	ArticleCount  int    `json:"article_count" gorm:"-"`
}

type TagVO struct {
	ID           uint   `json:"id" gorm:"column:id"`
	TagName      string `json:"tag_name" gorm:"column:name"`
	ArticleCount int    `json:"article_count" gorm:"-"`
}

type ArticleViewVO struct {
	ID            uint   `json:"id"`
	ArticleTitle  string `json:"article_title"`
	ViewCount     int    `json:"view_count"`
}

type ArticleStatisticsVO struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type AdminHomeInfoResp struct {
	UserCount         int                   `json:"user_count"`
	ArticleCount      int                   `json:"article_count"`
	MessageCount      int                   `json:"message_count"`
	ProductCount      int                   `json:"product_count"`
	CategoryList      []CategoryVO          `json:"category_list"`
	TagList           []TagVO               `json:"tag_list"`
	ArticleViewRanks  []ArticleViewVO       `json:"article_view_ranks"`
	ArticleStatistics []ArticleStatisticsVO `json:"article_statistics"`
}

type VisitTrendItem struct {
	Date    string `json:"date"`
	PVCount int    `json:"pv_count"`
	UVCount int    `json:"uv_count"`
}

type UserAreaVO struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func (l *DashboardLogic) GetStats(ctx context.Context) (interface{}, error) {
	var articleCount, userCount, commentCount, messageCount, productCount, viewCount int64

	// 文章总数
	l.svcCtx.DB.Model(&model.Article{}).Count(&articleCount)

	// 用户总数
	l.svcCtx.DB.Model(&model.User{}).Count(&userCount)

	// 评论总数
	l.svcCtx.DB.Model(&model.Comment{}).Count(&commentCount)

	// 留言总数
	l.svcCtx.DB.Model(&model.Message{}).Count(&messageCount)

	// 产品总数
	l.svcCtx.DB.Model(&model.Product{}).Count(&productCount)

	// 访问总量（文章浏览数总和）
	l.svcCtx.DB.Model(&model.Article{}).Select("COALESCE(SUM(view_count), 0)").Scan(&viewCount)

	// 分类列表
	var categoryList []CategoryVO
	l.svcCtx.DB.Model(&model.Category{}).Find(&categoryList)
	for i := range categoryList {
		var count int64
		l.svcCtx.DB.Model(&model.Article{}).Where("category_id = ?", categoryList[i].ID).Count(&count)
		categoryList[i].ArticleCount = int(count)
	}

	// 标签列表
	var tagList []TagVO
	l.svcCtx.DB.Model(&model.Tag{}).Find(&tagList)
	for i := range tagList {
		var count int64
		l.svcCtx.DB.Table("article_tags").Where("tag_id = ?", tagList[i].ID).Count(&count)
		tagList[i].ArticleCount = int(count)
	}

	// 文章浏览量排行 (前10)
	var articleViewRanks []ArticleViewVO
	l.svcCtx.DB.Model(&model.Article{}).Select("id, title as article_title, view_count").Order("view_count DESC").Limit(10).Find(&articleViewRanks)

	// 文章提交统计 (近30天)
	var articleStatistics []ArticleStatisticsVO
	l.svcCtx.DB.Raw(`
		SELECT DATE(created_at) as date, COUNT(*) as count
		FROM article
		WHERE created_at >= DATE_SUB(NOW(), INTERVAL 30 DAY) AND deleted_at IS NULL
		GROUP BY DATE(created_at)
		ORDER BY date ASC
	`).Scan(&articleStatistics)

	return AdminHomeInfoResp{
		UserCount:         int(userCount),
		ArticleCount:      int(articleCount),
		MessageCount:      int(messageCount),
		ProductCount:      int(productCount),
		CategoryList:      categoryList,
		TagList:           tagList,
		ArticleViewRanks:  articleViewRanks,
		ArticleStatistics: articleStatistics,
	}, nil
}

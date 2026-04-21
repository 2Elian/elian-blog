package blog

import (
	"context"
	"errors"
	"time"

	"elian-blog/internal/model"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"

	"gorm.io/gorm"
)

type ArticleLogic struct {
	svcCtx *svc.ServiceContext
}

func NewArticleLogic(svcCtx *svc.ServiceContext) *ArticleLogic {
	return &ArticleLogic{svcCtx: svcCtx}
}

// ListArticles 获取文章列表
func (l *ArticleLogic) ListArticles(ctx context.Context, req *types.QueryArticleHomeReq) (list interface{}, total int64, err error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	// 默认查询已发布的文章
	status := req.Status
	if status < 0 {
		status = 1 // 默认只显示已发布
	}

	articles, total, err := l.svcCtx.ArticleDao.List(page, pageSize, status, req.CategoryID, req.TagID)
	if err != nil {
		return nil, 0, err
	}

	// 转换为 ArticleHome 列表
	list = l.convertToArticleHomeList(articles)
	return list, total, nil
}

// GetArticle 获取文章详情
func (l *ArticleLogic) GetArticle(ctx context.Context, id uint) (interface{}, error) {
	article, err := l.svcCtx.ArticleDao.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("文章不存在")
		}
		return nil, err
	}

	// 增加浏览量
	go l.svcCtx.ArticleDao.IncrementViewCount(id)

	// 转换为 ArticleDetails
	details := l.convertToArticleDetails(article)
	return details, nil
}

// convertToArticleHomeList 转换文章列表
func (l *ArticleLogic) convertToArticleHomeList(articles []model.Article) []types.ArticleHome {
	result := make([]types.ArticleHome, 0, len(articles))
	for _, article := range articles {
		item := types.ArticleHome{
			ID:           article.ID,
			Title:        article.Title,
			Summary:      article.Summary,
			Content:      article.Content,
			Cover:        article.Cover,
			CategoryID:   article.CategoryID,
			CategoryName: article.Category.Name,
			AuthorID:     article.AuthorID,
			Status:       article.Status,
			IsTop:        article.IsTop,
			IsOriginal:   article.IsOriginal,
			Type:         article.Type,
			ViewCount:    article.ViewCount,
			LikeCount:    article.LikeCount,
			TagNameList:  l.getTagNames(article.Tags),
			CreatedAt:    article.CreatedAt.Format(time.DateTime),
			UpdatedAt:    article.UpdatedAt.Format(time.DateTime),
		}
		result = append(result, item)
	}
	return result
}

// convertToArticleDetails 转换文章详情
func (l *ArticleLogic) convertToArticleDetails(article *model.Article) *types.ArticleDetails {
	details := &types.ArticleDetails{
		ArticleHome: types.ArticleHome{
			ID:           article.ID,
			Title:        article.Title,
			Summary:      article.Summary,
			Content:      article.Content,
			Cover:        article.Cover,
			CategoryID:   article.CategoryID,
			CategoryName: article.Category.Name,
			AuthorID:     article.AuthorID,
			Status:       article.Status,
			IsTop:        article.IsTop,
			IsOriginal:   article.IsOriginal,
			Type:         article.Type,
			ViewCount:    article.ViewCount,
			LikeCount:    article.LikeCount,
			TagNameList:  l.getTagNames(article.Tags),
			CreatedAt:    article.CreatedAt.Format(time.DateTime),
			UpdatedAt:    article.UpdatedAt.Format(time.DateTime),
		},
		PrevArticle: nil,
		NextArticle: nil,
	}
	return details
}

// getTagNames 获取标签名称列表
func (l *ArticleLogic) getTagNames(tags []model.Tag) []string {
	names := make([]string, 0, len(tags))
	for _, tag := range tags {
		names = append(names, tag.Name)
	}
	return names
}

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

func (l *ArticleLogic) ListArticles(ctx context.Context, req *types.QueryArticleHomeReq) (list interface{}, total int64, err error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	status := req.Status
	if status < 0 {
		status = 1
	}

	articles, total, err := l.svcCtx.ArticleDao.List(page, pageSize, status, req.CategoryID, req.TagID, 0)
	if err != nil {
		return nil, 0, err
	}

	list = l.convertToArticleHomeList(articles)
	return list, total, nil
}

func (l *ArticleLogic) GetArticle(ctx context.Context, id uint) (interface{}, error) {
	article, err := l.svcCtx.ArticleDao.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("文章不存在")
		}
		return nil, err
	}

	go l.svcCtx.ArticleDao.IncrementViewCount(id)

	details := l.convertToArticleDetails(article)
	return details, nil
}

func (l *ArticleLogic) convertToArticleHomeList(articles []model.Article) []types.ArticleHome {
	result := make([]types.ArticleHome, 0, len(articles))
	for _, article := range articles {
		result = append(result, l.articleToHome(&article))
	}
	return result
}

func (l *ArticleLogic) convertToArticleDetails(article *model.Article) *types.ArticleDetails {
	return &types.ArticleDetails{
		ArticleHome: l.articleToHome(article),
		PrevArticle: nil,
		NextArticle: nil,
	}
}

func (l *ArticleLogic) articleToHome(article *model.Article) types.ArticleHome {
	var catID uint
	if article.CategoryID != nil {
		catID = *article.CategoryID
	}
	return types.ArticleHome{
		ID:           article.ID,
		Title:        article.Title,
		Summary:      article.Summary,
		Content:      article.Content,
		Cover:        article.Cover,
		CategoryID:   catID,
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
}

func (l *ArticleLogic) getTagNames(tags []model.Tag) []string {
	names := make([]string, 0, len(tags))
	for _, tag := range tags {
		names = append(names, tag.Name)
	}
	return names
}

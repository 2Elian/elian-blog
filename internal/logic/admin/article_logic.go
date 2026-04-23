package admin

import (
	"context"

	"elian-blog/internal/model"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type ArticleLogic struct {
	svcCtx *svc.ServiceContext
}

func NewArticleLogic(svcCtx *svc.ServiceContext) *ArticleLogic {
	return &ArticleLogic{svcCtx: svcCtx}
}

func (l *ArticleLogic) Create(ctx context.Context, req *types.CreateArticleReq) (interface{}, error) {
	article := &model.Article{
		Title:      req.Title,
		Summary:    req.Summary,
		Content:    req.Content,
		Cover:      req.Cover,
		CategoryID: req.CategoryID,
		Status:     req.Status,
		IsTop:      req.IsTop,
		IsOriginal: req.IsOriginal,
		Type:       req.Type,
		Password:   req.Password,
	}

	if err := l.svcCtx.ArticleDao.Create(article); err != nil {
		return nil, err
	}

	// Handle tags
	if len(req.TagNames) > 0 {
		tags, err := l.findOrCreateTags(req.TagNames)
		if err != nil {
			return nil, err
		}
		if err := l.svcCtx.ArticleDao.UpdateTags(article, tags); err != nil {
			return nil, err
		}
	}

	// Reload article with associations
	return l.svcCtx.ArticleDao.GetByID(article.ID)
}

func (l *ArticleLogic) Update(ctx context.Context, req *types.UpdateArticleReq) error {
	article, err := l.svcCtx.ArticleDao.GetByID(req.ID)
	if err != nil {
		return err
	}

	// Update fields
	if req.Title != "" {
		article.Title = req.Title
	}
	if req.Summary != "" {
		article.Summary = req.Summary
	}
	if req.Content != "" {
		article.Content = req.Content
	}
	if req.Cover != "" {
		article.Cover = req.Cover
	}
	if req.CategoryID != 0 {
		article.CategoryID = req.CategoryID
	}
	if req.Status != 0 {
		article.Status = req.Status
	}
	if req.IsTop != 0 {
		article.IsTop = req.IsTop
	}
	if req.IsOriginal != 0 {
		article.IsOriginal = req.IsOriginal
	}
	if req.Type != 0 {
		article.Type = req.Type
	}
	if req.Password != "" {
		article.Password = req.Password
	}

	if err := l.svcCtx.ArticleDao.Update(article); err != nil {
		return err
	}

	// Update tags if provided
	if req.TagNames != nil {
		tags, err := l.findOrCreateTags(req.TagNames)
		if err != nil {
			return err
		}
		if err := l.svcCtx.ArticleDao.UpdateTags(article, tags); err != nil {
			return err
		}
	}

	return nil
}

func (l *ArticleLogic) Delete(ctx context.Context, id uint) error {
	return l.svcCtx.ArticleDao.Delete(id)
}

func (l *ArticleLogic) List(ctx context.Context, req *types.QueryArticleHomeReq) (list interface{}, total int64, err error) {
	page := req.Page
	pageSize := req.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	articles, total, err := l.svcCtx.ArticleDao.List(page, pageSize, req.Status, req.CategoryID, req.TagID)
	if err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

func (l *ArticleLogic) Get(ctx context.Context, id uint) (interface{}, error) {
	return l.svcCtx.ArticleDao.GetByID(id)
}

// findOrCreateTags resolves tag names to Tag models, creating them if necessary.
func (l *ArticleLogic) findOrCreateTags(names []string) ([]model.Tag, error) {
	tags := make([]model.Tag, 0, len(names))
	for _, name := range names {
		tag, err := l.svcCtx.TagDao.FindOrCreate(name)
		if err != nil {
			return nil, err
		}
		tags = append(tags, *tag)
	}
	return tags, nil
}

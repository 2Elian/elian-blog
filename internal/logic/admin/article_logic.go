package admin

import (
	"context"
	"strings"

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

func (l *ArticleLogic) Create(ctx context.Context, req *types.CreateArticleReq, userID uint) (interface{}, error) {
	article := &model.Article{
		Title:     req.ArticleTitle,
		Content:   req.ArticleContent,
		Cover:     req.ArticleCover,
		Type:      req.ArticleType,
		SourceURL: req.OriginalURL,
		Status:    req.Status,
		IsTop:     req.IsTop,
		AuthorID:  userID,
	}

	// Resolve category by name
	if req.CategoryName != "" {
		var cat model.Category
		l.svcCtx.DB.Where("name = ?", req.CategoryName).First(&cat)
		if cat.ID > 0 {
			article.CategoryID = &cat.ID
		} else {
			// Create category if not exists
			cat = model.Category{Name: req.CategoryName}
			l.svcCtx.CategoryDao.Create(&cat)
			article.CategoryID = &cat.ID
		}
	}

	if err := l.svcCtx.ArticleDao.Create(article); err != nil {
		return nil, err
	}

	// Handle tags
	if len(req.TagNameList) > 0 {
		tags, err := l.findOrCreateTags(req.TagNameList)
		if err != nil {
			return nil, err
		}
		if err := l.svcCtx.ArticleDao.UpdateTags(article, tags); err != nil {
			return nil, err
		}
	}

	// Reload article with associations
	loaded, err := l.svcCtx.ArticleDao.GetByID(article.ID)
	if err != nil {
		return nil, err
	}
	return toArticleBackVO(loaded), nil
}

func (l *ArticleLogic) Update(ctx context.Context, req *types.UpdateArticleReq) error {
	article, err := l.svcCtx.ArticleDao.GetByID(req.ID)
	if err != nil {
		return err
	}

	if req.ArticleTitle != "" {
		article.Title = req.ArticleTitle
	}
	if req.ArticleContent != "" {
		article.Content = req.ArticleContent
	}
	if req.ArticleCover != "" {
		article.Cover = req.ArticleCover
	}
	if req.ArticleType != 0 {
		article.Type = req.ArticleType
	}
	if req.OriginalURL != "" {
		article.SourceURL = req.OriginalURL
	}
	if req.Status != 0 {
		article.Status = req.Status
	}
	if req.IsTop != 0 {
		article.IsTop = req.IsTop
	}
	if req.CategoryName != "" {
		var cat model.Category
		l.svcCtx.DB.Where("name = ?", req.CategoryName).First(&cat)
		if cat.ID > 0 {
			article.CategoryID = &cat.ID
		}
	}

	if err := l.svcCtx.ArticleDao.Update(article); err != nil {
		return err
	}

	// Update tags if provided
	if req.TagNameList != nil {
		tags, err := l.findOrCreateTags(req.TagNameList)
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

	// Handle category_name filter
	categoryID := req.CategoryID
	if req.CategoryName != "" && categoryID == 0 {
		var cat model.Category
		l.svcCtx.DB.Where("name = ?", req.CategoryName).First(&cat)
		if cat.ID > 0 {
			categoryID = cat.ID
		}
	}

	articles, total, err := l.svcCtx.ArticleDao.List(page, pageSize, req.Status, categoryID, req.TagID, req.IsDelete)
	if err != nil {
		return nil, 0, err
	}

	// Convert to ArticleBackVO
	voList := make([]types.ArticleBackVO, 0, len(articles))
	for _, a := range articles {
		voList = append(voList, toArticleBackVO(&a))
	}

	return voList, total, nil
}

func (l *ArticleLogic) Get(ctx context.Context, id uint) (interface{}, error) {
	article, err := l.svcCtx.ArticleDao.GetByID(id)
	if err != nil {
		return nil, err
	}
	return toArticleBackVO(article), nil
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

func toArticleBackVO(a *model.Article) types.ArticleBackVO {
	cover := a.Cover
	if cover != "" && !strings.HasPrefix(cover, "http") {
		cover = "http://localhost:8080" + cover
	}

	vo := types.ArticleBackVO{
		ID:             a.ID,
		ArticleTitle:   a.Title,
		ArticleContent: a.Content,
		ArticleCover:   cover,
		ArticleType:    a.Type,
		OriginalURL:    a.SourceURL,
		IsTop:          a.IsTop,
		IsDelete: func() int {
			if a.DeletedAt.Valid {
				return 1
			}
			return 0
		}(),
		Status:     a.Status,
		ViewsCount: a.ViewCount,
		LikeCount:  a.LikeCount,
		CreatedAt:  formatTime(a.CreatedAt),
		UpdatedAt:  formatTime(a.UpdatedAt),
	}

	if a.Category.ID > 0 {
		vo.CategoryName = a.Category.Name
	}

	tagNames := make([]string, 0, len(a.Tags))
	for _, t := range a.Tags {
		tagNames = append(tagNames, t.Name)
	}
	vo.TagNameList = tagNames

	return vo
}

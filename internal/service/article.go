package service

import (
	"elian-blog/internal/dao"
	"elian-blog/internal/model"
)

type ArticleService struct {
	articleDao  *dao.ArticleDao
	categoryDao *dao.CategoryDao
	tagDao      *dao.TagDao
}

func NewArticleService(articleDao *dao.ArticleDao, categoryDao *dao.CategoryDao, tagDao *dao.TagDao) *ArticleService {
	return &ArticleService{
		articleDao:  articleDao,
		categoryDao: categoryDao,
		tagDao:      tagDao,
	}
}

type ArticleCreateReq struct {
	Title      string   `json:"title" binding:"required"`
	Summary    string   `json:"summary"`
	Content    string   `json:"content" binding:"required"`
	Cover      string   `json:"cover"`
	CategoryID uint     `json:"category_id"`
	AuthorID   uint     `json:"author_id"`
	Status     int      `json:"status"`
	IsTop      int      `json:"is_top"`
	IsOriginal int      `json:"is_original"`
	SourceURL  string   `json:"source_url"`
	Type       int      `json:"type"`
	Password   string   `json:"password"`
	TagNames   []string `json:"tag_names"`
}

type ArticleUpdateReq struct {
	ID         uint     `json:"id" binding:"required"`
	Title      string   `json:"title"`
	Summary    string   `json:"summary"`
	Content    string   `json:"content"`
	Cover      string   `json:"cover"`
	CategoryID uint     `json:"category_id"`
	Status     int      `json:"status"`
	IsTop      int      `json:"is_top"`
	IsOriginal int      `json:"is_original"`
	SourceURL  string   `json:"source_url"`
	Type       int      `json:"type"`
	Password   string   `json:"password"`
	TagNames   []string `json:"tag_names"`
}

type ArticleQueryReq struct {
	Page       int  `json:"page" form:"page"`
	PageSize   int  `json:"page_size" form:"page_size"`
	Status     int  `json:"status" form:"status"`
	CategoryID uint `json:"category_id" form:"category_id"`
	TagID      uint `json:"tag_id" form:"tag_id"`
}

func (s *ArticleService) Create(req *ArticleCreateReq) (*model.Article, error) {
	var catID *uint
	if req.CategoryID > 0 {
		catID = &req.CategoryID
	}
	article := &model.Article{
		Title:      req.Title,
		Summary:    req.Summary,
		Content:    req.Content,
		Cover:      req.Cover,
		CategoryID: catID,
		AuthorID:   req.AuthorID,
		Status:     req.Status,
		IsTop:      req.IsTop,
		IsOriginal: req.IsOriginal,
		SourceURL:  req.SourceURL,
		Type:       req.Type,
		Password:   req.Password,
	}

	if err := s.articleDao.Create(article); err != nil {
		return nil, err
	}

	if len(req.TagNames) > 0 {
		tags := make([]model.Tag, 0, len(req.TagNames))
		for _, name := range req.TagNames {
			tag, err := s.tagDao.FindOrCreate(name)
			if err != nil {
				continue
			}
			tags = append(tags, *tag)
		}
		_ = s.articleDao.UpdateTags(article, tags)
	}

	return s.articleDao.GetByID(article.ID)
}

func (s *ArticleService) Update(req *ArticleUpdateReq) (*model.Article, error) {
	article, err := s.articleDao.GetByID(req.ID)
	if err != nil {
		return nil, err
	}

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
	if req.CategoryID > 0 {
		article.CategoryID = &req.CategoryID
	} else {
		article.CategoryID = nil
	}
	article.Status = req.Status
	article.IsTop = req.IsTop
	article.IsOriginal = req.IsOriginal
	article.SourceURL = req.SourceURL
	article.Type = req.Type
	article.Password = req.Password

	if err := s.articleDao.Update(article); err != nil {
		return nil, err
	}

	if req.TagNames != nil {
		tags := make([]model.Tag, 0, len(req.TagNames))
		for _, name := range req.TagNames {
			tag, err := s.tagDao.FindOrCreate(name)
			if err != nil {
				continue
			}
			tags = append(tags, *tag)
		}
		_ = s.articleDao.UpdateTags(article, tags)
	}

	return s.articleDao.GetByID(article.ID)
}

func (s *ArticleService) Delete(id uint) error {
	return s.articleDao.Delete(id)
}

func (s *ArticleService) GetByID(id uint) (*model.Article, error) {
	article, err := s.articleDao.GetByID(id)
	if err != nil {
		return nil, err
	}
	_ = s.articleDao.IncrementViewCount(id)
	return article, nil
}

func (s *ArticleService) List(req *ArticleQueryReq) ([]model.Article, int64, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	return s.articleDao.List(req.Page, req.PageSize, req.Status, req.CategoryID, req.TagID, 0)
}

func (s *ArticleService) ListPublished(page, pageSize int) ([]model.Article, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return s.articleDao.List(page, pageSize, 1, 0, 0, 0)
}

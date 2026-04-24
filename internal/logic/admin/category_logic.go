package admin

import (
	"context"
	"fmt"
	"time"

	"elian-blog/internal/model"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type CategoryLogic struct {
	svcCtx *svc.ServiceContext
}

func NewCategoryLogic(svcCtx *svc.ServiceContext) *CategoryLogic {
	return &CategoryLogic{svcCtx: svcCtx}
}

func (l *CategoryLogic) Create(ctx context.Context, req *types.CreateCategoryReq) (interface{}, error) {
	cat := &model.Category{
		Name:        req.CategoryName,
		Description: req.Description,
		Sort:        req.Sort,
	}
	if err := l.svcCtx.CategoryDao.Create(cat); err != nil {
		return nil, err
	}
	return toCategoryVO(cat), nil
}

func (l *CategoryLogic) Update(ctx context.Context, req *types.UpdateCategoryReq) error {
	cat, err := l.svcCtx.CategoryDao.GetByID(req.ID)
	if err != nil {
		return err
	}

	if req.CategoryName != "" {
		cat.Name = req.CategoryName
	}
	if req.Description != "" {
		cat.Description = req.Description
	}
	if req.Sort != 0 {
		cat.Sort = req.Sort
	}

	return l.svcCtx.CategoryDao.Update(cat)
}

func (l *CategoryLogic) Delete(ctx context.Context, id uint) error {
	return l.svcCtx.CategoryDao.Delete(id)
}

func (l *CategoryLogic) List(ctx context.Context, req *types.QueryCategoryReq) (interface{}, int64, error) {
	cats, err := l.svcCtx.CategoryDao.List()
	if err != nil {
		return nil, 0, err
	}

	list := make([]types.CategoryVO, 0, len(cats))
	for _, cat := range cats {
		count, _ := l.svcCtx.CategoryDao.CountArticles(cat.ID)
		vo := toCategoryVO(&cat)
		vo.ArticleCount = int(count)
		list = append(list, vo)
	}

	total := int64(len(list))
	page := req.Page
	pageSize := req.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	return list, total, nil
}

func toCategoryVO(cat *model.Category) types.CategoryVO {
	return types.CategoryVO{
		ID:           cat.ID,
		CategoryName: cat.Name,
		Description:  cat.Description,
		Sort:         cat.Sort,
		CreatedAt:    formatTime(cat.CreatedAt),
		UpdatedAt:    formatTime(cat.UpdatedAt),
	}
}

func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return fmt.Sprintf("%d", t.Unix())
}

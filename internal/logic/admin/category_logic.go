package admin

import (
	"context"

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
		Name:        req.Name,
		Description: req.Description,
		Sort:        req.Sort,
	}
	if err := l.svcCtx.CategoryDao.Create(cat); err != nil {
		return nil, err
	}
	return cat, nil
}

func (l *CategoryLogic) Update(ctx context.Context, req *types.UpdateCategoryReq) error {
	cat, err := l.svcCtx.CategoryDao.GetByID(req.ID)
	if err != nil {
		return err
	}

	if req.Name != "" {
		cat.Name = req.Name
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

func (l *CategoryLogic) List(ctx context.Context) (interface{}, error) {
	cats, err := l.svcCtx.CategoryDao.List()
	if err != nil {
		return nil, err
	}

	// Build VO list with article count
	list := make([]types.CategoryVO, 0, len(cats))
	for _, cat := range cats {
		count, err := l.svcCtx.CategoryDao.CountArticles(cat.ID)
		if err != nil {
			count = 0
		}
		list = append(list, types.CategoryVO{
			ID:           cat.ID,
			Name:         cat.Name,
			Description:  cat.Description,
			Sort:         cat.Sort,
			ArticleCount: int(count),
		})
	}

	return list, nil
}

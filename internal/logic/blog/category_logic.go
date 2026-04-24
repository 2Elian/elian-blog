package blog

import (
	"context"

	"elian-blog/internal/svc"
)

type CategoryLogic struct {
	svcCtx *svc.ServiceContext
}

func NewCategoryLogic(svcCtx *svc.ServiceContext) *CategoryLogic {
	return &CategoryLogic{svcCtx: svcCtx}
}

// List 获取分类列表
func (l *CategoryLogic) List(ctx context.Context) (interface{}, error) {
	categories, err := l.svcCtx.CategoryDao.List()
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0, len(categories))
	for _, cat := range categories {
		count, _ := l.svcCtx.CategoryDao.CountArticles(cat.ID)
		result = append(result, map[string]interface{}{
			"id":            cat.ID,
			"name":          cat.Name,
			"description":   cat.Description,
			"sort":          cat.Sort,
			"article_count": count,
		})
	}

	return result, nil
}

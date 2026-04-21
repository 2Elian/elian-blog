package blog

import (
	"context"

	"elian-blog/internal/svc"
	"elian-blog/internal/types"
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

	// 转换为 CategoryVO 列表，并添加文章数量
	result := make([]types.CategoryVO, 0, len(categories))
	for _, cat := range categories {
		count, _ := l.svcCtx.CategoryDao.CountArticles(cat.ID)
		result = append(result, types.CategoryVO{
			ID:           cat.ID,
			Name:         cat.Name,
			Description:  cat.Description,
			Sort:         cat.Sort,
			ArticleCount: int(count),
		})
	}

	return result, nil
}
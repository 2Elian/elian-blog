package blog

import (
	"context"
	"errors"

	"elian-blog/internal/svc"
	"elian-blog/internal/types"

	"gorm.io/gorm"
)

type PageLogic struct {
	svcCtx *svc.ServiceContext
}

func NewPageLogic(svcCtx *svc.ServiceContext) *PageLogic {
	return &PageLogic{svcCtx: svcCtx}
}

// List 获取页面列表
func (l *PageLogic) List(ctx context.Context) (interface{}, error) {
	pages, err := l.svcCtx.PageDao.List()
	if err != nil {
		return nil, err
	}

	// 转换为 PageVO 列表
	result := make([]types.PageVO, 0, len(pages))
	for _, page := range pages {
		result = append(result, types.PageVO{
			ID:        page.ID,
			Title:     page.Title,
			Content:   page.Content,
			Slug:      page.Slug,
			Sort:      page.Sort,
			Status:    page.Status,
			CreatedAt: page.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: page.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return result, nil
}

// GetBySlug 根据slug获取页面
func (l *PageLogic) GetBySlug(ctx context.Context, slug string) (interface{}, error) {
	page, err := l.svcCtx.PageDao.GetBySlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("页面不存在")
		}
		return nil, err
	}

	return types.PageVO{
		ID:        page.ID,
		Title:     page.Title,
		Content:   page.Content,
		Slug:      page.Slug,
		Sort:      page.Sort,
		Status:    page.Status,
		CreatedAt: page.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: page.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
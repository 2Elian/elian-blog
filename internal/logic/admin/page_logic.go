package admin

import (
	"context"

	"elian-blog/internal/model"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type PageLogic struct {
	svcCtx *svc.ServiceContext
}

func NewPageLogic(svcCtx *svc.ServiceContext) *PageLogic {
	return &PageLogic{svcCtx: svcCtx}
}

func (l *PageLogic) List(ctx context.Context, req *types.QueryPageReq) (interface{}, int64, error) {
	pages, err := l.svcCtx.PageDao.ListAdmin()
	if err != nil {
		return nil, 0, err
	}

	list := make([]types.PageBackVO, 0, len(pages))
	for _, p := range pages {
		list = append(list, toPageBackVO(&p))
	}

	total := int64(len(list))
	return list, total, nil
}

func (l *PageLogic) Create(ctx context.Context, req *types.CreatePageReq) (interface{}, error) {
	// Accept both field naming conventions
	title := req.Title
	if title == "" {
		title = req.PageName
	}
	slug := req.Slug
	if slug == "" {
		slug = req.PageLabel
	}
	cover := req.Cover
	if cover == "" {
		cover = req.PageCover
	}

	page := &model.Page{
		Title:   title,
		Content: req.Content,
		Slug:    slug,
		Cover:   cover,
		Sort:    req.Sort,
		Status:  req.Status,
	}
	if page.Status == 0 {
		page.Status = 1
	}
	if err := l.svcCtx.PageDao.Create(page); err != nil {
		return nil, err
	}
	return toPageBackVO(page), nil
}

func (l *PageLogic) Update(ctx context.Context, req *types.UpdatePageReq) error {
	page, err := l.svcCtx.PageDao.GetByID(req.ID)
	if err != nil {
		return err
	}

	// Accept both field naming conventions
	title := req.Title
	if title == "" {
		title = req.PageName
	}
	slug := req.Slug
	if slug == "" {
		slug = req.PageLabel
	}
	cover := req.Cover
	if cover == "" {
		cover = req.PageCover
	}

	if title != "" {
		page.Title = title
	}
	if req.Content != "" {
		page.Content = req.Content
	}
	if slug != "" {
		page.Slug = slug
	}
	if cover != "" {
		page.Cover = cover
	}
	if req.Sort != 0 {
		page.Sort = req.Sort
	}
	if req.Status != 0 {
		page.Status = req.Status
	}

	return l.svcCtx.PageDao.Update(page)
}

func (l *PageLogic) Delete(ctx context.Context, id uint) error {
	return l.svcCtx.PageDao.Delete(id)
}

func toPageBackVO(p *model.Page) types.PageBackVO {
	return types.PageBackVO{
		ID:        p.ID,
		Title:     p.Title,
		PageName:  p.Title,
		Content:   p.Content,
		Slug:      p.Slug,
		PageLabel: p.Slug,
		Cover:     p.Cover,
		PageCover: p.Cover,
		Sort:      p.Sort,
		Status:    p.Status,
		CreatedAt: formatTime(p.CreatedAt),
		UpdatedAt: formatTime(p.UpdatedAt),
	}
}

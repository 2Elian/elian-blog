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

func (l *PageLogic) List(ctx context.Context) (interface{}, error) {
	return l.svcCtx.PageDao.ListAdmin()
}

func (l *PageLogic) Create(ctx context.Context, req *types.CreatePageReq) (interface{}, error) {
	page := &model.Page{
		Title:   req.Title,
		Content: req.Content,
		Slug:    req.Slug,
		Sort:    req.Sort,
		Status:  req.Status,
	}
	if err := l.svcCtx.PageDao.Create(page); err != nil {
		return nil, err
	}
	return page, nil
}

func (l *PageLogic) Update(ctx context.Context, req *types.UpdatePageReq) error {
	page, err := l.svcCtx.PageDao.GetByID(req.ID)
	if err != nil {
		return err
	}

	if req.Title != "" {
		page.Title = req.Title
	}
	if req.Content != "" {
		page.Content = req.Content
	}
	if req.Slug != "" {
		page.Slug = req.Slug
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

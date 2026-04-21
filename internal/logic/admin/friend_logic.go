package admin

import (
	"context"

	"elian-blog/internal/model"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type FriendLogic struct {
	svcCtx *svc.ServiceContext
}

func NewFriendLogic(svcCtx *svc.ServiceContext) *FriendLogic {
	return &FriendLogic{svcCtx: svcCtx}
}

func (l *FriendLogic) List(ctx context.Context) (interface{}, error) {
	return l.svcCtx.FriendDao.ListAdmin()
}

func (l *FriendLogic) Create(ctx context.Context, req *types.CreateFriendLinkReq) (interface{}, error) {
	link := &model.FriendLink{
		Name:        req.Name,
		URL:         req.URL,
		Logo:        req.Logo,
		Description: req.Description,
		Sort:        req.Sort,
	}
	if err := l.svcCtx.FriendDao.Create(link); err != nil {
		return nil, err
	}
	return link, nil
}

func (l *FriendLogic) Update(ctx context.Context, req *types.UpdateFriendLinkReq) error {
	link, err := l.svcCtx.FriendDao.GetByID(req.ID)
	if err != nil {
		return err
	}

	if req.Name != "" {
		link.Name = req.Name
	}
	if req.URL != "" {
		link.URL = req.URL
	}
	if req.Logo != "" {
		link.Logo = req.Logo
	}
	if req.Description != "" {
		link.Description = req.Description
	}
	if req.Sort != 0 {
		link.Sort = req.Sort
	}
	if req.Status != 0 {
		link.Status = req.Status
	}

	return l.svcCtx.FriendDao.Update(link)
}

func (l *FriendLogic) Delete(ctx context.Context, id uint) error {
	return l.svcCtx.FriendDao.Delete(id)
}

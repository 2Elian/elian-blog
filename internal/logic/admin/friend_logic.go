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

func (l *FriendLogic) List(ctx context.Context, req *types.QueryFriendReq) (interface{}, int64, error) {
	links, err := l.svcCtx.FriendDao.ListAdmin()
	if err != nil {
		return nil, 0, err
	}

	list := make([]types.FriendLinkVO, 0, len(links))
	for _, link := range links {
		list = append(list, toFriendLinkVO(&link))
	}

	total := int64(len(list))
	return list, total, nil
}

func (l *FriendLogic) Create(ctx context.Context, req *types.CreateFriendLinkReq) (interface{}, error) {
	link := &model.FriendLink{
		Name:        req.LinkName,
		URL:         req.LinkAddress,
		Logo:        req.LinkAvatar,
		Description: req.LinkIntro,
		Sort:        req.Sort,
	}
	if err := l.svcCtx.FriendDao.Create(link); err != nil {
		return nil, err
	}
	return toFriendLinkVO(link), nil
}

func (l *FriendLogic) Update(ctx context.Context, req *types.UpdateFriendLinkReq) error {
	link, err := l.svcCtx.FriendDao.GetByID(req.ID)
	if err != nil {
		return err
	}

	if req.LinkName != "" {
		link.Name = req.LinkName
	}
	if req.LinkAddress != "" {
		link.URL = req.LinkAddress
	}
	if req.LinkAvatar != "" {
		link.Logo = req.LinkAvatar
	}
	if req.LinkIntro != "" {
		link.Description = req.LinkIntro
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

func toFriendLinkVO(link *model.FriendLink) types.FriendLinkVO {
	return types.FriendLinkVO{
		ID:          link.ID,
		LinkName:    link.Name,
		LinkAvatar:  link.Logo,
		LinkAddress: link.URL,
		LinkIntro:   link.Description,
		Sort:        link.Sort,
		Status:      link.Status,
		CreatedAt:   formatTime(link.CreatedAt),
		UpdatedAt:   formatTime(link.UpdatedAt),
	}
}

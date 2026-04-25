package blog

import (
	"context"

	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type FriendLinkLogic struct {
	svcCtx *svc.ServiceContext
}

func NewFriendLinkLogic(svcCtx *svc.ServiceContext) *FriendLinkLogic {
	return &FriendLinkLogic{svcCtx: svcCtx}
}

// List 获取友链列表
func (l *FriendLinkLogic) List(ctx context.Context) (interface{}, error) {
	links, err := l.svcCtx.FriendDao.List()
	if err != nil {
		return nil, err
	}

	result := make([]types.FriendLinkVO, 0, len(links))
	for _, link := range links {
		result = append(result, types.FriendLinkVO{
			ID:          link.ID,
			LinkName:    link.Name,
			LinkAddress: link.URL,
			LinkAvatar:  link.Logo,
			LinkIntro:   link.Description,
			Sort:        link.Sort,
			Status:      link.Status,
		})
	}

	return result, nil
}

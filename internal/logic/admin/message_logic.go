package admin

import (
	"context"

	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type MessageLogic struct {
	svcCtx *svc.ServiceContext
}

func NewMessageLogic(svcCtx *svc.ServiceContext) *MessageLogic {
	return &MessageLogic{svcCtx: svcCtx}
}

func (l *MessageLogic) List(ctx context.Context, req *types.PageQuery) (interface{}, error) {
	page := req.Page
	pageSize := req.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	messages, total, err := l.svcCtx.MessageDao.ListAdmin(page, pageSize)
	if err != nil {
		return nil, err
	}

	return types.PageResp{
		List:     messages,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

func (l *MessageLogic) Delete(ctx context.Context, id uint) error {
	return l.svcCtx.MessageDao.Delete(id)
}

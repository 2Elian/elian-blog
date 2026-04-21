package admin

import (
	"context"

	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type UserLogic struct {
	svcCtx *svc.ServiceContext
}

func NewUserLogic(svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{svcCtx: svcCtx}
}

func (l *UserLogic) List(ctx context.Context, req *types.PageQuery) (interface{}, error) {
	page := req.Page
	pageSize := req.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	users, total, err := l.svcCtx.UserDao.List(page, pageSize)
	if err != nil {
		return nil, err
	}

	return types.PageResp{
		List:     users,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

func (l *UserLogic) Delete(ctx context.Context, id uint) error {
	return l.svcCtx.UserDao.Delete(id)
}

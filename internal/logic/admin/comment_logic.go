package admin

import (
	"context"

	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type CommentLogic struct {
	svcCtx *svc.ServiceContext
}

func NewCommentLogic(svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{svcCtx: svcCtx}
}

func (l *CommentLogic) List(ctx context.Context, req *types.QueryCommentReq) (interface{}, error) {
	page := req.Page
	pageSize := req.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	comments, total, err := l.svcCtx.CommentDao.ListAdmin(page, pageSize)
	if err != nil {
		return nil, err
	}

	return types.PageResp{
		List:     comments,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

func (l *CommentLogic) Delete(ctx context.Context, id uint) error {
	return l.svcCtx.CommentDao.Delete(id)
}

func (l *CommentLogic) UpdateStatus(ctx context.Context, req *types.UpdateCommentStatusReq) error {
	for _, id := range req.IDs {
		if err := l.svcCtx.CommentDao.UpdateStatus(id, req.Status); err != nil {
			return err
		}
	}
	return nil
}
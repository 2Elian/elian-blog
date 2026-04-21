package blog

import (
	"context"
	"errors"
	"time"

	"elian-blog/internal/model"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type CommentLogic struct {
	svcCtx *svc.ServiceContext
}

func NewCommentLogic(svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{svcCtx: svcCtx}
}

// ListByArticle 获取文章评论列表
func (l *CommentLogic) ListByArticle(ctx context.Context, req *types.QueryCommentReq) (interface{}, error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	comments, total, err := l.svcCtx.CommentDao.ListByArticle(req.ArticleID, page, pageSize)
	if err != nil {
		return nil, err
	}

	// 转换为 CommentVO 列表
	result := l.convertToCommentVOList(comments)

	return map[string]interface{}{
		"list":  result,
		"total": total,
		"page":  page,
		"size":  pageSize,
	}, nil
}

// Recent 获取最近评论
func (l *CommentLogic) Recent(ctx context.Context) (interface{}, error) {
	comments, err := l.svcCtx.CommentDao.ListRecent(10)
	if err != nil {
		return nil, err
	}

	// 转换为 CommentVO 列表
	result := l.convertToCommentVOList(comments)
	return result, nil
}

// Create 创建评论
func (l *CommentLogic) Create(ctx context.Context, req *types.CreateCommentReq) (interface{}, error) {
	// 从上下文中获取用户ID
	userID, exists := ctx.Value("user_id").(uint)
	if !exists || userID == 0 {
		return nil, errors.New("请先登录")
	}

	comment := &model.Comment{
		ArticleID: req.ArticleID,
		UserID:    userID,
		Content:   req.Content,
		Type:      req.Type,
		ParentID:  req.ParentID,
		Status:    1, // 评论默认通过审核
	}

	if err := l.svcCtx.CommentDao.Create(comment); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id": comment.ID,
	}, nil
}

// convertToCommentVOList 转换评论列表
func (l *CommentLogic) convertToCommentVOList(comments []model.Comment) []types.CommentVO {
	result := make([]types.CommentVO, 0, len(comments))
	for _, comment := range comments {
		username := ""
		avatar := ""
		if comment.User.ID != 0 {
			username = comment.User.Username
			avatar = comment.User.Avatar
		}

		vo := types.CommentVO{
			ID:        comment.ID,
			ArticleID: comment.ArticleID,
			UserID:    comment.UserID,
			Content:   comment.Content,
			Type:      comment.Type,
			ParentID:  comment.ParentID,
			Username:  username,
			Avatar:    avatar,
			CreatedAt: comment.CreatedAt.Format(time.DateTime),
		}
		result = append(result, vo)
	}
	return result
}

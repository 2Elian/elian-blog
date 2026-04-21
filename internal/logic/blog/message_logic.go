package blog

import (
	"context"
	"errors"
	"time"

	"elian-blog/internal/model"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type MessageLogic struct {
	svcCtx *svc.ServiceContext
}

func NewMessageLogic(svcCtx *svc.ServiceContext) *MessageLogic {
	return &MessageLogic{svcCtx: svcCtx}
}

// List 获取留言列表
func (l *MessageLogic) List(ctx context.Context, req *types.PageQuery) (interface{}, error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	messages, total, err := l.svcCtx.MessageDao.List(page, pageSize)
	if err != nil {
		return nil, err
	}

	// 转换为 MessageVO 列表
	result := l.convertToMessageVOList(messages)

	return map[string]interface{}{
		"list":  result,
		"total": total,
		"page":  page,
		"size":  pageSize,
	}, nil
}

// Create 创建留言
func (l *MessageLogic) Create(ctx context.Context, req *types.CreateMessageReq) (interface{}, error) {
	// 从上下文中获取用户ID
	userID, exists := ctx.Value("user_id").(uint)
	if !exists || userID == 0 {
		return nil, errors.New("请先登录")
	}

	message := &model.Message{
		UserID:  userID,
		Content: req.Content,
		Images:  req.Images,
		Status:  1,
	}

	if err := l.svcCtx.MessageDao.Create(message); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id": message.ID,
	}, nil
}

// convertToMessageVOList 转换留言列表
func (l *MessageLogic) convertToMessageVOList(messages []model.Message) []types.MessageVO {
	result := make([]types.MessageVO, 0, len(messages))
	for _, msg := range messages {
		username := ""
		avatar := ""
		if msg.User.ID != 0 {
			username = msg.User.Username
			avatar = msg.User.Avatar
		}

		vo := types.MessageVO{
			ID:        msg.ID,
			UserID:    msg.UserID,
			Content:   msg.Content,
			Images:    msg.Images,
			Username:  username,
			Avatar:    avatar,
			CreatedAt: msg.CreatedAt.Format(time.DateTime),
		}
		result = append(result, vo)
	}
	return result
}

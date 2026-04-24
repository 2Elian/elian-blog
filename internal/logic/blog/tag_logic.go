package blog

import (
	"context"

	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type TagLogic struct {
	svcCtx *svc.ServiceContext
}

func NewTagLogic(svcCtx *svc.ServiceContext) *TagLogic {
	return &TagLogic{svcCtx: svcCtx}
}

// List 获取标签列表
func (l *TagLogic) List(ctx context.Context) (interface{}, error) {
	tags, err := l.svcCtx.TagDao.List()
	if err != nil {
		return nil, err
	}

	result := make([]types.TagVO, 0, len(tags))
	for _, tag := range tags {
		result = append(result, types.TagVO{
			ID:      tag.ID,
			TagName: tag.Name,
			Color:   tag.Color,
		})
	}

	return result, nil
}

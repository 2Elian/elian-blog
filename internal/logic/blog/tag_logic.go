package blog

import (
	"context"

	"elian-blog/internal/svc"
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

	result := make([]map[string]interface{}, 0, len(tags))
	for _, tag := range tags {
		var articleCount int64
		l.svcCtx.DB.Table("article_tags").Where("tag_id = ?", tag.ID).Count(&articleCount)
		result = append(result, map[string]interface{}{
			"id":            tag.ID,
			"name":          tag.Name,
			"color":         tag.Color,
			"article_count": articleCount,
		})
	}

	return result, nil
}

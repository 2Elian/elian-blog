package admin

import (
	"context"

	"elian-blog/internal/model"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type TagLogic struct {
	svcCtx *svc.ServiceContext
}

func NewTagLogic(svcCtx *svc.ServiceContext) *TagLogic {
	return &TagLogic{svcCtx: svcCtx}
}

func (l *TagLogic) Create(ctx context.Context, req *types.CreateTagReq) (interface{}, error) {
	tag := &model.Tag{
		Name:  req.TagName,
		Color: req.Color,
	}
	if err := l.svcCtx.TagDao.Create(tag); err != nil {
		return nil, err
	}
	return toTagVO(tag), nil
}

func (l *TagLogic) Update(ctx context.Context, req *types.UpdateTagReq) error {
	tag, err := l.svcCtx.TagDao.GetByID(req.ID)
	if err != nil {
		return err
	}

	if req.TagName != "" {
		tag.Name = req.TagName
	}
	if req.Color != "" {
		tag.Color = req.Color
	}

	return l.svcCtx.TagDao.Update(tag)
}

func (l *TagLogic) Delete(ctx context.Context, id uint) error {
	return l.svcCtx.TagDao.Delete(id)
}

func (l *TagLogic) List(ctx context.Context, req *types.QueryTagReq) (interface{}, int64, error) {
	tags, err := l.svcCtx.TagDao.List()
	if err != nil {
		return nil, 0, err
	}

	list := make([]types.TagVO, 0, len(tags))
	for _, tag := range tags {
		list = append(list, toTagVO(&tag))
	}

	total := int64(len(list))
	return list, total, nil
}

func toTagVO(tag *model.Tag) types.TagVO {
	return types.TagVO{
		ID:        tag.ID,
		TagName:   tag.Name,
		Color:     tag.Color,
		CreatedAt: formatTime(tag.CreatedAt),
		UpdatedAt: formatTime(tag.UpdatedAt),
	}
}

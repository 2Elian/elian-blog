package admin

import (
	"context"

	"elian-blog/internal/model"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type SiteLogic struct {
	svcCtx *svc.ServiceContext
}

func NewSiteLogic(svcCtx *svc.ServiceContext) *SiteLogic {
	return &SiteLogic{svcCtx: svcCtx}
}

func (l *SiteLogic) GetConfig(ctx context.Context) (interface{}, error) {
	return l.svcCtx.SiteDao.List()
}

func (l *SiteLogic) SetConfig(ctx context.Context, req *types.SetSiteConfigReq) error {
	cfg := &model.SiteConfig{
		Key:   req.Key,
		Value: req.Value,
	}
	return l.svcCtx.SiteDao.Upsert(cfg)
}

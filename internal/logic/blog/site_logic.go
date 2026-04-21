package blog

import (
	"context"

	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type SiteLogic struct {
	svcCtx *svc.ServiceContext
}

func NewSiteLogic(svcCtx *svc.ServiceContext) *SiteLogic {
	return &SiteLogic{svcCtx: svcCtx}
}

// GetConfig 获取网站配置
func (l *SiteLogic) GetConfig(ctx context.Context) (interface{}, error) {
	configs, err := l.svcCtx.SiteDao.List()
	if err != nil {
		return nil, err
	}

	// 转换为 SiteConfigVO 列表
	result := make([]types.SiteConfigVO, 0, len(configs))
	for _, cfg := range configs {
		result = append(result, types.SiteConfigVO{
			ID:    cfg.ID,
			Key:   cfg.Key,
			Value: cfg.Value,
		})
	}

	return result, nil
}
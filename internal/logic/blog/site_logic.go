package blog

import (
	"context"
	"encoding/json"

	"elian-blog/internal/svc"
)

type SiteLogic struct {
	svcCtx *svc.ServiceContext
}

func NewSiteLogic(svcCtx *svc.ServiceContext) *SiteLogic {
	return &SiteLogic{svcCtx: svcCtx}
}

// GetConfig 获取网站配置（返回与 admin 相同的结构化 JSON）
func (l *SiteLogic) GetConfig(ctx context.Context) (interface{}, error) {
	defaultConfig := map[string]interface{}{
		"admin_url":      "",
		"websocket_url":  "",
		"tourist_avatar": "",
		"user_avatar":    "",
		"website_info": map[string]interface{}{
			"website_author":      "",
			"website_avatar":      "",
			"website_create_time": "",
			"website_intro":       "",
			"website_name":        "",
			"website_notice":      "",
			"website_record_no":   "",
		},
		"website_feature": map[string]interface{}{
			"is_chat_room":     1,
			"is_ai_assistant":  1,
			"is_music_player":  1,
			"is_comment_review": 0,
			"is_email_notice":  1,
			"is_message_review": 0,
			"is_reward":        0,
		},
		"reward_qr_code": map[string]interface{}{
			"alipay_qr_code": "",
			"weixin_qr_code": "",
		},
		"social_login_list": []interface{}{},
		"social_url_list":   []interface{}{},
	}

	cfg, err := l.svcCtx.SiteDao.GetByKey("website_config")
	if err != nil || cfg == nil {
		return defaultConfig, nil
	}
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(cfg.Value), &result); err != nil {
		return defaultConfig, nil
	}
	// Merge with defaults to ensure all fields exist
	for k, v := range defaultConfig {
		if _, exists := result[k]; !exists {
			result[k] = v
		}
		if m, ok := v.(map[string]interface{}); ok {
			if rm, ok := result[k].(map[string]interface{}); ok {
				for dk, dv := range m {
					if _, exists := rm[dk]; !exists {
						rm[dk] = dv
					}
				}
			}
		}
	}
	return result, nil
}

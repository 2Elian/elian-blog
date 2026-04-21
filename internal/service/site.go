package service

import (
	"elian-blog/internal/dao"
	"elian-blog/internal/model"
	"encoding/json"
)

type SiteConfigService struct {
	dao *dao.SiteConfigDao
}

func NewSiteConfigService(dao *dao.SiteConfigDao) *SiteConfigService {
	return &SiteConfigService{dao: dao}
}

func (s *SiteConfigService) GetByKey(key string) (string, error) {
	cfg, err := s.dao.GetByKey(key)
	if err != nil {
		return "", err
	}
	return cfg.Value, nil
}

func (s *SiteConfigService) GetJSON(key string, out interface{}) error {
	val, err := s.GetByKey(key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), out)
}

func (s *SiteConfigService) Set(key, value string) error {
	cfg := &model.SiteConfig{Key: key, Value: value}
	return s.dao.Upsert(cfg)
}

func (s *SiteConfigService) List() (map[string]string, error) {
	configs, err := s.dao.List()
	if err != nil {
		return nil, err
	}
	result := make(map[string]string, len(configs))
	for _, c := range configs {
		result[c.Key] = c.Value
	}
	return result, nil
}

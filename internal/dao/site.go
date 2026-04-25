package dao

import (
	"elian-blog/internal/model"

	"gorm.io/gorm"
)

type SiteConfigDao struct {
	db *gorm.DB
}

func NewSiteConfigDao(db *gorm.DB) *SiteConfigDao {
	return &SiteConfigDao{db: db}
}

func (d *SiteConfigDao) GetByKey(key string) (*model.SiteConfig, error) {
	var cfg model.SiteConfig
	err := d.db.Where("`key` = ?", key).First(&cfg).Error
	return &cfg, err
}

func (d *SiteConfigDao) Upsert(cfg *model.SiteConfig) error {
	return d.db.Save(cfg).Error
}

func (d *SiteConfigDao) Set(key, value string) error {
	var cfg model.SiteConfig
	err := d.db.Where("`key` = ?", key).First(&cfg).Error
	if err != nil {
		cfg = model.SiteConfig{Key: key, Value: value}
	} else {
		cfg.Value = value
	}
	return d.db.Save(&cfg).Error
}

func (d *SiteConfigDao) List() ([]model.SiteConfig, error) {
	var configs []model.SiteConfig
	err := d.db.Find(&configs).Error
	return configs, err
}

type OperationLogDao struct {
	db *gorm.DB
}

func NewOperationLogDao(db *gorm.DB) *OperationLogDao {
	return &OperationLogDao{db: db}
}

func (d *OperationLogDao) Create(log *model.OperationLog) error {
	return d.db.Create(log).Error
}

func (d *OperationLogDao) List(page, pageSize int) ([]model.OperationLog, int64, error) {
	var logs []model.OperationLog
	var total int64
	d.db.Model(&model.OperationLog{}).Count(&total)
	err := d.db.Preload("User").Offset((page - 1) * pageSize).Limit(pageSize).
		Order("created_at DESC").Find(&logs).Error
	return logs, total, err
}

package dao

import (
	"elian-blog/internal/model"

	"gorm.io/gorm"
)

type MessageDao struct {
	db *gorm.DB
}

func NewMessageDao(db *gorm.DB) *MessageDao {
	return &MessageDao{db: db}
}

func (d *MessageDao) Create(msg *model.Message) error {
	return d.db.Create(msg).Error
}

func (d *MessageDao) Delete(id uint) error {
	return d.db.Delete(&model.Message{}, id).Error
}

func (d *MessageDao) List(page, pageSize int) ([]model.Message, int64, error) {
	var messages []model.Message
	var total int64

	d.db.Model(&model.Message{}).Where("status = 1").Count(&total)
	err := d.db.Preload("User").Where("status = 1").
		Order("created_at DESC").
		Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&messages).Error
	return messages, total, err
}

func (d *MessageDao) ListAdmin(page, pageSize int) ([]model.Message, int64, error) {
	var messages []model.Message
	var total int64
	d.db.Model(&model.Message{}).Count(&total)
	err := d.db.Preload("User").Offset((page - 1) * pageSize).Limit(pageSize).
		Order("created_at DESC").Find(&messages).Error
	return messages, total, err
}

package dao

import (
	"elian-blog/internal/model"

	"gorm.io/gorm"
)

type FriendLinkDao struct {
	db *gorm.DB
}

func NewFriendLinkDao(db *gorm.DB) *FriendLinkDao {
	return &FriendLinkDao{db: db}
}

func (d *FriendLinkDao) Create(link *model.FriendLink) error {
	return d.db.Create(link).Error
}

func (d *FriendLinkDao) Update(link *model.FriendLink) error {
	return d.db.Save(link).Error
}

func (d *FriendLinkDao) Delete(id uint) error {
	return d.db.Delete(&model.FriendLink{}, id).Error
}

func (d *FriendLinkDao) GetByID(id uint) (*model.FriendLink, error) {
	var link model.FriendLink
	err := d.db.First(&link, id).Error
	return &link, err
}

func (d *FriendLinkDao) List() ([]model.FriendLink, error) {
	var links []model.FriendLink
	err := d.db.Where("status = 1").Order("sort ASC, id ASC").Find(&links).Error
	return links, err
}

func (d *FriendLinkDao) ListAdmin() ([]model.FriendLink, error) {
	var links []model.FriendLink
	err := d.db.Order("sort ASC, id ASC").Find(&links).Error
	return links, err
}

package dao

import (
	"elian-blog/internal/model"

	"gorm.io/gorm"
)

type PageDao struct {
	db *gorm.DB
}

func NewPageDao(db *gorm.DB) *PageDao {
	return &PageDao{db: db}
}

func (d *PageDao) Create(page *model.Page) error {
	return d.db.Create(page).Error
}

func (d *PageDao) Update(page *model.Page) error {
	return d.db.Save(page).Error
}

func (d *PageDao) Delete(id uint) error {
	return d.db.Delete(&model.Page{}, id).Error
}

func (d *PageDao) GetByID(id uint) (*model.Page, error) {
	var page model.Page
	err := d.db.First(&page, id).Error
	return &page, err
}

func (d *PageDao) GetBySlug(slug string) (*model.Page, error) {
	var page model.Page
	err := d.db.Where("slug = ? AND status = 1", slug).First(&page).Error
	return &page, err
}

func (d *PageDao) List() ([]model.Page, error) {
	var pages []model.Page
	err := d.db.Where("status = 1").Order("sort ASC, id ASC").Find(&pages).Error
	return pages, err
}

func (d *PageDao) ListAdmin() ([]model.Page, error) {
	var pages []model.Page
	err := d.db.Order("sort ASC, id ASC").Find(&pages).Error
	return pages, err
}

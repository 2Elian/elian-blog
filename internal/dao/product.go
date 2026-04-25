package dao

import (
	"elian-blog/internal/model"

	"gorm.io/gorm"
)

type ProductDao struct {
	db *gorm.DB
}

func NewProductDao(db *gorm.DB) *ProductDao {
	return &ProductDao{db: db}
}

func (d *ProductDao) Create(p *model.Product) error {
	return d.db.Create(p).Error
}

func (d *ProductDao) Update(p *model.Product) error {
	return d.db.Save(p).Error
}

func (d *ProductDao) Delete(id uint) error {
	return d.db.Delete(&model.Product{}, id).Error
}

func (d *ProductDao) GetByID(id uint) (*model.Product, error) {
	var p model.Product
	err := d.db.First(&p, id).Error
	return &p, err
}

func (d *ProductDao) List(page, pageSize int) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64
	d.db.Model(&model.Product{}).Count(&total)
	err := d.db.Order("sort ASC, created_at DESC").
		Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&products).Error
	return products, total, err
}

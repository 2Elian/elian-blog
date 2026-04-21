package dao

import (
	"elian-blog/internal/model"

	"gorm.io/gorm"
)

type CategoryDao struct {
	db *gorm.DB
}

func NewCategoryDao(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db: db}
}

func (d *CategoryDao) Create(cat *model.Category) error {
	return d.db.Create(cat).Error
}

func (d *CategoryDao) Update(cat *model.Category) error {
	return d.db.Save(cat).Error
}

func (d *CategoryDao) Delete(id uint) error {
	return d.db.Delete(&model.Category{}, id).Error
}

func (d *CategoryDao) GetByID(id uint) (*model.Category, error) {
	var cat model.Category
	err := d.db.First(&cat, id).Error
	return &cat, err
}

func (d *CategoryDao) List() ([]model.Category, error) {
	var cats []model.Category
	err := d.db.Order("sort ASC, id ASC").Find(&cats).Error
	return cats, err
}

func (d *CategoryDao) CountArticles(categoryID uint) (int64, error) {
	var count int64
	err := d.db.Model(&model.Article{}).Where("category_id = ?", categoryID).Count(&count).Error
	return count, err
}

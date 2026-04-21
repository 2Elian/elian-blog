package dao

import (
	"elian-blog/internal/model"

	"gorm.io/gorm"
)

type TagDao struct {
	db *gorm.DB
}

func NewTagDao(db *gorm.DB) *TagDao {
	return &TagDao{db: db}
}

func (d *TagDao) Create(tag *model.Tag) error {
	return d.db.Create(tag).Error
}

func (d *TagDao) Update(tag *model.Tag) error {
	return d.db.Save(tag).Error
}

func (d *TagDao) Delete(id uint) error {
	return d.db.Delete(&model.Tag{}, id).Error
}

func (d *TagDao) GetByID(id uint) (*model.Tag, error) {
	var tag model.Tag
	err := d.db.First(&tag, id).Error
	return &tag, err
}

func (d *TagDao) GetByName(name string) (*model.Tag, error) {
	var tag model.Tag
	err := d.db.Where("name = ?", name).First(&tag).Error
	return &tag, err
}

func (d *TagDao) List() ([]model.Tag, error) {
	var tags []model.Tag
	err := d.db.Order("id ASC").Find(&tags).Error
	return tags, err
}

func (d *TagDao) FindOrCreate(name string) (*model.Tag, error) {
	tag, err := d.GetByName(name)
	if err == nil {
		return tag, nil
	}
	if err != gorm.ErrRecordNotFound {
		return nil, err
	}
	tag = &model.Tag{Name: name}
	if err := d.Create(tag); err != nil {
		return nil, err
	}
	return tag, nil
}

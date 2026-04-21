package dao

import (
	"elian-blog/internal/model"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (d *UserDao) Create(user *model.User) error {
	return d.db.Create(user).Error
}

func (d *UserDao) Update(user *model.User) error {
	return d.db.Save(user).Error
}

func (d *UserDao) Delete(id uint) error {
	return d.db.Delete(&model.User{}, id).Error
}

func (d *UserDao) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := d.db.Preload("Roles").First(&user, id).Error
	return &user, err
}

func (d *UserDao) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := d.db.Preload("Roles").Where("username = ?", username).First(&user).Error
	return &user, err
}

func (d *UserDao) List(page, pageSize int) ([]model.User, int64, error) {
	var users []model.User
	var total int64
	d.db.Model(&model.User{}).Count(&total)
	err := d.db.Preload("Roles").Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error
	return users, total, err
}

func (d *UserDao) UpdateLoginTime(id uint) error {
	return d.db.Model(&model.User{}).Where("id = ?", id).Update("last_login", gorm.Expr("NOW()")).Error
}

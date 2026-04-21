package dao

import (
	"elian-blog/internal/model"

	"gorm.io/gorm"
)

type RoleDao struct {
	db *gorm.DB
}

func NewRoleDao(db *gorm.DB) *RoleDao {
	return &RoleDao{db: db}
}

func (d *RoleDao) Create(role *model.Role) error {
	return d.db.Create(role).Error
}

func (d *RoleDao) Update(role *model.Role) error {
	return d.db.Save(role).Error
}

func (d *RoleDao) Delete(id uint) error {
	return d.db.Select("Menus").Delete(&model.Role{}, id).Error
}

func (d *RoleDao) GetByID(id uint) (*model.Role, error) {
	var role model.Role
	err := d.db.Preload("Menus").First(&role, id).Error
	return &role, err
}

func (d *RoleDao) List() ([]model.Role, error) {
	var roles []model.Role
	err := d.db.Preload("Menus").Order("sort ASC").Find(&roles).Error
	return roles, err
}

func (d *RoleDao) UpdateMenus(role *model.Role, menus []model.Menu) error {
	return d.db.Model(role).Association("Menus").Replace(menus)
}

type MenuDao struct {
	db *gorm.DB
}

func NewMenuDao(db *gorm.DB) *MenuDao {
	return &MenuDao{db: db}
}

func (d *MenuDao) Create(menu *model.Menu) error {
	return d.db.Create(menu).Error
}

func (d *MenuDao) Update(menu *model.Menu) error {
	return d.db.Save(menu).Error
}

func (d *MenuDao) Delete(id uint) error {
	return d.db.Delete(&model.Menu{}, id).Error
}

func (d *MenuDao) GetByID(id uint) (*model.Menu, error) {
	var menu model.Menu
	err := d.db.First(&menu, id).Error
	return &menu, err
}

func (d *MenuDao) List() ([]model.Menu, error) {
	var menus []model.Menu
	err := d.db.Order("sort ASC, id ASC").Find(&menus).Error
	return menus, err
}

package service

import (
	"elian-blog/internal/dao"
	"elian-blog/internal/model"
)

type RoleService struct {
	dao     *dao.RoleDao
	menuDao *dao.MenuDao
}

func NewRoleService(dao *dao.RoleDao, menuDao *dao.MenuDao) *RoleService {
	return &RoleService{dao: dao, menuDao: menuDao}
}

func (s *RoleService) Create(name, label, desc string, sort int) (*model.Role, error) {
	role := &model.Role{Name: name, Label: label, Description: desc, Sort: sort, Status: 1}
	return role, s.dao.Create(role)
}

func (s *RoleService) Update(role *model.Role) error {
	return s.dao.Update(role)
}

func (s *RoleService) Delete(id uint) error {
	return s.dao.Delete(id)
}

func (s *RoleService) GetByID(id uint) (*model.Role, error) {
	return s.dao.GetByID(id)
}

func (s *RoleService) List() ([]model.Role, error) {
	return s.dao.List()
}

func (s *RoleService) UpdateMenus(roleID uint, menuIDs []uint) error {
	role, err := s.dao.GetByID(roleID)
	if err != nil {
		return err
	}
	menus := make([]model.Menu, 0, len(menuIDs))
	for _, id := range menuIDs {
		menu, err := s.menuDao.GetByID(id)
		if err != nil {
			continue
		}
		menus = append(menus, *menu)
	}
	return s.dao.UpdateMenus(role, menus)
}

type MenuService struct {
	dao *dao.MenuDao
}

func NewMenuService(dao *dao.MenuDao) *MenuService {
	return &MenuService{dao: dao}
}

func (s *MenuService) Create(menu *model.Menu) error {
	return s.dao.Create(menu)
}

func (s *MenuService) Update(menu *model.Menu) error {
	return s.dao.Update(menu)
}

func (s *MenuService) Delete(id uint) error {
	return s.dao.Delete(id)
}

func (s *MenuService) GetByID(id uint) (*model.Menu, error) {
	return s.dao.GetByID(id)
}

func (s *MenuService) List() ([]model.Menu, error) {
	return s.dao.List()
}

func (s *MenuService) BuildTree(menus []model.Menu) []model.Menu {
	menuMap := make(map[uint]*model.Menu)
	for i := range menus {
		menuMap[menus[i].ID] = &menus[i]
	}

	var roots []model.Menu
	for i := range menus {
		if menus[i].ParentID == 0 {
			roots = append(roots, menus[i])
		} else {
			if parent, ok := menuMap[menus[i].ParentID]; ok {
				parent.Children = append(parent.Children, menus[i])
			}
		}
	}
	return roots
}
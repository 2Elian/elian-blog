package admin

import (
	"context"

	"elian-blog/internal/model"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type RoleLogic struct {
	svcCtx *svc.ServiceContext
}

func NewRoleLogic(svcCtx *svc.ServiceContext) *RoleLogic {
	return &RoleLogic{svcCtx: svcCtx}
}

func (l *RoleLogic) List(ctx context.Context) (interface{}, error) {
	return l.svcCtx.RoleDao.List()
}

func (l *RoleLogic) Create(ctx context.Context, req *types.CreateRoleReq) (interface{}, error) {
	role := &model.Role{
		Name:        req.Name,
		Label:       req.Label,
		Description: req.Description,
		Sort:        req.Sort,
	}
	if err := l.svcCtx.RoleDao.Create(role); err != nil {
		return nil, err
	}
	return role, nil
}

func (l *RoleLogic) Update(ctx context.Context, req *types.UpdateRoleReq) error {
	role, err := l.svcCtx.RoleDao.GetByID(req.ID)
	if err != nil {
		return err
	}

	if req.Name != "" {
		role.Name = req.Name
	}
	if req.Label != "" {
		role.Label = req.Label
	}
	if req.Description != "" {
		role.Description = req.Description
	}
	if req.Sort != 0 {
		role.Sort = req.Sort
	}

	return l.svcCtx.RoleDao.Update(role)
}

func (l *RoleLogic) Delete(ctx context.Context, id uint) error {
	return l.svcCtx.RoleDao.Delete(id)
}

func (l *RoleLogic) UpdateMenus(ctx context.Context, req *types.UpdateRoleMenusReq) error {
	role, err := l.svcCtx.RoleDao.GetByID(req.ID)
	if err != nil {
		return err
	}

	menus := make([]model.Menu, 0, len(req.MenuIDs))
	for _, menuID := range req.MenuIDs {
		menus = append(menus, model.Menu{Model: model.Model{ID: menuID}})
	}

	return l.svcCtx.RoleDao.UpdateMenus(role, menus)
}

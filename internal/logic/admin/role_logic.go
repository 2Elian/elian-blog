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

func (l *RoleLogic) List(ctx context.Context, req *types.QueryRoleReq) (interface{}, int64, error) {
	roles, err := l.svcCtx.RoleDao.List()
	if err != nil {
		return nil, 0, err
	}

	list := make([]types.RoleBackVO, 0, len(roles))
	for _, role := range roles {
		list = append(list, toRoleBackVO(&role))
	}

	total := int64(len(list))
	return list, total, nil
}

func (l *RoleLogic) Create(ctx context.Context, req *types.CreateRoleReq) (interface{}, error) {
	role := &model.Role{
		Name:        req.RoleLabel,
		Label:       req.RoleKey,
		Description: req.RoleComment,
	}
	if err := l.svcCtx.RoleDao.Create(role); err != nil {
		return nil, err
	}
	return toRoleBackVO(role), nil
}

func (l *RoleLogic) Update(ctx context.Context, req *types.UpdateRoleReq) error {
	role, err := l.svcCtx.RoleDao.GetByID(req.ID)
	if err != nil {
		return err
	}

	if req.RoleLabel != "" {
		role.Name = req.RoleLabel
	}
	if req.RoleKey != "" {
		role.Label = req.RoleKey
	}
	if req.RoleComment != "" {
		role.Description = req.RoleComment
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

func toRoleBackVO(role *model.Role) types.RoleBackVO {
	return types.RoleBackVO{
		ID:          role.ID,
		ParentID:    0,
		RoleKey:     role.Label,
		RoleLabel:   role.Name,
		RoleComment: role.Description,
		Status:      role.Status,
		CreatedAt:   formatTime(role.CreatedAt),
		UpdatedAt:   formatTime(role.UpdatedAt),
	}
}

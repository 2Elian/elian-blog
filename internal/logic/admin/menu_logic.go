package admin

import (
	"context"

	"elian-blog/internal/model"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type MenuLogic struct {
	svcCtx *svc.ServiceContext
}

func NewMenuLogic(svcCtx *svc.ServiceContext) *MenuLogic {
	return &MenuLogic{svcCtx: svcCtx}
}

func (l *MenuLogic) List(ctx context.Context) (interface{}, error) {
	menus, err := l.svcCtx.MenuDao.List()
	if err != nil {
		return nil, err
	}

	return buildMenuTree(menus), nil
}

func (l *MenuLogic) Create(ctx context.Context, req *types.CreateMenuReq) (interface{}, error) {
	menu := &model.Menu{
		Name:     req.Name,
		Path:     req.Path,
		Icon:     req.Icon,
		ParentID: req.ParentID,
		Sort:     req.Sort,
	}
	if err := l.svcCtx.MenuDao.Create(menu); err != nil {
		return nil, err
	}
	return menu, nil
}

func (l *MenuLogic) Update(ctx context.Context, req *types.UpdateMenuReq) error {
	menu, err := l.svcCtx.MenuDao.GetByID(req.ID)
	if err != nil {
		return err
	}

	if req.Name != "" {
		menu.Name = req.Name
	}
	if req.Path != "" {
		menu.Path = req.Path
	}
	if req.Icon != "" {
		menu.Icon = req.Icon
	}
	if req.Sort != 0 {
		menu.Sort = req.Sort
	}

	return l.svcCtx.MenuDao.Update(menu)
}

func (l *MenuLogic) Delete(ctx context.Context, id uint) error {
	return l.svcCtx.MenuDao.Delete(id)
}

// buildMenuTree converts a flat menu list into a tree structure.
func buildMenuTree(menus []model.Menu) []types.MenuVO {
	nodeMap := make(map[uint]*types.MenuVO)
	roots := make([]types.MenuVO, 0)

	// First pass: create all nodes
	for _, m := range menus {
		nodeMap[m.ID] = &types.MenuVO{
			ID:       m.ID,
			ParentID: m.ParentID,
			Name:     m.Name,
			Path:     m.Path,
			Icon:     m.Icon,
			Sort:     m.Sort,
			Children: make([]types.MenuVO, 0),
		}
	}

	// Second pass: build tree
	for _, m := range menus {
		node := nodeMap[m.ID]
		if m.ParentID == 0 {
			roots = append(roots, *node)
		} else if parent, ok := nodeMap[m.ParentID]; ok {
			parent.Children = append(parent.Children, *node)
		}
	}

	return roots
}

package types

type RoleVO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Sort        int    `json:"sort"`
}

type CreateRoleReq struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description,optional"`
	Sort        int    `json:"sort,optional"`
}

type UpdateRoleReq struct {
	ID          uint   `json:"id"`
	Name        string `json:"name,optional"`
	Label       string `json:"label,optional"`
	Description string `json:"description,optional"`
	Sort        int    `json:"sort,optional"`
}

type UpdateRoleMenusReq struct {
	ID      uint   `json:"id"`
	MenuIDs []uint `json:"menu_ids"`
}

type MenuVO struct {
	ID       uint     `json:"id"`
	ParentID uint     `json:"parent_id"`
	Name     string   `json:"name"`
	Path     string   `json:"path"`
	Icon     string   `json:"icon"`
	Sort     int      `json:"sort"`
	Children []MenuVO `json:"children"`
}

type CreateMenuReq struct {
	ParentID uint   `json:"parent_id,optional"`
	Name     string `json:"name"`
	Path     string `json:"path,optional"`
	Icon     string `json:"icon,optional"`
	Sort     int    `json:"sort,optional"`
}

type UpdateMenuReq struct {
	ID       uint   `json:"id"`
	ParentID uint   `json:"parent_id,optional"`
	Name     string `json:"name,optional"`
	Path     string `json:"path,optional"`
	Icon     string `json:"icon,optional"`
	Sort     int    `json:"sort,optional"`
}

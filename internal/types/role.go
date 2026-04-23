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

type MenuMetaVO struct {
	Title      string `json:"title"`
	Icon       string `json:"icon"`
	Hidden     bool   `json:"hidden"`
	AlwaysShow bool   `json:"always_show"`
	Affix      bool   `json:"affix"`
	KeepAlive  bool   `json:"keep_alive"`
	Breadcrumb bool   `json:"breadcrumb"`
}

type MenuVO struct {
	ID        uint       `json:"id"`
	ParentID  uint       `json:"parent_id"`
	Name      string     `json:"name"`
	Path      string     `json:"path"`
	Component string     `json:"component"`
	Redirect  string     `json:"redirect"`
	Icon      string     `json:"icon"`
	Sort      int        `json:"sort"`
	Meta      MenuMetaVO `json:"meta"`
	Children  []MenuVO   `json:"children"`
}

type CreateMenuReq struct {
	ParentID  uint   `json:"parent_id,optional"`
	Name      string `json:"name"`
	Title     string `json:"title,optional"`
	Path      string `json:"path,optional"`
	Component string `json:"component,optional"`
	Redirect  string `json:"redirect,optional"`
	Icon      string `json:"icon,optional"`
	Sort      int    `json:"sort,optional"`
	Type      int    `json:"type,optional"`
}

type UpdateMenuReq struct {
	ID        uint   `json:"id"`
	ParentID  uint   `json:"parent_id,optional"`
	Name      string `json:"name,optional"`
	Title     string `json:"title,optional"`
	Path      string `json:"path,optional"`
	Component string `json:"component,optional"`
	Redirect  string `json:"redirect,optional"`
	Icon      string `json:"icon,optional"`
	Sort      int    `json:"sort,optional"`
}

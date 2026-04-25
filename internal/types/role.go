package types

type RoleBackVO struct {
	ID          uint   `json:"id"`
	ParentID    int    `json:"parent_id"`
	RoleKey     string `json:"role_key"`
	RoleLabel   string `json:"role_label"`
	RoleComment string `json:"role_comment"`
	IsDefault   int    `json:"is_default"`
	Status      int    `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type RoleVO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Sort        int    `json:"sort"`
}

type CreateRoleReq struct {
	RoleKey     string `json:"role_key"`
	RoleLabel   string `json:"role_label"`
	RoleComment string `json:"role_comment,optional"`
	IsDefault   int    `json:"is_default,optional"`
	Status      int    `json:"status,optional"`
}

type UpdateRoleReq struct {
	ID          uint   `json:"id"`
	RoleKey     string `json:"role_key,optional"`
	RoleLabel   string `json:"role_label,optional"`
	RoleComment string `json:"role_comment,optional"`
	IsDefault   int    `json:"is_default,optional"`
	Status      int    `json:"status,optional"`
}

type QueryRoleReq struct {
	PageQuery
	RoleKey   string `json:"role_key,optional" form:"role_key,optional"`
	RoleLabel string `json:"role_label,optional" form:"role_label,optional"`
	Status    int    `json:"status,optional" form:"status,optional"`
}

type UpdateRoleMenusReq struct {
	ID      uint   `json:"id"`
	MenuIDs []uint `json:"menu_ids"`
}

type UpdateRoleApisReq struct {
	RoleID uint   `json:"role_id"`
	ApiIDs []uint `json:"api_ids"`
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

type MenuBackVO struct {
	ID        uint         `json:"id"`
	ParentID  uint         `json:"parent_id"`
	Name      string       `json:"name"`
	Path      string       `json:"path"`
	Component string       `json:"component"`
	Redirect  string       `json:"redirect"`
	Meta      MenuMetaVO   `json:"meta"`
	Children  []MenuBackVO `json:"children"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
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

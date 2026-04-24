package types

type PageBackVO struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	PageName  string `json:"page_name"`     // 前端期望的字段名，映射到 title
	Content   string `json:"content"`
	Slug      string `json:"slug"`
	PageLabel string `json:"page_label"`    // 前端期望的字段名，映射到 slug
	Cover     string `json:"cover"`
	PageCover string `json:"page_cover"`    // 前端期望的字段名，映射到 cover
	Sort      int    `json:"sort"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PageVO struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Slug      string `json:"slug"`
	Sort      int    `json:"sort"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreatePageReq struct {
	Title     string `json:"title,optional"`
	PageName  string `json:"page_name,optional"`  // 前端字段名，映射到 title
	Content   string `json:"content,optional"`
	Slug      string `json:"slug,optional"`
	PageLabel string `json:"page_label,optional"` // 前端字段名，映射到 slug
	Cover     string `json:"cover,optional"`
	PageCover string `json:"page_cover,optional"` // 前端字段名，映射到 cover
	Sort      int    `json:"sort,optional"`
	Status    int    `json:"status,optional"`
}

type UpdatePageReq struct {
	ID        uint   `json:"id"`
	Title     string `json:"title,optional"`
	PageName  string `json:"page_name,optional"`  // 前端字段名，映射到 title
	Content   string `json:"content,optional"`
	Slug      string `json:"slug,optional"`
	PageLabel string `json:"page_label,optional"` // 前端字段名，映射到 slug
	Cover     string `json:"cover,optional"`
	PageCover string `json:"page_cover,optional"` // 前端字段名，映射到 cover
	Sort      int    `json:"sort,optional"`
	Status    int    `json:"status,optional"`
}

type QueryPageReq struct {
	PageQuery
	Title    string `json:"title,optional" form:"title,optional"`
	PageName string `json:"page_name,optional" form:"page_name,optional"`
}

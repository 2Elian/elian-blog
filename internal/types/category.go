package types

type CategoryVO struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	Sort         int    `json:"sort"`
	ArticleCount int    `json:"article_count"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type CreateCategoryReq struct {
	CategoryName string `json:"category_name"`
	Description  string `json:"description,optional"`
	Sort         int    `json:"sort,optional"`
}

type UpdateCategoryReq struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name,optional"`
	Description  string `json:"description,optional"`
	Sort         int    `json:"sort,optional"`
}

type QueryCategoryReq struct {
	PageQuery
	CategoryName string `json:"category_name,optional" form:"category_name,optional"`
}

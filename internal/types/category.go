package types

type CategoryVO struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Sort         int    `json:"sort"`
	ArticleCount int    `json:"article_count"`
}

type CreateCategoryReq struct {
	Name        string `json:"name"`
	Description string `json:"description,optional"`
	Sort        int    `json:"sort,optional"`
}

type UpdateCategoryReq struct {
	ID          uint   `json:"id"`
	Name        string `json:"name,optional"`
	Description string `json:"description,optional"`
	Sort        int    `json:"sort,optional"`
}

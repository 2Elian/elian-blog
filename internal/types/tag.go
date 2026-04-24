package types

type TagVO struct {
	ID           uint   `json:"id"`
	TagName      string `json:"tag_name"`
	Color        string `json:"color"`
	ArticleCount int    `json:"article_count"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type CreateTagReq struct {
	TagName string `json:"tag_name"`
	Color   string `json:"color,optional"`
}

type UpdateTagReq struct {
	ID      uint   `json:"id"`
	TagName string `json:"tag_name,optional"`
	Color   string `json:"color,optional"`
}

type QueryTagReq struct {
	PageQuery
	TagName string `json:"tag_name,optional" form:"tag_name,optional"`
}

package types

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
	Title   string `json:"title"`
	Content string `json:"content"`
	Slug    string `json:"slug"`
	Sort    int    `json:"sort,optional"`
	Status  int    `json:"status,optional"`
}

type UpdatePageReq struct {
	ID      uint   `json:"id"`
	Title   string `json:"title,optional"`
	Content string `json:"content,optional"`
	Slug    string `json:"slug,optional"`
	Sort    int    `json:"sort,optional"`
	Status  int    `json:"status,optional"`
}

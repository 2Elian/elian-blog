package types

type TagVO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type CreateTagReq struct {
	Name  string `json:"name"`
	Color string `json:"color,optional"`
}

type UpdateTagReq struct {
	ID    uint   `json:"id"`
	Name  string `json:"name,optional"`
	Color string `json:"color,optional"`
}

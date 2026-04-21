package types

// --- 公共类型 ---

type EmptyReq struct{}

type EmptyResp struct{}

type IDReq struct {
	ID uint `json:"id" path:"id"`
}

type PageQuery struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

type PageResp struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

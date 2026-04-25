package types

type ProductVO struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Content     string  `json:"content"`
	Price       float64 `json:"price"`
	Cover       string  `json:"cover"`
	Status      int     `json:"status"`
	Sort        int     `json:"sort"`
	Type        string  `json:"type"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type CreateProductReq struct {
	Name        string  `json:"name"`
	Description string  `json:"description,optional"`
	Content     string  `json:"content,optional"`
	Price       float64 `json:"price,optional"`
	Cover       string  `json:"cover,optional"`
	Status      int     `json:"status,optional"`
	Sort        int     `json:"sort,optional"`
	Type        string  `json:"type,optional"`
}

type UpdateProductReq struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name,optional"`
	Description string  `json:"description,optional"`
	Content     string  `json:"content,optional"`
	Price       float64 `json:"price,optional"`
	Cover       string  `json:"cover,optional"`
	Status      int     `json:"status,optional"`
	Sort        int     `json:"sort,optional"`
	Type        string  `json:"type,optional"`
}

package model

type Product struct {
	Model
	Name        string  `json:"name" gorm:"size:200;not null"`
	Description string  `json:"description" gorm:"size:500"`
	Content     string  `json:"content" gorm:"type:longtext"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2);default:0"`
	Cover       string  `json:"cover" gorm:"size:500"`
	Status      int     `json:"status" gorm:"default:1;comment:1-上架 0-下架"`
	Sort        int     `json:"sort" gorm:"default:0"`
	Type        string  `json:"type" gorm:"size:50;default:其他"`
}

func (Product) TableName() string { return "product" }

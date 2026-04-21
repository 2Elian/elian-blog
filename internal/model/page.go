package model

// Page 页面表（关于、学习、友链等自定义页面）
type Page struct {
	Model
	Title   string `json:"title" gorm:"size:200;not null"`
	Content string `json:"content" gorm:"type:longtext;not null"`
	Cover   string `json:"cover" gorm:"size:500"`
	Summary string `json:"summary" gorm:"size:500"`
	Slug    string `json:"slug" gorm:"uniqueIndex;size:100;not null;comment:URL标识"`
	Status  int    `json:"status" gorm:"default:1;comment:0-隐藏 1-公开"`
	Sort    int    `json:"sort" gorm:"default:0"`
}

func (Page) TableName() string { return "page" }

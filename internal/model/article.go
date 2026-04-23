package model

// Article 文章表
type Article struct {
	Model
	Title      string   `json:"title" gorm:"size:200;not null;index"`
	Summary    string   `json:"summary" gorm:"size:500"`
	Content    string   `json:"content" gorm:"type:longtext;not null"`
	Cover      string   `json:"cover" gorm:"size:500"`
	CategoryID uint     `json:"category_id" gorm:"index"`
	AuthorID   uint     `json:"author_id" gorm:"index"`
	Status     int      `json:"status" gorm:"default:0;comment:0-草稿 1-公开 2-私密"`
	IsTop      int      `json:"is_top" gorm:"default:0;comment:0-否 1-是"`
	IsOriginal int      `json:"is_original" gorm:"default:1;comment:0-转载 1-原创"`
	SourceURL  string   `json:"source_url" gorm:"size:500"`
	Type       int      `json:"type" gorm:"default:1;comment:1-普通 2-置顶"`
	ViewCount  int      `json:"view_count" gorm:"default:0"`
	LikeCount  int      `json:"like_count" gorm:"default:0"`
	Password   string   `json:"-" gorm:"size:100;comment:访问密码"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"`
	Tags       []Tag    `json:"tags" gorm:"many2many:article_tags;"`
	Author     User     `json:"author" gorm:"foreignKey:AuthorID"`
}

func (Article) TableName() string { return "article" }

// Category 分类表
type Category struct {
	Model
	Name         string `json:"name" gorm:"uniqueIndex;size:50;not null"`
	Description  string `json:"description" gorm:"size:200"`
	Sort         int    `json:"sort" gorm:"default:0"`
	ArticleCount int    `json:"article_count" gorm:"-"`
}

func (Category) TableName() string { return "category" }

// Tag 标签表
type Tag struct {
	Model
	Name         string `json:"name" gorm:"uniqueIndex;size:50;not null"`
	Color        string `json:"color" gorm:"size:20"`
	ArticleCount int    `json:"article_count" gorm:"-"`
}

func (Tag) TableName() string { return "tag" }

package types

// --- 文章 ---

type ArticleHome struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	Summary     string   `json:"summary"`
	Content     string   `json:"content"`
	Cover       string   `json:"cover"`
	CategoryID  uint     `json:"category_id"`
	CategoryName string  `json:"category_name"`
	AuthorID    uint     `json:"author_id"`
	Status      int      `json:"status"`
	IsTop       int      `json:"is_top"`
	IsOriginal  int      `json:"is_original"`
	Type        int      `json:"type"`
	ViewCount   int      `json:"view_count"`
	LikeCount   int      `json:"like_count"`
	TagNameList []string `json:"tag_name_list"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

type ArticleDetails struct {
	ArticleHome
	PrevArticle *ArticlePreview   `json:"prev_article"`
	NextArticle *ArticlePreview   `json:"next_article"`
}

type ArticlePreview struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	ViewCount int    `json:"view_count"`
	CreatedAt string `json:"created_at"`
}

type QueryArticleHomeReq struct {
	PageQuery
	Title      string `json:"title,optional" form:"title,optional"`
	CategoryID uint   `json:"category_id,optional" form:"category_id,optional"`
	TagID      uint   `json:"tag_id,optional" form:"tag_id,optional"`
	Status     int    `json:"status,optional" form:"status,optional"`
}

type CreateArticleReq struct {
	Title      string   `json:"title"`
	Summary    string   `json:"summary,optional"`
	Content    string   `json:"content"`
	Cover      string   `json:"cover,optional"`
	CategoryID uint     `json:"category_id,optional"`
	Status     int      `json:"status,optional"`
	IsTop      int      `json:"is_top,optional"`
	IsOriginal int      `json:"is_original,optional"`
	Type       int      `json:"type,optional"`
	Password   string   `json:"password,optional"`
	TagNames   []string `json:"tag_names,optional"`
}

type UpdateArticleReq struct {
	ID         uint     `json:"id"`
	Title      string   `json:"title,optional"`
	Summary    string   `json:"summary,optional"`
	Content    string   `json:"content,optional"`
	Cover      string   `json:"cover,optional"`
	CategoryID uint     `json:"category_id,optional"`
	Status     int      `json:"status,optional"`
	IsTop      int      `json:"is_top,optional"`
	IsOriginal int      `json:"is_original,optional"`
	Type       int      `json:"type,optional"`
	Password   string   `json:"password,optional"`
	TagNames   []string `json:"tag_names,optional"`
}

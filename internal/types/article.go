package types

// --- 文章 ---

type ArticleBackVO struct {
	ID             uint     `json:"id"`
	ArticleCover   string   `json:"article_cover"`
	ArticleTitle   string   `json:"article_title"`
	ArticleContent string   `json:"article_content"`
	ArticleSummary string   `json:"article_summary"`
	ArticleType    int      `json:"article_type"`
	OriginalURL    string   `json:"original_url"`
	IsTop          int      `json:"is_top"`
	IsDelete       int      `json:"is_delete"`
	Status         int      `json:"status"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
	CategoryName   string   `json:"category_name"`
	TagNameList    []string `json:"tag_name_list"`
	LikeCount      int      `json:"like_count"`
	ViewsCount     int      `json:"views_count"`
}

type ArticleHome struct {
	ID           uint     `json:"id"`
	Title        string   `json:"title"`
	Summary      string   `json:"summary"`
	Content      string   `json:"content"`
	Cover        string   `json:"cover"`
	CategoryID   uint     `json:"category_id"`
	CategoryName string   `json:"category_name"`
	Category     CategoryInfo `json:"category"`
	AuthorID     uint     `json:"author_id"`
	Status       int      `json:"status"`
	IsTop        int      `json:"is_top"`
	IsOriginal   int      `json:"is_original"`
	Type         int      `json:"type"`
	ViewCount    int      `json:"view_count"`
	Views        int      `json:"views"`
	LikeCount    int      `json:"like_count"`
	TagNameList  []string `json:"tag_name_list"`
	Tags         []TagInfo `json:"tags"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
}

type CategoryInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type TagInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ArticleDetails struct {
	ArticleHome
	PrevArticle *ArticlePreview `json:"prev_article"`
	NextArticle *ArticlePreview `json:"next_article"`
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
	Title        string `json:"title,optional" form:"title,optional"`
	ArticleTitle string `json:"article_title,optional" form:"article_title,optional"`
	CategoryID   uint   `json:"category_id,optional" form:"category_id,optional"`
	CategoryName string `json:"category_name,optional" form:"category_name,optional"`
	TagID        uint   `json:"tag_id,optional" form:"tag_id,optional"`
	Status       int    `json:"status,optional" form:"status,optional"`
	ArticleType  int    `json:"article_type,optional" form:"article_type,optional"`
	IsTop        int    `json:"is_top,optional" form:"is_top,optional"`
	IsDelete     int    `json:"is_delete,optional" form:"is_delete,optional"`
}

type CreateArticleReq struct {
	ArticleTitle   string   `json:"article_title"`
	ArticleContent string   `json:"article_content"`
	ArticleCover   string   `json:"article_cover,optional"`
	ArticleSummary string   `json:"article_summary,optional"`
	ArticleType    int      `json:"article_type,optional"`
	OriginalURL    string   `json:"original_url,optional"`
	CategoryName   string   `json:"category_name,optional"`
	TagNameList    []string `json:"tag_name_list,optional"`
	Status         int      `json:"status,optional"`
	IsTop          int      `json:"is_top,optional"`
}

type UpdateArticleReq struct {
	ID             uint     `json:"id"`
	ArticleTitle   string   `json:"article_title,optional"`
	ArticleContent string   `json:"article_content,optional"`
	ArticleCover   string   `json:"article_cover,optional"`
	ArticleSummary string   `json:"article_summary,optional"`
	ArticleType    int      `json:"article_type,optional"`
	OriginalURL    string   `json:"original_url,optional"`
	CategoryName   string   `json:"category_name,optional"`
	TagNameList    []string `json:"tag_name_list,optional"`
	Status         int      `json:"status,optional"`
	IsTop          int      `json:"is_top,optional"`
}

type UpdateArticleTopReq struct {
	ID    uint `json:"id"`
	IsTop int  `json:"is_top"`
}

type UpdateArticleDeleteReq struct {
	ID       uint `json:"id"`
	IsDelete int  `json:"is_delete"`
}

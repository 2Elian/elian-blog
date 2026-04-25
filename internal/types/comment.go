package types

type CommentBackVO struct {
	ID             uint        `json:"id"`
	UserID         uint        `json:"user_id"`
	TerminalID     int         `json:"terminal_id"`
	Type           int         `json:"type"`
	TopicTitle     string      `json:"topic_title"`
	ReplyUserID    uint        `json:"reply_user_id"`
	CommentContent string      `json:"comment_content"`
	Status         int         `json:"status"`
	CreatedAt      string      `json:"created_at"`
	UserInfo       *CommentUser `json:"user_info"`
	ReplyUserInfo  *CommentUser `json:"reply_user_info"`
}

type CommentUser struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}

type CommentVO struct {
	ID        uint   `json:"id"`
	ArticleID uint   `json:"article_id"`
	UserID    uint   `json:"user_id"`
	Content   string `json:"content"`
	Type      int    `json:"type"`
	ParentID  uint   `json:"parent_id"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"created_at"`
}

type CreateCommentReq struct {
	ArticleID uint   `json:"article_id"`
	Content   string `json:"content"`
	Type      int    `json:"type,optional"`
	ParentID  uint   `json:"parent_id,optional"`
}

type QueryCommentReq struct {
	PageQuery
	ArticleID uint `json:"article_id,optional" form:"article_id,optional"`
	Type      int  `json:"type,optional" form:"type,optional"`
	UserID    uint `json:"user_id,optional" form:"user_id,optional"`
	Status    int  `json:"status,optional" form:"status,optional"`
}

type UpdateCommentStatusReq struct {
	IDs    []uint `json:"ids"`
	Status int    `json:"status"`
}

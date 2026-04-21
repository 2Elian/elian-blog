package model

// Comment 评论表
type Comment struct {
	Model
	ArticleID  uint     `json:"article_id" gorm:"index;comment:文章ID,0为留言板"`
	UserID     uint     `json:"user_id" gorm:"index"`
	ParentID   uint     `json:"parent_id" gorm:"default:0;index"`
	ReplyID    uint     `json:"reply_id" gorm:"default:0"`
	Content    string   `json:"content" gorm:"type:text;not null"`
	Type       int      `json:"type" gorm:"default:1;comment:1-文章评论 2-友链评论 3-留言板"`
	IsAdmin    int      `json:"is_admin" gorm:"default:0;comment:0-用户 1-博主"`
	Status     int      `json:"status" gorm:"default:1;comment:0-审核中 1-通过 2-拒绝"`
	IPAddress  string   `json:"ip_address" gorm:"size:50"`
	User       User     `json:"user" gorm:"foreignKey:UserID"`
	ReplyUser  *User    `json:"reply_user" gorm:"foreignKey:ReplyID"`
	Children   []Comment `json:"children" gorm:"-"`
}

func (Comment) TableName() string { return "comment" }

package model

// Message 消息/说说表
type Message struct {
	Model
	UserID    uint   `json:"user_id" gorm:"index"`
	Content   string `json:"content" gorm:"type:text;not null"`
	Images    string `json:"images" gorm:"type:text;comment:图片URL,逗号分隔"`
	Status    int    `json:"status" gorm:"default:1;comment:0-隐藏 1-公开"`
	LikeCount int    `json:"like_count" gorm:"default:0"`
	User      User   `json:"user" gorm:"foreignKey:UserID"`
}

func (Message) TableName() string { return "message" }

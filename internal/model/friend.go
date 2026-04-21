package model

// FriendLink 友链表
type FriendLink struct {
	Model
	Name        string `json:"name" gorm:"size:100;not null"`
	URL         string `json:"url" gorm:"size:500;not null"`
	Logo        string `json:"logo" gorm:"size:500"`
	Description string `json:"description" gorm:"size:200"`
	Sort        int    `json:"sort" gorm:"default:0"`
	Status      int    `json:"status" gorm:"default:1;comment:0-下架 1-正常"`
}

func (FriendLink) TableName() string { return "friend_link" }

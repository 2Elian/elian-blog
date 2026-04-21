package types

type MessageVO struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Content   string `json:"content"`
	Images    string `json:"images"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"created_at"`
}

type CreateMessageReq struct {
	Content string `json:"content"`
	Images  string `json:"images,optional"`
	Type    int    `json:"type,optional"`
}

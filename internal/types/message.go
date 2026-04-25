package types

type MessageBackVO struct {
	ID             uint         `json:"id"`
	UserID         uint         `json:"user_id"`
	TerminalID     int          `json:"terminal_id"`
	MessageContent string       `json:"message_content"`
	Status         int          `json:"status"`
	CreatedAt      string       `json:"created_at"`
	UpdatedAt      string       `json:"updated_at"`
	UserInfo       *MessageUser `json:"user_info"`
}

type MessageUser struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}

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

type QueryMessageReq struct {
	PageQuery
	UserID uint `json:"user_id,optional" form:"user_id,optional"`
	Status int  `json:"status,optional" form:"status,optional"`
}

type UpdateMessageStatusReq struct {
	IDs    []uint `json:"ids"`
	Status int    `json:"status"`
}

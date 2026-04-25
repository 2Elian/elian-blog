package types

type FriendLinkVO struct {
	ID         uint   `json:"id"`
	LinkName   string `json:"link_name"`
	LinkAvatar string `json:"link_avatar"`
	LinkAddress string `json:"link_address"`
	LinkIntro  string `json:"link_intro"`
	Sort       int    `json:"sort"`
	Status     int    `json:"status"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type CreateFriendLinkReq struct {
	LinkName    string `json:"link_name"`
	LinkAvatar  string `json:"link_avatar,optional"`
	LinkAddress string `json:"link_address"`
	LinkIntro   string `json:"link_intro,optional"`
	Sort        int    `json:"sort,optional"`
}

type UpdateFriendLinkReq struct {
	ID          uint   `json:"id"`
	LinkName    string `json:"link_name,optional"`
	LinkAvatar  string `json:"link_avatar,optional"`
	LinkAddress string `json:"link_address,optional"`
	LinkIntro   string `json:"link_intro,optional"`
	Sort        int    `json:"sort,optional"`
	Status      int    `json:"status,optional"`
}

type QueryFriendReq struct {
	PageQuery
	LinkName string `json:"link_name,optional" form:"link_name,optional"`
}

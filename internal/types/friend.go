package types

type FriendLinkVO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	Sort        int    `json:"sort"`
	Status      int    `json:"status"`
}

type CreateFriendLinkReq struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Logo        string `json:"logo,optional"`
	Description string `json:"description,optional"`
	Sort        int    `json:"sort,optional"`
}

type UpdateFriendLinkReq struct {
	ID          uint   `json:"id"`
	Name        string `json:"name,optional"`
	URL         string `json:"url,optional"`
	Logo        string `json:"logo,optional"`
	Description string `json:"description,optional"`
	Sort        int    `json:"sort,optional"`
	Status      int    `json:"status,optional"`
}

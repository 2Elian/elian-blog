package types

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email,optional"`
}

type LoginResp struct {
	Token    string    `json:"token"`
	UserInfo UserInfo  `json:"user_info"`
}

type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

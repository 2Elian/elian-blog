package types

type UserVO struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

type UpdateUserReq struct {
	ID       uint   `json:"id"`
	Username string `json:"username,optional"`
	Avatar   string `json:"avatar,optional"`
	Email    string `json:"email,optional"`
}

type UpdatePasswordReq struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

package types

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar,optional"`
	Intro    string `json:"intro,optional"`
	Website  string `json:"website,optional"`
}

// VeLoginResp ve-admin-element 期望的登录响应
type VeLoginResp struct {
	UserID string   `json:"user_id"`
	Scope  string   `json:"scope"`
	Token  *VeToken `json:"token"`
}

type VeToken struct {
	TokenType        string `json:"token_type"`
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresIn int64  `json:"refresh_expires_in"`
	RefreshExpiresAt int64  `json:"refresh_expires_at"`
}

// VeUserInfoResp ve-admin-element 期望的用户信息
type VeUserInfoResp struct {
	UserID       string   `json:"user_id"`
	Username     string   `json:"username"`
	Nickname     string   `json:"nickname"`
	Avatar       string   `json:"avatar"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	CreatedAt    int64    `json:"created_at"`
	RegisterType string   `json:"register_type"`
	ThirdParty   []any    `json:"third_party"`
	Roles        []string `json:"roles"`
	Perms        []string `json:"perms"`
	Intro        string   `json:"intro"`
	Website      string   `json:"website"`
	Gender       int      `json:"gender"`
}

// 旧格式兼容 (blog API)
type LoginResp struct {
	Token    string   `json:"token"`
	UserInfo UserInfo `json:"user_info"`
}

type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Intro    string `json:"intro"`
	Website  string `json:"website"`
	Role     string `json:"role"`
}

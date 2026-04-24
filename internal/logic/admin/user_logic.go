package admin

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"elian-blog/internal/model"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

type UserLogic struct {
	svcCtx *svc.ServiceContext
}

func NewUserLogic(svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{svcCtx: svcCtx}
}

func (l *UserLogic) List(ctx context.Context, req *types.PageQuery) (interface{}, error) {
	page := req.Page
	pageSize := req.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	users, total, err := l.svcCtx.UserDao.List(page, pageSize)
	if err != nil {
		return nil, err
	}

	type RoleInfo struct {
		RoleKey   string `json:"role_key"`
		RoleLabel string `json:"role_label"`
	}
	type UserBackVO struct {
		UserID       string     `json:"user_id"`
		Username     string     `json:"username"`
		Nickname     string     `json:"nickname"`
		Avatar       string     `json:"avatar"`
		Email        string     `json:"email"`
		Status       int        `json:"status"`
		RegisterType string     `json:"register_type"`
		CreatedAt    string     `json:"created_at"`
		UpdatedAt    string     `json:"updated_at"`
		Roles        []RoleInfo `json:"roles"`
	}

	list := make([]UserBackVO, 0, len(users))
	for _, u := range users {
		roles := make([]RoleInfo, 0, len(u.Roles))
		for _, r := range u.Roles {
			roles = append(roles, RoleInfo{
				RoleKey:   r.Label,
				RoleLabel: r.Name,
			})
		}
		avatar := u.Avatar
		if avatar == "" || avatar == "https://example.com/avatar.png" || (!strings.HasPrefix(avatar, "http") && !strings.HasPrefix(avatar, "/")) {
			avatar = "https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
		}
		list = append(list, UserBackVO{
			UserID:       strconv.FormatUint(uint64(u.ID), 10),
			Username:     u.Username,
			Nickname:     u.Nickname,
			Avatar:       avatar,
			Email:        u.Email,
			Status:       u.Status,
			RegisterType: u.LoginType,
			CreatedAt:    formatTime(u.CreatedAt),
			UpdatedAt:    formatTime(u.UpdatedAt),
			Roles:        roles,
		})
	}

	return types.PageResp{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

func (l *UserLogic) Delete(ctx context.Context, id uint) error {
	return l.svcCtx.UserDao.Delete(id)
}

// formatTime is defined in category_logic.go in the same package.
// If not present, define it here:
func init() {
	_ = fmt.Sprintf
	_ = model.Article{}
}

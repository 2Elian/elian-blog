package admin

import (
	"context"
	"errors"
	"strconv"
	"time"

	"elian-blog/internal/middleware"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"
	"elian-blog/internal/utils"

	"gorm.io/gorm"
)

type VeAuthLogic struct {
	svcCtx *svc.ServiceContext
}

func NewVeAuthLogic(svcCtx *svc.ServiceContext) *VeAuthLogic {
	return &VeAuthLogic{svcCtx: svcCtx}
}

// VeLogin ve-admin-element 登录
func (l *VeAuthLogic) VeLogin(ctx context.Context, req *types.LoginReq) (*types.VeLoginResp, error) {
	user, err := l.svcCtx.UserDao.GetByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户名或密码错误")
		}
		return nil, err
	}

	if user.Status != 1 {
		return nil, errors.New("账号已被禁用")
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("用户名或密码错误")
	}

	role := "user"
	if len(user.Roles) > 0 {
		role = user.Roles[0].Label
	}

	token, err := utils.GenerateToken(user.ID, user.Username, role, l.svcCtx.Config.JWT.Secret, l.svcCtx.Config.JWT.ExpireHours)
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	_ = l.svcCtx.UserDao.UpdateLoginTime(user.ID)

	expiresIn := int64(l.svcCtx.Config.JWT.ExpireHours * 3600)
	now := time.Now()

	return &types.VeLoginResp{
		UserID: strconv.FormatUint(uint64(user.ID), 10),
		Scope:  role,
		Token: &types.VeToken{
			TokenType:        "Bearer",
			AccessToken:      token,
			ExpiresIn:        expiresIn,
			RefreshToken:     token,
			RefreshExpiresIn: expiresIn,
			RefreshExpiresAt: now.Add(time.Duration(expiresIn) * time.Second).Unix(),
		},
	}, nil
}

// VeGetUserInfo 获取用户信息
func (l *VeAuthLogic) VeGetUserInfo(ctx context.Context) (*types.VeUserInfoResp, error) {
	userID := getUserIDFromCtx(ctx)
	if userID == 0 {
		return nil, errors.New("未获取到用户信息")
	}

	user, err := l.svcCtx.UserDao.GetByID(userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	roles := make([]string, 0, len(user.Roles))
	for _, r := range user.Roles {
		roles = append(roles, r.Label)
	}

	perms := []string{"*"}

	return &types.VeUserInfoResp{
		UserID:       strconv.FormatUint(uint64(user.ID), 10),
		Username:     user.Username,
		Nickname:     user.Nickname,
		Avatar:       user.Avatar,
		Email:        user.Email,
		Phone:        "",
		CreatedAt:    user.CreatedAt.Unix(),
		RegisterType: user.LoginType,
		ThirdParty:   []any{},
		Roles:        roles,
		Perms:        perms,
		Intro:        user.Intro,
		Website:      user.Website,
	}, nil
}

// VeGetUserMenus 获取用户菜单
func (l *VeAuthLogic) VeGetUserMenus(ctx context.Context) (interface{}, error) {
	userID := getUserIDFromCtx(ctx)
	if userID == 0 {
		return nil, errors.New("未获取到用户信息")
	}

	menuLogic := NewMenuLogic(l.svcCtx)
	return menuLogic.List(ctx)
}

// VeGetUserRoles 获取用户角色列表
func (l *VeAuthLogic) VeGetUserRoles(ctx context.Context) (interface{}, error) {
	userID := getUserIDFromCtx(ctx)
	if userID == 0 {
		return nil, errors.New("未获取到用户信息")
	}

	user, err := l.svcCtx.UserDao.GetByID(userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	type UserRole struct {
		ID          uint   `json:"id"`
		ParentID    int    `json:"parent_id"`
		RoleKey     string `json:"role_key"`
		RoleLabel   string `json:"role_label"`
		RoleComment string `json:"role_comment"`
	}

	roles := make([]UserRole, 0, len(user.Roles))
	for _, r := range user.Roles {
		roles = append(roles, UserRole{
			ID:          r.ID,
			ParentID:    0,
			RoleKey:     r.Label,
			RoleLabel:   r.Name,
			RoleComment: r.Description,
		})
	}

	return map[string]interface{}{"list": roles}, nil
}

// VeGetUserApis 获取用户API权限
func (l *VeAuthLogic) VeGetUserApis(ctx context.Context) (interface{}, error) {
	userID := getUserIDFromCtx(ctx)
	if userID == 0 {
		return nil, errors.New("未获取到用户信息")
	}

	user, err := l.svcCtx.UserDao.GetByID(userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	isAdmin := false
	for _, r := range user.Roles {
		if r.Label == "admin" || r.Label == "root" {
			isAdmin = true
			break
		}
	}

	if isAdmin {
		return map[string]interface{}{"list": []map[string]interface{}{
			{"id": 1, "parent_id": 0, "name": "全部", "path": "*", "method": "*"},
		}}, nil
	}

	return map[string]interface{}{"list": []map[string]interface{}{}}, nil
}

func getUserIDFromCtx(ctx context.Context) uint {
	return middleware.GetUserID(ctx)
}

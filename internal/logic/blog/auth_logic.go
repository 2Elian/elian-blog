package blog

import (
	"context"
	"errors"

	"elian-blog/internal/model"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"
	"elian-blog/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type AuthLogic struct {
	svcCtx *svc.ServiceContext
}

func NewAuthLogic(svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{svcCtx: svcCtx}
}

// Login 用户登录
func (l *AuthLogic) Login(ctx context.Context, req *types.LoginReq) (interface{}, error) {
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

	// 更新登录时间
	_ = l.svcCtx.UserDao.UpdateLoginTime(user.ID)

	return types.LoginResp{
		Token: token,
		UserInfo: types.UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Avatar:   user.Avatar,
			Email:    user.Email,
			Intro:    user.Intro,
			Website:  user.Website,
			Role:     role,
		},
	}, nil
}

// Register 用户注册
func (l *AuthLogic) Register(ctx context.Context, req *types.RegisterReq) (interface{}, error) {
	// 检查用户名是否已存在
	_, err := l.svcCtx.UserDao.GetByUsername(req.Username)
	if err == nil {
		return nil, errors.New("用户名已存在")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 密码加密
	hashedPwd, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	user := &model.User{
		Username: req.Username,
		Password: hashedPwd,
		Nickname: req.Username,
		Email:    req.Email,
		Avatar:   req.Avatar,
		Intro:    req.Intro,
		Website:  req.Website,
		Status:   1,
	}

	if err := l.svcCtx.UserDao.Create(user); err != nil {
		return nil, err
	}

	// 查找或创建默认 "user" 角色
	defaultRole, err := l.svcCtx.RoleDao.FindByLabel("user")
	if err != nil {
		defaultRole = &model.Role{
			Name:        "普通用户",
			Label:       "user",
			Description: "默认用户角色",
			Status:      1,
			Sort:        3,
		}
		if createErr := l.svcCtx.RoleDao.Create(defaultRole); createErr != nil {
			logx.Errorf("创建默认角色失败: %v", createErr)
		}
	}

	// 创建 user_roles 关联
	if defaultRole != nil && defaultRole.ID > 0 {
		if assignErr := l.svcCtx.UserDao.AssignRole(user.ID, defaultRole.ID); assignErr != nil {
			logx.Errorf("分配角色失败: %v", assignErr)
		}
	}

	// 生成token
	token, err := utils.GenerateToken(user.ID, user.Username, "user", l.svcCtx.Config.JWT.Secret, l.svcCtx.Config.JWT.ExpireHours)
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	return types.LoginResp{
		Token: token,
		UserInfo: types.UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Avatar:   user.Avatar,
			Email:    user.Email,
			Intro:    user.Intro,
			Website:  user.Website,
			Role:     "user",
		},
	}, nil
}

// GetUserInfo 获取用户信息
func (l *AuthLogic) GetUserInfo(ctx context.Context) (interface{}, error) {
	// 从上下文中获取用户ID
	userID, exists := ctx.Value("user_id").(uint)
	if !exists || userID == 0 {
		return nil, errors.New("未获取到用户信息")
	}

	user, err := l.svcCtx.UserDao.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	role := "user"
	if len(user.Roles) > 0 {
		role = user.Roles[0].Label
	}

	return types.UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Intro:    user.Intro,
		Website:  user.Website,
		Role:     role,
	}, nil
}

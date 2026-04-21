package service

import (
	"elian-blog/internal/dao"
	"elian-blog/internal/model"
	"elian-blog/internal/utils"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthService struct {
	userDao *dao.UserDao
	secret  string
	expire  int
	log     *zap.Logger
}

func NewAuthService(userDao *dao.UserDao, secret string, expireHours int, log *zap.Logger) *AuthService {
	return &AuthService{
		userDao: userDao,
		secret:  secret,
		expire:  expireHours,
		log:     log,
	}
}

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterReq struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type LoginResp struct {
	Token    string `json:"token"`
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

func (s *AuthService) Login(req *LoginReq) (*LoginResp, error) {
	user, err := s.userDao.GetByUsername(req.Username)
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

	token, err := utils.GenerateToken(user.ID, user.Username, role, s.secret, s.expire)
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	_ = s.userDao.UpdateLoginTime(user.ID)

	return &LoginResp{
		Token:    token,
		UserID:   user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}, nil
}

func (s *AuthService) Register(req *RegisterReq) (*LoginResp, error) {
	_, err := s.userDao.GetByUsername(req.Username)
	if err == nil {
		return nil, errors.New("用户名已存在")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hashedPwd, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	nickname := req.Nickname
	if nickname == "" {
		// TODO 如果用户不传自己的名称 那么nickname = elian-uuid
		nickname = req.Username
	}
	// TODO 邮箱校验 然后发验证码
	user := &model.User{
		Username:  req.Username,
		Password:  hashedPwd,
		Nickname:  nickname,
		Email:     req.Email,
		Status:    1,
		LastLogin: time.Now(),
	}

	if err := s.userDao.Create(user); err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(user.ID, user.Username, "user", s.secret, s.expire)
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	return &LoginResp{
		Token:    token,
		UserID:   user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
	}, nil
}

func (s *AuthService) GetUserInfo(userID uint) (*model.User, error) {
	return s.userDao.GetByID(userID)
}

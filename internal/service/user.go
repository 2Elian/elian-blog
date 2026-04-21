package service

import (
	"elian-blog/internal/dao"
	"elian-blog/internal/model"
	"elian-blog/internal/utils"
	"errors"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService(userDao *dao.UserDao) *UserService {
	return &UserService{userDao: userDao}
}

func (s *UserService) GetByID(id uint) (*model.User, error) {
	return s.userDao.GetByID(id)
}

func (s *UserService) List(page, pageSize int) ([]model.User, int64, error) {
	return s.userDao.List(page, pageSize)
}

func (s *UserService) Update(user *model.User) error {
	return s.userDao.Update(user)
}

func (s *UserService) Delete(id uint) error {
	return s.userDao.Delete(id)
}

type UserUpdateReq struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Intro    string `json:"intro"`
	Website  string `json:"website"`
}

func (s *UserService) UpdateProfile(userID uint, req *UserUpdateReq) error {
	user, err := s.userDao.GetByID(userID)
	if err != nil {
		return err
	}
	user.Nickname = req.Nickname
	user.Avatar = req.Avatar
	user.Email = req.Email
	user.Intro = req.Intro
	user.Website = req.Website
	return s.userDao.Update(user)
}

func (s *UserService) UpdatePassword(userID uint, oldPwd, newPwd string) error {
	user, err := s.userDao.GetByID(userID)
	if err != nil {
		return err
	}
	if !utils.CheckPassword(oldPwd, user.Password) {
		return errors.New("旧密码错误")
	}
	hashed, err := utils.HashPassword(newPwd)
	if err != nil {
		return err
	}
	user.Password = hashed
	return s.userDao.Update(user)
}

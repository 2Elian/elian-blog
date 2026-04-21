package service

import (
	"elian-blog/internal/dao"
	"elian-blog/internal/model"
)

type FriendLinkService struct {
	dao *dao.FriendLinkDao
}

func NewFriendLinkService(dao *dao.FriendLinkDao) *FriendLinkService {
	return &FriendLinkService{dao: dao}
}

func (s *FriendLinkService) Create(name, url, logo, desc string, sort int) (*model.FriendLink, error) {
	link := &model.FriendLink{
		Name:        name,
		URL:         url,
		Logo:        logo,
		Description: desc,
		Sort:        sort,
		Status:      1,
	}
	return link, s.dao.Create(link)
}

func (s *FriendLinkService) Update(link *model.FriendLink) error {
	return s.dao.Update(link)
}

func (s *FriendLinkService) Delete(id uint) error {
	return s.dao.Delete(id)
}

func (s *FriendLinkService) GetByID(id uint) (*model.FriendLink, error) {
	return s.dao.GetByID(id)
}

func (s *FriendLinkService) List() ([]model.FriendLink, error) {
	return s.dao.List()
}

func (s *FriendLinkService) ListAdmin() ([]model.FriendLink, error) {
	return s.dao.ListAdmin()
}

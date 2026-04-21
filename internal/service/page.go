package service

import (
	"elian-blog/internal/dao"
	"elian-blog/internal/model"
)

type PageService struct {
	dao *dao.PageDao
}

func NewPageService(dao *dao.PageDao) *PageService {
	return &PageService{dao: dao}
}

func (s *PageService) Create(page *model.Page) error {
	return s.dao.Create(page)
}

func (s *PageService) Update(page *model.Page) error {
	return s.dao.Update(page)
}

func (s *PageService) Delete(id uint) error {
	return s.dao.Delete(id)
}

func (s *PageService) GetByID(id uint) (*model.Page, error) {
	return s.dao.GetByID(id)
}

func (s *PageService) GetBySlug(slug string) (*model.Page, error) {
	return s.dao.GetBySlug(slug)
}

func (s *PageService) List() ([]model.Page, error) {
	return s.dao.List()
}

func (s *PageService) ListAdmin() ([]model.Page, error) {
	return s.dao.ListAdmin()
}

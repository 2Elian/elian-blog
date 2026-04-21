package service

import (
	"elian-blog/internal/dao"
	"elian-blog/internal/model"
)

type CategoryService struct {
	dao *dao.CategoryDao
}

func NewCategoryService(dao *dao.CategoryDao) *CategoryService {
	return &CategoryService{dao: dao}
}

func (s *CategoryService) Create(name, desc string, sort int) (*model.Category, error) {
	cat := &model.Category{Name: name, Description: desc, Sort: sort}
	return cat, s.dao.Create(cat)
}

func (s *CategoryService) Update(cat *model.Category) error {
	return s.dao.Update(cat)
}

func (s *CategoryService) Delete(id uint) error {
	return s.dao.Delete(id)
}

func (s *CategoryService) GetByID(id uint) (*model.Category, error) {
	return s.dao.GetByID(id)
}

func (s *CategoryService) List() ([]model.Category, error) {
	cats, err := s.dao.List()
	if err != nil {
		return nil, err
	}
	for i := range cats {
		count, _ := s.dao.CountArticles(cats[i].ID)
		cats[i].ArticleCount = int(count)
	}
	return cats, nil
}

type TagService struct {
	dao *dao.TagDao
}

func NewTagService(dao *dao.TagDao) *TagService {
	return &TagService{dao: dao}
}

func (s *TagService) Create(name, color string) (*model.Tag, error) {
	tag := &model.Tag{Name: name, Color: color}
	return tag, s.dao.Create(tag)
}

func (s *TagService) Update(tag *model.Tag) error {
	return s.dao.Update(tag)
}

func (s *TagService) Delete(id uint) error {
	return s.dao.Delete(id)
}

func (s *TagService) GetByID(id uint) (*model.Tag, error) {
	return s.dao.GetByID(id)
}

func (s *TagService) List() ([]model.Tag, error) {
	return s.dao.List()
}

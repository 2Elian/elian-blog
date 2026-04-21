package service

import (
	"elian-blog/internal/dao"
	"elian-blog/internal/model"
)

type MessageService struct {
	dao *dao.MessageDao
}

func NewMessageService(dao *dao.MessageDao) *MessageService {
	return &MessageService{dao: dao}
}

type MessageCreateReq struct {
	Content string `json:"content" binding:"required"`
	Images  string `json:"images"`
}

func (s *MessageService) Create(userID uint, req *MessageCreateReq) error {
	msg := &model.Message{
		UserID:  userID,
		Content: req.Content,
		Images:  req.Images,
		Status:  1,
	}
	return s.dao.Create(msg)
}

func (s *MessageService) Delete(id uint) error {
	return s.dao.Delete(id)
}

func (s *MessageService) List(page, pageSize int) ([]model.Message, int64, error) {
	return s.dao.List(page, pageSize)
}

func (s *MessageService) ListAdmin(page, pageSize int) ([]model.Message, int64, error) {
	return s.dao.ListAdmin(page, pageSize)
}
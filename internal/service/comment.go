package service

import (
	"elian-blog/internal/dao"
	"elian-blog/internal/model"
)

type CommentService struct {
	dao *dao.CommentDao
}

func NewCommentService(dao *dao.CommentDao) *CommentService {
	return &CommentService{dao: dao}
}

type CommentCreateReq struct {
	ArticleID uint   `json:"article_id"`
	ParentID  uint   `json:"parent_id"`
	ReplyID   uint   `json:"reply_id"`
	Content   string `json:"content" binding:"required"`
	Type      int    `json:"type"`
}

func (s *CommentService) Create(userID uint, req *CommentCreateReq) error {
	comment := &model.Comment{
		ArticleID: req.ArticleID,
		UserID:    userID,
		ParentID:  req.ParentID,
		ReplyID:   req.ReplyID,
		Content:   req.Content,
		Type:      req.Type,
		Status:    1,
	}
	return s.dao.Create(comment)
}

func (s *CommentService) Delete(id uint) error {
	return s.dao.Delete(id)
}

func (s *CommentService) UpdateStatus(id uint, status int) error {
	return s.dao.UpdateStatus(id, status)
}

func (s *CommentService) ListByArticle(articleID uint, page, pageSize int) ([]model.Comment, int64, error) {
	return s.dao.ListByArticle(articleID, page, pageSize)
}

func (s *CommentService) ListRecent(limit int) ([]model.Comment, error) {
	return s.dao.ListRecent(limit)
}

func (s *CommentService) ListAdmin(page, pageSize int) ([]model.Comment, int64, error) {
	return s.dao.ListAdmin(page, pageSize)
}

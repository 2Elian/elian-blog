package dao

import (
	"elian-blog/internal/model"

	"gorm.io/gorm"
)

type CommentDao struct {
	db *gorm.DB
}

func NewCommentDao(db *gorm.DB) *CommentDao {
	return &CommentDao{db: db}
}

func (d *CommentDao) Create(comment *model.Comment) error {
	return d.db.Create(comment).Error
}

func (d *CommentDao) Delete(id uint) error {
	return d.db.Delete(&model.Comment{}, id).Error
}

func (d *CommentDao) UpdateStatus(id uint, status int) error {
	return d.db.Model(&model.Comment{}).Where("id = ?", id).Update("status", status).Error
}

func (d *CommentDao) ListByArticle(articleID uint, page, pageSize int) ([]model.Comment, int64, error) {
	var comments []model.Comment
	var total int64

	query := d.db.Model(&model.Comment{}).Where("article_id = ? AND status = 1", articleID)
	query.Count(&total)

	err := query.Preload("User").Preload("ReplyUser").
		Order("created_at DESC").
		Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&comments).Error
	return comments, total, err
}

func (d *CommentDao) ListRecent(limit int) ([]model.Comment, error) {
	var comments []model.Comment
	err := d.db.Preload("User").Where("status = 1").
		Order("created_at DESC").Limit(limit).Find(&comments).Error
	return comments, err
}

func (d *CommentDao) ListAdmin(page, pageSize int) ([]model.Comment, int64, error) {
	var comments []model.Comment
	var total int64
	d.db.Model(&model.Comment{}).Count(&total)
	err := d.db.Preload("User").Offset((page - 1) * pageSize).Limit(pageSize).
		Order("created_at DESC").Find(&comments).Error
	return comments, total, err
}

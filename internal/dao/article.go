package dao

import (
	"elian-blog/internal/model"

	"gorm.io/gorm"
)

type ArticleDao struct {
	db *gorm.DB
}

func NewArticleDao(db *gorm.DB) *ArticleDao {
	return &ArticleDao{db: db}
}

func (d *ArticleDao) Create(article *model.Article) error {
	return d.db.Create(article).Error
}

func (d *ArticleDao) Update(article *model.Article) error {
	return d.db.Save(article).Error
}

func (d *ArticleDao) Delete(id uint) error {
	return d.db.Select("Tags").Delete(&model.Article{}, id).Error
}

func (d *ArticleDao) GetByID(id uint) (*model.Article, error) {
	var article model.Article
	err := d.db.Preload("Category").Preload("Tags").Preload("Author").First(&article, id).Error
	return &article, err
}

func (d *ArticleDao) List(page, pageSize int, status int, categoryID uint, tagID uint) ([]model.Article, int64, error) {
	var articles []model.Article
	var total int64

	query := d.db.Model(&model.Article{})
	if status >= 0 {
		query = query.Where("status = ?", status)
	}
	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}
	if tagID > 0 {
		query = query.Joins("JOIN article_tags ON article_tags.article_id = article.id").
			Where("article_tags.tag_id = ?", tagID)
	}

	query.Count(&total)
	err := query.Preload("Category").Preload("Tags").Preload("Author").
		Order("is_top DESC, created_at DESC").
		Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&articles).Error
	return articles, total, err
}

func (d *ArticleDao) IncrementViewCount(id uint) error {
	return d.db.Model(&model.Article{}).Where("id = ?", id).
		UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error
}

func (d *ArticleDao) UpdateTags(article *model.Article, tags []model.Tag) error {
	return d.db.Model(article).Association("Tags").Replace(tags)
}

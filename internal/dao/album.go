package dao

import (
	"elian-blog/internal/model"

	"gorm.io/gorm"
)

type AlbumDao struct {
	db *gorm.DB
}

func NewAlbumDao(db *gorm.DB) *AlbumDao {
	return &AlbumDao{db: db}
}

func (d *AlbumDao) Create(album *model.Album) error {
	return d.db.Create(album).Error
}

func (d *AlbumDao) Update(album *model.Album) error {
	return d.db.Save(album).Error
}

func (d *AlbumDao) Delete(id uint) error {
	return d.db.Delete(&model.Album{}, id).Error
}

func (d *AlbumDao) GetByID(id uint) (*model.Album, error) {
	var album model.Album
	err := d.db.First(&album, id).Error
	return &album, err
}

func (d *AlbumDao) List(page, pageSize int, albumName string, isDelete int) ([]model.Album, int64, error) {
	var albums []model.Album
	var total int64

	query := d.db.Model(&model.Album{})
	if albumName != "" {
		query = query.Where("album_name LIKE ?", "%"+albumName+"%")
	}
	if isDelete >= 0 {
		query = query.Where("status != 0 OR id > 0") // show all
	}

	query.Count(&total)
	err := query.Order("id DESC").
		Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&albums).Error
	return albums, total, err
}

func (d *AlbumDao) CountPhotos(albumID uint) (int64, error) {
	var count int64
	err := d.db.Model(&model.Photo{}).Where("album_id = ?", albumID).Count(&count).Error
	return count, err
}

type PhotoDao struct {
	db *gorm.DB
}

func NewPhotoDao(db *gorm.DB) *PhotoDao {
	return &PhotoDao{db: db}
}

func (d *PhotoDao) Create(photo *model.Photo) error {
	return d.db.Create(photo).Error
}

func (d *PhotoDao) Update(photo *model.Photo) error {
	return d.db.Save(photo).Error
}

func (d *PhotoDao) Delete(id uint) error {
	return d.db.Delete(&model.Photo{}, id).Error
}

func (d *PhotoDao) GetByID(id uint) (*model.Photo, error) {
	var photo model.Photo
	err := d.db.First(&photo, id).Error
	return &photo, err
}

func (d *PhotoDao) List(page, pageSize int, albumID uint, isDelete int) ([]model.Photo, int64, error) {
	var photos []model.Photo
	var total int64

	query := d.db.Model(&model.Photo{})
	if albumID > 0 {
		query = query.Where("album_id = ?", albumID)
	}

	query.Count(&total)
	err := query.Order("id DESC").
		Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&photos).Error
	return photos, total, err
}

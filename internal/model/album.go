package model

type Album struct {
	Model
	AlbumName string `json:"album_name" gorm:"size:100;not null;comment:相册名称"`
	AlbumDesc string `json:"album_desc" gorm:"size:500;comment:相册描述"`
	AlbumCover string `json:"album_cover" gorm:"size:500;comment:相册封面"`
	Status    int    `json:"status" gorm:"default:1;comment:1-公开 2-私密"`
}

func (Album) TableName() string { return "album" }

type Photo struct {
	Model
	AlbumID   uint   `json:"album_id" gorm:"index;not null;comment:相册ID"`
	PhotoName string `json:"photo_name" gorm:"size:200;not null;comment:照片名称"`
	PhotoDesc string `json:"photo_desc" gorm:"size:500;comment:照片描述"`
	PhotoSrc  string `json:"photo_src" gorm:"size:500;not null;comment:照片地址"`
}

func (Photo) TableName() string { return "photo" }

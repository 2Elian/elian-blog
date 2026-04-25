package types

type AlbumBackVO struct {
	ID         uint   `json:"id"`
	AlbumName  string `json:"album_name"`
	AlbumDesc  string `json:"album_desc"`
	AlbumCover string `json:"album_cover"`
	IsDelete   int    `json:"is_delete"`
	Status     int    `json:"status"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	PhotoCount int    `json:"photo_count"`
}

type NewAlbumReq struct {
	ID         uint   `json:"id"`
	AlbumName  string `json:"album_name"`
	AlbumDesc  string `json:"album_desc"`
	AlbumCover string `json:"album_cover"`
	Status     int    `json:"status"`
	IsDelete   int    `json:"is_delete"`
}

type QueryAlbumReq struct {
	PageQuery
	AlbumName string `json:"album_name,optional" form:"album_name,optional"`
	IsDelete  int    `json:"is_delete,optional" form:"is_delete,optional"`
}

type UpdateAlbumDeleteReq struct {
	IDs      []uint `json:"ids"`
	IsDelete int    `json:"is_delete"`
}

type PhotoBackVO struct {
	ID        uint   `json:"id"`
	AlbumID   uint   `json:"album_id"`
	PhotoName string `json:"photo_name"`
	PhotoDesc string `json:"photo_desc"`
	PhotoSrc  string `json:"photo_src"`
	IsDelete  int    `json:"is_delete"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type NewPhotoReq struct {
	ID        uint   `json:"id"`
	AlbumID   uint   `json:"album_id"`
	PhotoName string `json:"photo_name"`
	PhotoDesc string `json:"photo_desc"`
	PhotoSrc  string `json:"photo_src"`
	IsDelete  int    `json:"is_delete"`
}

type QueryPhotoReq struct {
	PageQuery
	AlbumID  uint `json:"album_id,optional" form:"album_id,optional"`
	IsDelete int  `json:"is_delete,optional" form:"is_delete,optional"`
}

type UpdatePhotoDeleteReq struct {
	IDs      []uint `json:"ids"`
	IsDelete int    `json:"is_delete"`
}

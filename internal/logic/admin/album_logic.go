package admin

import (
	"context"
	"strings"

	"elian-blog/internal/model"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

func (l *AlbumLogic) fixAlbumCoverURL(cover string) string {
	if cover == "" {
		return ""
	}
	if strings.HasPrefix(cover, "http") {
		return cover
	}
	baseURL := l.svcCtx.Config.Upload.BaseURL
	if strings.HasPrefix(cover, "/") {
		return baseURL + cover
	}
	return baseURL + "/" + cover
}

type AlbumLogic struct {
	svcCtx *svc.ServiceContext
}

func NewAlbumLogic(svcCtx *svc.ServiceContext) *AlbumLogic {
	return &AlbumLogic{svcCtx: svcCtx}
}

func (l *AlbumLogic) List(ctx context.Context, req *types.QueryAlbumReq) (interface{}, int64, error) {
	page := req.Page
	pageSize := req.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	albums, total, err := l.svcCtx.AlbumDao.List(page, pageSize, req.AlbumName, req.IsDelete)
	if err != nil {
		return nil, 0, err
	}

	list := make([]types.AlbumBackVO, 0, len(albums))
	for _, album := range albums {
		count, _ := l.svcCtx.AlbumDao.CountPhotos(album.ID)
		list = append(list, types.AlbumBackVO{
			ID:         album.ID,
			AlbumName:  album.AlbumName,
			AlbumDesc:  album.AlbumDesc,
			AlbumCover: l.fixAlbumCoverURL(album.AlbumCover),
			Status:     album.Status,
			CreatedAt:  formatTime(album.CreatedAt),
			UpdatedAt:  formatTime(album.UpdatedAt),
			PhotoCount: int(count),
		})
	}

	return list, total, nil
}

func (l *AlbumLogic) Get(ctx context.Context, id uint) (interface{}, error) {
	album, err := l.svcCtx.AlbumDao.GetByID(id)
	if err != nil {
		return nil, err
	}
	count, _ := l.svcCtx.AlbumDao.CountPhotos(album.ID)
	return types.AlbumBackVO{
		ID:         album.ID,
		AlbumName:  album.AlbumName,
		AlbumDesc:  album.AlbumDesc,
		AlbumCover: l.fixAlbumCoverURL(album.AlbumCover),
		Status:     album.Status,
		CreatedAt:  formatTime(album.CreatedAt),
		UpdatedAt:  formatTime(album.UpdatedAt),
		PhotoCount: int(count),
	}, nil
}

func (l *AlbumLogic) Create(ctx context.Context, req *types.NewAlbumReq) (interface{}, error) {
	album := &model.Album{
		AlbumName:  req.AlbumName,
		AlbumDesc:  req.AlbumDesc,
		AlbumCover: req.AlbumCover,
		Status:     req.Status,
	}
	if album.Status == 0 {
		album.Status = 1
	}
	if err := l.svcCtx.AlbumDao.Create(album); err != nil {
		return nil, err
	}
	return types.AlbumBackVO{
		ID:         album.ID,
		AlbumName:  album.AlbumName,
		AlbumDesc:  album.AlbumDesc,
		AlbumCover: l.fixAlbumCoverURL(album.AlbumCover),
		Status:     album.Status,
		CreatedAt:  formatTime(album.CreatedAt),
	}, nil
}

func (l *AlbumLogic) Update(ctx context.Context, req *types.NewAlbumReq) error {
	album, err := l.svcCtx.AlbumDao.GetByID(req.ID)
	if err != nil {
		return err
	}
	album.AlbumName = req.AlbumName
	album.AlbumDesc = req.AlbumDesc
	album.AlbumCover = req.AlbumCover
	album.Status = req.Status
	return l.svcCtx.AlbumDao.Update(album)
}

func (l *AlbumLogic) Delete(ctx context.Context, ids []uint) error {
	for _, id := range ids {
		_ = l.svcCtx.AlbumDao.Delete(id)
	}
	return nil
}

func (l *AlbumLogic) UpdateDeleteStatus(ctx context.Context, req *types.UpdateAlbumDeleteReq) error {
	for _, id := range req.IDs {
		l.svcCtx.DB.Model(&model.Album{}).Where("id = ?", id).Update("status", 0)
	}
	return nil
}

type PhotoLogic struct {
	svcCtx *svc.ServiceContext
}

func NewPhotoLogic(svcCtx *svc.ServiceContext) *PhotoLogic {
	return &PhotoLogic{svcCtx: svcCtx}
}

func (l *PhotoLogic) List(ctx context.Context, req *types.QueryPhotoReq) (interface{}, int64, error) {
	page := req.Page
	pageSize := req.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}

	photos, total, err := l.svcCtx.PhotoDao.List(page, pageSize, req.AlbumID, req.IsDelete)
	if err != nil {
		return nil, 0, err
	}

	list := make([]types.PhotoBackVO, 0, len(photos))
	for _, photo := range photos {
		list = append(list, types.PhotoBackVO{
			ID:        photo.ID,
			AlbumID:   photo.AlbumID,
			PhotoName: photo.PhotoName,
			PhotoDesc: photo.PhotoDesc,
			PhotoSrc:  photo.PhotoSrc,
			CreatedAt: formatTime(photo.CreatedAt),
			UpdatedAt: formatTime(photo.UpdatedAt),
		})
	}

	return list, total, nil
}

func (l *PhotoLogic) Create(ctx context.Context, req *types.NewPhotoReq) (interface{}, error) {
	photo := &model.Photo{
		AlbumID:   req.AlbumID,
		PhotoName: req.PhotoName,
		PhotoDesc: req.PhotoDesc,
		PhotoSrc:  req.PhotoSrc,
	}
	if err := l.svcCtx.PhotoDao.Create(photo); err != nil {
		return nil, err
	}
	return types.PhotoBackVO{
		ID:        photo.ID,
		AlbumID:   photo.AlbumID,
		PhotoName: photo.PhotoName,
		PhotoDesc: photo.PhotoDesc,
		PhotoSrc:  photo.PhotoSrc,
		CreatedAt: formatTime(photo.CreatedAt),
	}, nil
}

func (l *PhotoLogic) Update(ctx context.Context, req *types.NewPhotoReq) error {
	photo, err := l.svcCtx.PhotoDao.GetByID(req.ID)
	if err != nil {
		return err
	}
	photo.PhotoName = req.PhotoName
	photo.PhotoDesc = req.PhotoDesc
	return l.svcCtx.PhotoDao.Update(photo)
}

func (l *PhotoLogic) Delete(ctx context.Context, ids []uint) error {
	for _, id := range ids {
		_ = l.svcCtx.PhotoDao.Delete(id)
	}
	return nil
}

func (l *PhotoLogic) UpdateDeleteStatus(ctx context.Context, req *types.UpdatePhotoDeleteReq) error {
	for _, id := range req.IDs {
		l.svcCtx.DB.Model(&model.Photo{}).Where("id = ?", id).Update("status", 0)
	}
	return nil
}
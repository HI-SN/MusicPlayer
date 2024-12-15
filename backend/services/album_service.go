package services

import (
	"backend/database"
	"backend/models"
)

// AlbumService 定义专辑相关的服务函数
type AlbumService struct{}

// CreateAlbum 在数据库中创建新专辑
func (a *AlbumService) CreateAlbum(album *models.Album) error {
	query := "INSERT INTO album_info (name, description, release_date, cover_url) VALUES ($1, $2, $3, $4) RETURNING album_id"
	err := database.DB.QueryRow(query, album.Name, album.Description, album.Release_date, album.Cover_url).Scan(&album.Album_id)
	return err
}

// GetAlbumByID 根据专辑ID获取专辑信息
func (a *AlbumService) GetAlbumByID(albumID int) (*models.Album, error) {
	album := &models.Album{}
	query := "SELECT id, name, description, release_date, cover_url FROM album_info WHERE id=?"
	err := database.DB.QueryRow(query, albumID).Scan(&album.Album_id, &album.Name, &album.Description, &album.Release_date, &album.Cover_url)
	if err != nil {
		return nil, err
	}
	return album, nil
}

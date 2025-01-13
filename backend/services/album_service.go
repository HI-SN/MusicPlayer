package services

import (
	"backend/database"
	"backend/models"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// AlbumService 定义专辑相关的服务函数
type AlbumService struct{}

// CreateAlbum 在数据库中创建新专辑
func (a *AlbumService) CreateAlbum(album *models.Album) error {
	if album.Name == "" || album.Release_date.IsZero() {
		return fmt.Errorf("name and release date are required")
	}

	query := "INSERT INTO album_info (name, description, release_date, cover_url) VALUES (?, ?, ?, ?)"
	result, err := database.DB.Exec(query, album.Name, album.Description, album.Release_date, album.Cover_url)
	if err != nil {
		return fmt.Errorf("failed to create album: %v", err)
	}

	// 获取插入的专辑 ID
	albumID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %v", err)
	}

	album.Album_id = int(albumID)
	return nil
}

// UploadAlbumCover 上传专辑封面
func (a *AlbumService) UploadAlbumCover(albumID int, file *multipart.FileHeader) (string, error) {
	// 创建上传目录
	uploadDir := "uploads/album_cover"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	// 生成文件名
	fileName := fmt.Sprintf("album_%d%s", albumID, filepath.Ext(file.Filename))
	filePath := filepath.Join(uploadDir, fileName)

	// 保存文件
	if err := saveUploadedFile(file, filePath); err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	// 返回文件路径
	return filePath, nil
}

// saveUploadedFile 保存上传的文件
func saveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// UpdateAlbum 更新专辑信息
func (a *AlbumService) UpdateAlbum(albumID int, name, description string, releaseDate time.Time, coverFile *multipart.FileHeader) error {
	// 获取当前专辑信息
	var currentCoverURL string
	query := "SELECT cover_url FROM album_info WHERE id=?"
	err := database.DB.QueryRow(query, albumID).Scan(&currentCoverURL)
	if err != nil {
		return fmt.Errorf("failed to get current album info: %v", err)
	}

	// 如果上传了新封面，保存并更新封面 URL
	if coverFile != nil {
		coverURL, err := a.UploadAlbumCover(albumID, coverFile)
		if err != nil {
			return fmt.Errorf("failed to upload cover: %v", err)
		}
		currentCoverURL = coverURL
	}

	// 更新专辑信息
	query = "UPDATE album_info SET name=?, description=?, release_date=?, cover_url=? WHERE id=?"
	_, err = database.DB.Exec(query, name, description, releaseDate, currentCoverURL, albumID)
	return err
}

// GetAlbumByID 根据专辑ID获取专辑信息，包括歌手信息和歌曲信息
func (a *AlbumService) GetAlbumByID(albumID int) (*models.Album, []models.Artist, []models.Song, error) {
	// 获取专辑基本信息
	album := &models.Album{}
	query := "SELECT id, name, description, release_date, cover_url FROM album_info WHERE id=?"
	err := database.DB.QueryRow(query, albumID).Scan(&album.Album_id, &album.Name, &album.Description, &album.Release_date, &album.Cover_url)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get album: %v", err)
	}

	// 获取专辑的歌手信息
	artists, err := a.getArtistsByAlbumID(albumID)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get artists: %v", err)
	}

	// 获取专辑的歌曲信息
	songs, err := a.getSongsByAlbumID(albumID)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get songs: %v", err)
	}

	return album, artists, songs, nil
}

// getArtistsByAlbumID 根据专辑ID获取关联的歌手信息
func (a *AlbumService) getArtistsByAlbumID(albumID int) ([]models.Artist, error) {
	var artists []models.Artist
	query := `
		SELECT ai.id, ai.name, ai.bio, ai.profile_pic, ai.type, ai.nation
		FROM artist_info ai
		JOIN artist_album_relation aar ON ai.id = aar.artist_id
		WHERE aar.album_id = ?
	`
	rows, err := database.DB.Query(query, albumID)
	if err != nil {
		return nil, fmt.Errorf("failed to query artists: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var artist models.Artist
		err := rows.Scan(&artist.Artist_id, &artist.Name, &artist.Bio, &artist.Profile_pic, &artist.Type, &artist.Nation)
		if err != nil {
			return nil, fmt.Errorf("failed to scan artist: %v", err)
		}
		artists = append(artists, artist)
	}

	return artists, nil
}

// getSongsByAlbumID 根据专辑ID获取关联的歌曲信息
func (a *AlbumService) getSongsByAlbumID(albumID int) ([]models.Song, error) {
	var songs []models.Song
	query := `
		SELECT id, title, duration, album_id, genre, release_date, song_url, lyrics, created_at, updated_at, song_hit
		FROM song_info
		WHERE album_id = ?
	`
	rows, err := database.DB.Query(query, albumID)
	if err != nil {
		return nil, fmt.Errorf("failed to query songs: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var song models.Song
		err := rows.Scan(&song.Song_id, &song.Title, &song.Duration, &song.Album_id, &song.Genre, &song.Release_date, &song.Song_url, &song.Lyrics, &song.Created_at, &song.Updated_at, &song.Song_hit)
		if err != nil {
			return nil, fmt.Errorf("failed to scan song: %v", err)
		}
		songs = append(songs, song)
	}

	return songs, nil
}

// DeleteAlbum 删除专辑
func (a *AlbumService) DeleteAlbum(albumID int) error {
	// 获取专辑封面路径
	var coverURL string
	query := "SELECT cover_url FROM album_info WHERE id=?"
	err := database.DB.QueryRow(query, albumID).Scan(&coverURL)
	if err != nil {
		return fmt.Errorf("failed to get cover URL: %v", err)
	}

	// 删除封面文件
	if coverURL != "" {
		if err := os.Remove(coverURL); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to delete cover file: %v", err)
		}
	}

	// 删除专辑记录
	_, err = database.DB.Exec("DELETE FROM album_info WHERE id=?", albumID)
	return err
}

package services

import (
	"backend/database"
	"backend/models"
)

// SongService 定义歌曲相关的服务函数
type SongService struct{}

// CreateSong 在数据库中创建新歌曲
func (s *SongService) CreateSong(song *models.Song) error {
	query := "INSERT INTO song_info (title, duration, album_id, genre, release_date, song_url, lyrics, created_at, updated_at, song_hit) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	err := database.DB.QueryRow(query, song.Title, song.Duration, song.Album_id, song.Genre, song.Release_date, song.Song_url, song.Lyrics, song.Created_at, song.Updated_at, song.Song_hit).Scan(&song.Song_id)
	return err
}

// GetSongByID 根据歌曲ID获取歌曲信息
func (s *SongService) GetSongByID(songID int) (*models.Song, error) {
	song := &models.Song{}
	query := "SELECT id, title, duration, album_id, genre, release_date, song_url, lyrics, created_at, updated_at, song_hit FROM song_info WHERE id=?"
	err := database.DB.QueryRow(query, songID).Scan(&song.Song_id, &song.Title, &song.Duration, &song.Album_id, &song.Genre, &song.Release_date, &song.Song_url, &song.Lyrics, &song.Created_at, &song.Updated_at, &song.Song_hit)
	if err != nil {
		return nil, err
	}
	return song, nil
}

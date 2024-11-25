package services

import (
	"backend/database"
	"backend/models"
	"time"
)

// PlayerService 定义播放器相关的服务函数
type PlayerService struct{}

// PlaySong 播放歌曲
func (p *PlayerService) PlaySong(songID int) (string, error) {
	song, err := (&SongService{}).GetSongByID(songID)
	if err != nil {
		return "", err
	}
	return song.Song_url, nil
}

// PauseSong 暂停歌曲
func (p *PlayerService) PauseSong(songID int) error {
	// 这里添加暂停歌曲的逻辑
	// 例如，可以记录当前播放位置或状态
	// 假设我们有一个播放状态表来记录播放状态
	query := "INSERT INTO play_status (song_id, paused_at) VALUES ($1, $2)"
	_, err := database.DB.Exec(query, songID, time.Now())
	return err
}

// ResumeSong 继续播放歌曲
func (p *PlayerService) ResumeSong(songID int) error {
	// 这里添加继续播放歌曲的逻辑
	// 例如，可以从上次暂停的位置继续播放
	// 假设我们有一个播放状态表来记录播放状态
	query := "DELETE FROM play_status WHERE song_id = $1"
	_, err := database.DB.Exec(query, songID)
	return err
}

// AdjustVolume 调整音量
func (p *PlayerService) AdjustVolume(songID int, volume int) error {
	// 这里添加调整音量的逻辑
	// 例如，可以记录当前音量设置
	// 假设我们有一个播放状态表来记录播放状态
	query := "INSERT INTO play_status (song_id, volume) VALUES ($1, $2)"
	_, err := database.DB.Exec(query, songID, volume)
	return err
}

// CreatePlaylist 创建播放列表
func (p *PlayerService) CreatePlaylist(title, description string) (int, error) {
	playlist := &models.Playlist{
		Title:       title,
		Description: description,
		Create_at:   time.Now(),
	}
	return playlist.Playlist_id, (&PlaylistService{}).CreatePlaylist(playlist)
}

// AddSongToPlaylist 添加歌曲到播放列表
func (p *PlayerService) AddSongToPlaylist(playlistID, songID int) error {
	relation := &models.SongPlaylistRelation{
		PlaylistID: playlistID,
		SongID:     songID,
	}
	return (&SongPlaylistRelationService{}).CreateSongPlaylistRelation(relation)
}

// RemoveSongFromPlaylist 从播放列表移除歌曲
func (p *PlayerService) RemoveSongFromPlaylist(playlistID, songID int) error {
	return (&SongPlaylistRelationService{}).DeleteSongPlaylistRelation(playlistID, songID)
}

// GetSongsByPlaylistID 获取播放列表中的所有歌曲
func (p *PlayerService) GetSongsByPlaylistID(playlistID int) ([]int, error) {
	return (&SongPlaylistRelationService{}).GetSongsByPlaylistID(playlistID)
}

// ShowLyrics 显示歌词
func (p *PlayerService) ShowLyrics(songID int) (string, error) {
	song, err := (&SongService{}).GetSongByID(songID)
	if err != nil {
		return "", err
	}
	return song.Lyrics, nil
}

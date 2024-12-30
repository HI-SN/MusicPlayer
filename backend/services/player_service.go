package services

import (
	"backend/database"
	"database/sql"
	"fmt"
)

// PlayerService 定义播放器相关的服务函数
type PlayerService struct{}

// PlaySong 播放歌曲
func (p *PlayerService) PlaySong(songID int) (string, error) {
	song, _, err := (&SongService{}).GetSongByID(songID)
	if err != nil {
		return "", err
	}
	return song.Song_url, nil
}

// PauseSong 暂停歌曲
func (p *PlayerService) PauseSong(songID int) error {
	// 不记录暂停状态，直接返回成功
	return nil
}

// ResumeSong 继续播放歌曲
func (p *PlayerService) ResumeSong(songID int) error {
	// 不记录恢复状态，直接返回成功
	return nil
}

// AdjustVolume 调整音量
func (p *PlayerService) AdjustVolume(songID int, volume int) error {
	// 不记录音量状态，直接返回成功
	return nil
}

// ShowLyrics 返回歌词文件路径
func (p *PlayerService) ShowLyrics(songID int) (string, error) {
	db := database.DB

	var lyricsPath string

	// 查询歌词文件路径
	query := `
        SELECT lyrics
        FROM song_info
        WHERE id = ?
    `
	err := db.QueryRow(query, songID).Scan(&lyricsPath)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("song not found")
		}
		return "", fmt.Errorf("failed to get lyrics path: %v", err)
	}

	// 如果歌词路径为空，返回错误
	if lyricsPath == "" {
		return "", fmt.Errorf("lyrics not found for song ID %d", songID)
	}

	return lyricsPath, nil
}

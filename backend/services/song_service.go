package services

import (
	"backend/database"
	"backend/models"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"
)

type SongService struct{}

// CreateSong creates a new song and returns the song ID
func (s *SongService) CreateSong(title string, artistID int, duration int, albumID int, genre string, releaseDate time.Time, songURL string, lyrics string) (int64, error) {
	db := database.DB

	// 插入歌曲到数据库
	query := `
		INSERT INTO song_info (title, duration, album_id, genre, release_date, song_url, lyrics, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := db.Exec(query, title, duration, albumID, genre, releaseDate, songURL, lyrics, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	// 获取插入的歌曲 ID
	songID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// 创建艺术家与歌曲的关系
	relationQuery := `
		INSERT INTO artist_song_relation (artist_id, song_id)
		VALUES (?, ?)
	`
	_, err = db.Exec(relationQuery, artistID, songID)
	if err != nil {
		return 0, err
	}

	return songID, nil
}

// GetSongByID retrieves a song by its ID along with artist name
func (s *SongService) GetSongByID(songID int) (*models.Song, string, error) {
	db := database.DB

	var song models.Song
	var artistName string

	// 查询歌曲信息
	query := `
		SELECT id, title, duration, album_id, genre, release_date, song_url, lyrics, created_at, updated_at, song_hit
		FROM song_info
		WHERE id = ?
	`
	err := db.QueryRow(query, songID).Scan(
		&song.Song_id, &song.Title, &song.Duration, &song.Album_id, &song.Genre, &song.Release_date,
		&song.Song_url, &song.Lyrics, &song.Created_at, &song.Updated_at, &song.Song_hit,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, "", errors.New("song not found")
		}
		return nil, "", err
	}

	// 查询艺术家与歌曲的关系
	var artistID int
	relationQuery := `
		SELECT artist_id
		FROM artist_song_relation
		WHERE song_id = ?
	`
	err = db.QueryRow(relationQuery, songID).Scan(&artistID)
	if err != nil {
		return nil, "", err
	}

	// 查询艺术家信息
	artistQuery := `
		SELECT name
		FROM artist_info
		WHERE id = ?
	`
	err = db.QueryRow(artistQuery, artistID).Scan(&artistName)
	if err != nil {
		return nil, "", err
	}

	return &song, artistName, nil
}

func (s *SongService) UpdateSongInfo(songID int, title string, duration int, albumID int, genre string, releaseDate time.Time, songURL string, lyrics string) error {
	db := database.DB

	// 更新歌曲信息
	query := `
		UPDATE song_info
		SET title = ?, duration = ?, album_id = ?, genre = ?, release_date = ?, song_url = ?, lyrics = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := db.Exec(query, title, duration, albumID, genre, releaseDate, songURL, lyrics, time.Now(), songID)
	if err != nil {
		return err
	}

	return nil
}

// UploadLyricsBySongID uploads lyrics for a song
func (s *SongService) UploadLyricsBySongID(songID int, lyricsPath string) error {
	db := database.DB

	// 更新歌词文件路径
	query := `
		UPDATE song_info
		SET lyrics = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := db.Exec(query, lyricsPath, time.Now(), songID)
	if err != nil {
		return err
	}

	return nil
}

func (s *SongService) DownloadAudioBySongID(songID int) (string, error) {
	db := database.DB

	var songURL string

	// 查询歌曲的音频 URL
	query := `
        SELECT song_url
        FROM song_info
        WHERE id = ?
    `
	err := db.QueryRow(query, songID).Scan(&songURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("audio file not found in database")
		}
		return "", err
	}

	// 检查文件是否存在
	if _, err := os.Stat(songURL); os.IsNotExist(err) {
		return "", fmt.Errorf("audio file does not exist at path: %s", songURL)
	}

	return songURL, nil
}

func (s *SongService) DownloadLyricsBySongID(songID int) (string, error) {
	db := database.DB

	var lyrics string

	// 查询歌曲的歌词
	query := `
        SELECT lyrics
        FROM song_info
        WHERE id = ?
    `
	err := db.QueryRow(query, songID).Scan(&lyrics)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("lyrics not found in database")
		}
		return "", err
	}

	// 检查文件是否存在
	if _, err := os.Stat(lyrics); os.IsNotExist(err) {
		return "", fmt.Errorf("lyrics file does not exist at path: %s", lyrics)
	}

	// 读取歌词文件内容
	content, err := os.ReadFile(lyrics)
	if err != nil {
		return "", fmt.Errorf("failed to read lyrics file: %v", err)
	}

	return string(content), nil
}

// DeleteSongByID 删除歌曲及其关联信息
func (s *SongService) DeleteSongByID(songID int) error {
	db := database.DB

	// 获取歌曲的音频和歌词文件路径
	var songURL, lyrics string
	query := `
        SELECT song_url, lyrics
        FROM song_info
        WHERE id = ?
    `
	err := db.QueryRow(query, songID).Scan(&songURL, &lyrics)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("song not found")
		}
		return err
	}

	// 删除音频文件
	if songURL != "" {
		if err := os.Remove(songURL); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to delete audio file: %v", err)
		}
	}

	// 删除歌词文件
	if lyrics != "" {
		if err := os.Remove(lyrics); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to delete lyrics file: %v", err)
		}
	}

	// 删除歌手与歌曲的关联
	_, err = db.Exec("DELETE FROM artist_song_relation WHERE song_id = ?", songID)
	if err != nil {
		return fmt.Errorf("failed to delete artist_song_relation: %v", err)
	}

	// 删除歌曲与专辑/歌单的关联
	_, err = db.Exec("DELETE FROM song_playlist_relation WHERE song_id = ?", songID)
	if err != nil {
		return fmt.Errorf("failed to delete song_playlist_relation: %v", err)
	}

	// 删除歌曲信息
	_, err = db.Exec("DELETE FROM song_info WHERE id = ?", songID)
	if err != nil {
		return fmt.Errorf("failed to delete song_info: %v", err)
	}

	return nil
}

// GetCommentsBySongID 获取歌曲相关评论
func (s *SongService) GetCommentsBySongID(songID int) ([]models.Comment, error) {
	db := database.DB

	// 查询歌曲相关评论
	query := `
        SELECT c.id, c.content, c.user_id, u.user_name, c.created_at, c.type, c.target_id
        FROM comment_info c
        JOIN user_info u ON c.user_id = u.user_id
        WHERE c.target_id = ? AND c.type = 'song'
        ORDER BY c.created_at DESC
    `
	rows, err := db.Query(query, songID)
	if err != nil {
		return nil, fmt.Errorf("failed to query comments: %v", err)
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.Comment_id,
			&comment.Content,
			&comment.User_id,
			&comment.User_name,
			&comment.Created_at,
			&comment.Type,
			&comment.Target_id,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comment: %v", err)
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating comments: %v", err)
	}

	return comments, nil
}

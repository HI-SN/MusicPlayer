package services

import (
	"backend/database"
	"backend/models"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// PlaylistService 定义播放列表相关的服务函数
type PlaylistService struct{}

// CheckPlaylistExists 检查播放列表是否存在
func (p *PlaylistService) CheckPlaylistExists(playlistID int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM playlist_info WHERE id = ?)"
	err := database.DB.QueryRow(query, playlistID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// CheckUserExists 检查用户是否存在
func (p *PlaylistService) CheckUserExists(userID string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM user_info WHERE user_id = ?)"
	err := database.DB.QueryRow(query, userID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// CreatePlaylist 在数据库中创建新歌单
func (p *PlaylistService) CreatePlaylist(playlist *models.Playlist) error {
	// 检查用户是否存在
	exists, err := p.CheckUserExists(playlist.User_id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user with ID %s does not exist", playlist.User_id)
	}

	// 如果 created_at 为空，则设置为当前时间
	if playlist.Create_at.IsZero() {
		playlist.Create_at = time.Now()
	}

	query := "INSERT INTO playlist_info (title, user_id, created_at, description, type, hits, cover_url) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := database.DB.Exec(query, playlist.Title, playlist.User_id, playlist.Create_at, playlist.Description, playlist.Type, playlist.Hits, playlist.Cover_url)
	if err != nil {
		return err
	}

	// 获取插入后的自增主键
	playlistID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// 将自增主键赋值给 playlist.Playlist_id
	playlist.Playlist_id = int(playlistID)

	return nil
}

// 删除歌单
func (p *PlaylistService) DeletePlaylistByID(playlistID int, user_id string) error {
	// 检查播放列表是否存在
	exists, err := p.CheckPlaylistExists(playlistID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("playlist with ID %d does not exist", playlistID)
	}

	// 检查用户是否存在
	exists, err = p.CheckUserExists(user_id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user with ID %s does not exist", user_id)
	}

	query := "DELETE FROM playlist_info WHERE id=? and user_id=?"
	_, err = database.DB.Exec(query, playlistID, user_id)
	return err
}

// GetPlaylistByID 根据播放列表ID获取播放列表信息，包括歌曲信息和用户是否like的信息
func (p *PlaylistService) GetPlaylistByID(playlistID int) (*models.Playlist, []models.Song, bool, error) {
	// 检查播放列表是否存在
	exists, err := p.CheckPlaylistExists(playlistID)
	if err != nil {
		return nil, nil, false, err
	}
	if !exists {
		return nil, nil, false, fmt.Errorf("playlist with ID %d does not exist", playlistID)
	}

	// 获取歌单基本信息
	playlist := &models.Playlist{}
	query := "SELECT id, title, user_id, created_at, description, type, hits, cover_url FROM playlist_info WHERE id=?"
	err = database.DB.QueryRow(query, playlistID).Scan(&playlist.Playlist_id, &playlist.Title, &playlist.User_id, &playlist.Create_at, &playlist.Description, &playlist.Type, &playlist.Hits, &playlist.Cover_url)
	if err != nil {
		return nil, nil, false, fmt.Errorf("failed to get playlist: %v", err)
	}

	// 获取歌单中的歌曲信息
	songs, err := p.getSongsByPlaylistID(playlistID)
	if err != nil {
		return nil, nil, false, fmt.Errorf("failed to get songs: %v", err)
	}

	// 获取用户是否like歌单的信息
	var isLiked bool
	query = "SELECT EXISTS(SELECT 1 FROM user_like_playlist WHERE playlist_id = ?)"
	err = database.DB.QueryRow(query, playlistID).Scan(&isLiked)
	if err != nil {
		return nil, nil, false, fmt.Errorf("failed to check if playlist is liked: %v", err)
	}

	return playlist, songs, isLiked, nil
}

// getSongsByPlaylistID 获取歌单中的歌曲信息
func (p *PlaylistService) getSongsByPlaylistID(playlistID int) ([]models.Song, error) {
	var songs []models.Song

	// 查询歌单中的歌曲
	query := `
		SELECT s.id, s.title, s.duration, s.album_id, s.genre, s.release_date, s.song_url, s.lyrics, s.created_at, s.updated_at, s.song_hit
		FROM song_info s
		JOIN song_playlist_relation spr ON s.id = spr.song_id
		WHERE spr.playlist_id = ?
	`
	rows, err := database.DB.Query(query, playlistID)
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

// UpdatePlaylist 更新播放列表信息
func (p *PlaylistService) UpdatePlaylist(playlist *models.Playlist) error {
	// 检查播放列表是否存在
	exists, err := p.CheckPlaylistExists(playlist.Playlist_id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("playlist with ID %d does not exist", playlist.Playlist_id)
	}

	// 检查用户是否存在
	exists, err = p.CheckUserExists(playlist.User_id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user with ID %s does not exist", playlist.User_id)
	}

	query := "UPDATE playlist_info SET title=?, user_id=?, description=?, type=?, hits=?, cover_url=? WHERE id=?"
	_, err = database.DB.Exec(query, playlist.Title, playlist.User_id, playlist.Description, playlist.Type, playlist.Hits, playlist.Cover_url, playlist.Playlist_id)
	return err
}

// DeletePlaylist 删除播放列表
func (p *PlaylistService) DeletePlaylist(playlistID int) error {
	query := "DELETE FROM playlist_info WHERE id=?"
	_, err := database.DB.Exec(query, playlistID)
	return err
}

// AddSongToPlaylist 添加歌曲到播放列表
func (p *PlaylistService) AddSongToPlaylist(playlistID, songID int) error {
	// 检查播放列表是否存在
	exists, err := p.CheckPlaylistExists(playlistID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("playlist with ID %d does not exist", playlistID)
	}

	// 检查歌曲是否存在
	var songExists bool
	query := "SELECT EXISTS(SELECT 1 FROM song_info WHERE id = ?)"
	err = database.DB.QueryRow(query, songID).Scan(&songExists)
	if err != nil {
		return err
	}
	if !songExists {
		return fmt.Errorf("song with ID %d does not exist", songID)
	}

	query = "INSERT INTO song_playlist_relation (playlist_id, song_id) VALUES (?, ?)"
	_, err = database.DB.Exec(query, playlistID, songID)
	return err
}

// RemoveSongFromPlaylist 从播放列表移除歌曲
func (p *PlaylistService) RemoveSongFromPlaylist(playlistID, songID int) error {
	// 检查播放列表是否存在
	exists, err := p.CheckPlaylistExists(playlistID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("playlist with ID %d does not exist", playlistID)
	}

	// 检查歌曲是否存在
	var songExists bool
	query := "SELECT EXISTS(SELECT 1 FROM song_info WHERE id = ?)"
	err = database.DB.QueryRow(query, songID).Scan(&songExists)
	if err != nil {
		return err
	}
	if !songExists {
		return fmt.Errorf("song with ID %d does not exist", songID)
	}

	query = "DELETE FROM song_playlist_relation WHERE playlist_id=? AND song_id=?"
	_, err = database.DB.Exec(query, playlistID, songID)
	return err
}

// GetSongsByPlaylistID 获取播放列表中的所有歌曲
func (p *PlaylistService) GetSongsByPlaylistID(playlistID int, userID string, isLoggedIn bool) ([]gin.H, error) {
	// 检查播放列表是否存在
	exists, err := p.CheckPlaylistExists(playlistID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("playlist with ID %d does not exist", playlistID)
	}

	// 查询播放列表中的歌曲 ID
	query := `
		SELECT song_id
		FROM song_playlist_relation
		WHERE playlist_id = ?
	`
	rows, err := database.DB.Query(query, playlistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []gin.H

	for rows.Next() {
		var songID int
		if err := rows.Scan(&songID); err != nil {
			return nil, err
		}

		// 获取歌曲详细信息（复现 GetSongsBySearch 逻辑）
		var song struct {
			ID          int
			Title       string
			Duration    int
			AlbumID     int
			Genre       string
			ReleaseDate string
			SongURL     string
			Lyrics      string
			CreatedAt   time.Time
			UpdatedAt   time.Time
			SongHit     int
		}
		songQuery := `
			SELECT id, title, duration, album_id, genre, release_date, song_url, lyrics, created_at, updated_at, song_hit
			FROM song_info
			WHERE id = ?
		`
		err := database.DB.QueryRow(songQuery, songID).Scan(
			&song.ID, &song.Title, &song.Duration, &song.AlbumID, &song.Genre, &song.ReleaseDate,
			&song.SongURL, &song.Lyrics, &song.CreatedAt, &song.UpdatedAt, &song.SongHit,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				continue // 如果没有找到歌曲，跳过
			}
			return nil, err
		}

		// 获取歌手名称（复现 GetArtistNameBySongID 逻辑）
		var artistName string
		artistQuery := `
			SELECT ai.name
			FROM artist_info ai
			JOIN artist_song_relation asr ON ai.id = asr.artist_id
			WHERE asr.song_id = ?
		`
		err = database.DB.QueryRow(artistQuery, songID).Scan(&artistName)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		// 获取专辑名称（复现 GetAlbumNameByID 逻辑）
		var albumName string
		albumQuery := `
			SELECT name
			FROM album_info
			WHERE id = ?
		`
		err = database.DB.QueryRow(albumQuery, song.AlbumID).Scan(&albumName)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		// 检查用户是否喜欢该歌曲（复现 IsSongLikedByUser 逻辑）
		var isLiked bool
		if isLoggedIn {
			var count int
			likeQuery := `
				SELECT COUNT(*)
				FROM user_like_song
				WHERE user_id = ? AND song_id = ?
			`
			err := database.DB.QueryRow(likeQuery, userID, songID).Scan(&count)
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				return nil, err
			}
			isLiked = count > 0
		} else {
			isLiked = false
		}

		// 格式化时长
		minutes := song.Duration / 60
		seconds := song.Duration % 60
		formattedDuration := fmt.Sprintf("%02d:%02d", minutes, seconds)

		// 构造歌曲信息
		songInfo := gin.H{
			"id":       strconv.Itoa(song.ID),
			"title":    song.Title,
			"singer":   artistName,
			"album":    albumName,
			"album_id": strconv.Itoa(song.AlbumID),
			"duration": formattedDuration,
			"liked":    strconv.FormatBool(isLiked), // 动态设置 liked 字段
		}

		songs = append(songs, songInfo)
	}

	return songs, nil
}

// GetPlaylistByUserID 根据用户ID获取用户创建的歌单列表
func (p *PlaylistService) GetPlaylistByUserID(userID string) ([]models.Playlist, error) {
	var playlists []models.Playlist
	query := "SELECT id, title, user_id, created_at, description, type, hits, cover_url FROM playlist_info WHERE user_id=?"
	rows, err := database.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var playlist models.Playlist
		err := rows.Scan(&playlist.Playlist_id, &playlist.Title, &playlist.User_id, &playlist.Create_at, &playlist.Description, &playlist.Type, &playlist.Hits, &playlist.Cover_url)
		if err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}
	return playlists, nil
}

// GetPlaylistByPlaylistID 根据歌单ID获取歌单列表
func (p *PlaylistService) GetPlaylistByPlaylistID(PlaylistID int) (*models.Playlist, error) {
	playlist := &models.Playlist{}
	query := "SELECT id, title, user_id, created_at, description, type, hits, cover_url FROM playlist_info WHERE id=?"
	err := database.DB.QueryRow(query, PlaylistID).Scan(&playlist.Playlist_id, &playlist.Title, &playlist.User_id, &playlist.Create_at, &playlist.Description, &playlist.Type, &playlist.Hits, &playlist.Cover_url)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

// UpdatePlaylistCover 更新歌单封面 URL
func (p *PlaylistService) UpdatePlaylistCover(playlistID int, coverURL string) error {
	// 检查播放列表是否存在
	exists, err := p.CheckPlaylistExists(playlistID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("playlist with ID %d does not exist", playlistID)
	}

	// 更新 cover_url
	query := "UPDATE playlist_info SET cover_url=? WHERE id=?"
	_, err = database.DB.Exec(query, coverURL, playlistID)
	return err
}

// SongIDResponse 定义返回的歌曲ID结构体
type SongIDResponse struct {
	SongID int `json:"song_id"`
}

// GetSongIDsByPlaylistID 根据歌单ID获取该歌单下的所有歌曲ID
func (p *PlaylistService) GetSongIDsByPlaylistID(playlistID int) ([]SongIDResponse, error) {
	// 检查歌单是否存在
	exists, err := p.CheckPlaylistExists(playlistID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("playlist with ID %d does not exist", playlistID)
	}

	// 查询歌单下的歌曲ID
	query := `
		SELECT song_id
		FROM song_playlist_relation
		WHERE playlist_id = ?
	`
	rows, err := database.DB.Query(query, playlistID)
	if err != nil {
		return nil, fmt.Errorf("failed to query song IDs: %v", err)
	}
	defer rows.Close()

	var songIDs []SongIDResponse

	for rows.Next() {
		var songID int
		if err := rows.Scan(&songID); err != nil {
			return nil, fmt.Errorf("failed to scan song ID: %v", err)
		}
		songIDs = append(songIDs, SongIDResponse{SongID: songID})
	}

	return songIDs, nil
}

// GetPlaylistsByType 根据歌单类型获取歌单列表
func (p *PlaylistService) GetPlaylistsByType(playlistType string) ([]gin.H, error) {
	var query string
	var args []interface{}

	// 如果 type 为 "推荐"，查询所有歌单
	if playlistType == "推荐" {
		query = `
			SELECT id, title
			FROM playlist_info
		`
	} else {
		// 否则，查询符合指定类型的歌单
		query = `
			SELECT id, title
			FROM playlist_info
			WHERE type = ?
		`
		args = append(args, playlistType)
	}

	// 执行查询
	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var playlists []gin.H

	for rows.Next() {
		var id int
		var title string

		if err := rows.Scan(&id, &title); err != nil {
			return nil, err
		}

		// 构造歌单信息
		playlist := gin.H{
			"id":    strconv.Itoa(id),
			"title": title,
		}

		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

// GetPlaylistsBySearch 根据搜索关键词获取歌单信息
func (p *PlaylistService) GetPlaylistsBySearch(keyword string) ([]models.Playlist, error) {
	query := `SELECT * FROM playlist_info WHERE title LIKE ?`
	rows, err := database.DB.Query(query, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var playlists []models.Playlist
	for rows.Next() {
		var playlist models.Playlist
		if err := rows.Scan(&playlist.Playlist_id, &playlist.Title, &playlist.User_id, &playlist.Create_at, &playlist.Description, &playlist.Type, &playlist.Hits, &playlist.Cover_url); err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

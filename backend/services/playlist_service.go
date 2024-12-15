package services

import (
	"backend/database"
	"backend/models"
)

// PlaylistService 定义播放列表相关的服务函数
type PlaylistService struct{}

// CreatePlaylist 在数据库中创建新播放列表
func (p *PlaylistService) CreatePlaylist(playlist *models.Playlist) error {
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

// GetPlaylistByID 根据播放列表ID获取播放列表信息
func (p *PlaylistService) GetPlaylistByID(playlistID int) (*models.Playlist, error) {
	playlist := &models.Playlist{}
	query := "SELECT playlist_id, title, user_id, create_at, description, type, hits FROM playlist_info WHERE playlist_id=$1"
	err := database.DB.QueryRow(query, playlistID).Scan(&playlist.Playlist_id, &playlist.Title, &playlist.User_id, &playlist.Create_at, &playlist.Description, &playlist.Type, &playlist.Hits)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

// UpdatePlaylist 更新播放列表信息
func (p *PlaylistService) UpdatePlaylist(playlist *models.Playlist) error {
	query := "UPDATE playlist_info SET title=$1, user_id=$2, description=$3, type=$4, hits=$5 WHERE playlist_id=$6"
	_, err := database.DB.Exec(query, playlist.Title, playlist.User_id, playlist.Description, playlist.Type, playlist.Hits, playlist.Playlist_id)
	return err
}

// DeletePlaylist 删除播放列表
func (p *PlaylistService) DeletePlaylist(playlistID int) error {
	query := "DELETE FROM playlist_info WHERE playlist_id=$1"
	_, err := database.DB.Exec(query, playlistID)
	return err
}

// AddSongToPlaylist 添加歌曲到播放列表
func (p *PlaylistService) AddSongToPlaylist(playlistID, songID int) error {
	relation := &models.SongPlaylistRelation{
		PlaylistID: playlistID,
		SongID:     songID,
	}
	return (&SongPlaylistRelationService{}).CreateSongPlaylistRelation(relation)
}

// RemoveSongFromPlaylist 从播放列表移除歌曲
func (p *PlaylistService) RemoveSongFromPlaylist(playlistID, songID int) error {
	return (&SongPlaylistRelationService{}).DeleteSongPlaylistRelation(playlistID, songID)
}

// GetSongsByPlaylistID 获取播放列表中的所有歌曲
func (p *PlaylistService) GetSongsByPlaylistID(playlistID int) ([]int, error) {
	return (&SongPlaylistRelationService{}).GetSongsByPlaylistID(playlistID)
}

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

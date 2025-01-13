package services

import (
	"backend/database"
	"backend/models"
)

// SongPlaylistRelationService 定义歌曲和播放列表关系相关的服务函数
type SongPlaylistRelationService struct{}

// CreateSongPlaylistRelation 在数据库中创建新的歌曲和播放列表关系
func (s *SongPlaylistRelationService) CreateSongPlaylistRelation(relation *models.SongPlaylistRelation) error {
	query := "INSERT INTO song_playlist_relation (playlist_id, song_id, added_time) VALUES ($1, $2, $3)"
	_, err := database.DB.Exec(query, relation.PlaylistID, relation.SongID)
	return err
}

// DeleteSongPlaylistRelation 根据播放列表ID和歌曲ID删除关系
func (s *SongPlaylistRelationService) DeleteSongPlaylistRelation(playlistID, songID int) error {
	query := "DELETE FROM song_playlist_relation WHERE playlist_id=$1 AND song_id=$2"
	_, err := database.DB.Exec(query, playlistID, songID)
	return err
}

// GetSongsByPlaylistID 根据播放列表ID获取所有歌曲ID
func (s *SongPlaylistRelationService) GetSongsByPlaylistID(playlistID int) ([]int, error) {
	var songIDs []int
	query := "SELECT song_id FROM song_playlist_relation WHERE playlist_id=$1"
	rows, err := database.DB.Query(query, playlistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var songID int
		if err := rows.Scan(&songID); err != nil {
			return nil, err
		}
		songIDs = append(songIDs, songID)
	}

	return songIDs, nil
}

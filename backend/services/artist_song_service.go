package services

import (
	"backend/database"
	"backend/models"
)

type ArtistSongService struct{}

// 创建歌手-音乐信息
func (ass *ArtistSongService) CreateArtistSong(as *models.ArtistSongRelation) error {
	query := `INSERT INTO artist_song_relation (artist_id, song_id) VALUES (?, ?)`
	_, err := database.DB.Exec(query, as.ArtistID, as.SongID)
	return err
}

// 通过歌手id获取歌手-音乐信息
func (ass *ArtistSongService) GetSongListByArtistID(artistID int) ([]int, error) {
	// 执行查询
	rows, err := database.DB.Query("SELECT song_id FROM artist_song_relation WHERE artist_id = ?", artistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []int
	// 遍历查询结果
	for rows.Next() {
		var song_id int
		// 使用Scan方法将列值映射到结构体的字段
		err = rows.Scan(&song_id)
		if err != nil {
			return nil, err
		}
		results = append(results, song_id)
	}

	// 检查遍历过程中是否有错误发生
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// 通过歌曲id获取歌手-音乐信息
func (ass *ArtistSongService) GetArtistListBySongID(songID int) ([]int, error) {
	// 执行查询
	rows, err := database.DB.Query("SELECT artist_id FROM artist_song_relation WHERE song_id = ?", songID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []int
	// 遍历查询结果
	for rows.Next() {
		var artistID int
		// 使用Scan方法将列值映射到结构体的字段
		err = rows.Scan(&artistID)
		if err != nil {
			return nil, err
		}
		results = append(results, artistID)
	}

	// 检查遍历过程中是否有错误发生
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

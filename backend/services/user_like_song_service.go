package services

import (
	"backend/database"
	"backend/models"
)

type UserSongService struct{}

// 创建我喜欢的音乐信息
func (uss *UserSongService) CreateUserLikeSong(uls *models.UserLikeSong) error {
	query := `INSERT INTO user_like_song (user_id, song_id) VALUES (?, ?)`
	_, err := database.DB.Exec(query, uls.UserID, uls.SongID)
	return err
}

// 删除我喜欢的音乐信息
func (uss *UserSongService) DeleteUserLikeSong(uls *models.UserLikeSong) error {
	query := `DELETE FROM user_like_song WHERE user_id = ? AND song_id = ?`
	_, err := database.DB.Exec(query, uls.UserID, uls.SongID)
	return err
}

// 根据用户id获取喜欢的歌曲列表
func (uss *UserSongService) GetUserLikeSongList(userID string) ([]int, error) {
	// 执行查询
	rows, err := database.DB.Query("SELECT song_id FROM user_like_song WHERE user_id = ?", userID)
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

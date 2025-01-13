package services

import (
	"backend/database"
	"backend/models"
)

type UserPlaylistService struct{}

// 创建我收藏的歌单信息
func (ups *UserPlaylistService) CreateUserLikePlaylist(ulp *models.UserLikePlaylist) error {
	query := `INSERT INTO user_like_playlist (user_id, playlist_id) VALUES (?, ?)`
	_, err := database.DB.Exec(query, ulp.UserID, ulp.PlaylistID)
	return err
}

// 删除收藏的歌单
func (ups *UserPlaylistService) DeleteUserLikePlaylist(ulp *models.UserLikePlaylist) error {
	query := `DELETE FROM user_like_playlist WHERE user_id=? AND playlist_id=?`
	_, err := database.DB.Exec(query, ulp.UserID, ulp.PlaylistID)
	return err
}

// 根据用户id获取我收藏的歌单列表
func (ups *UserPlaylistService) GetUserLikePlaylistList(userID string) ([]int, error) {
	// 执行查询
	rows, err := database.DB.Query("SELECT playlist_id FROM user_like_playlist WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []int
	// 遍历查询结果
	for rows.Next() {
		var playlist_id int
		// 使用Scan方法将列值映射到结构体的字段
		err = rows.Scan(&playlist_id)
		if err != nil {
			return nil, err
		}
		results = append(results, playlist_id)
	}

	// 检查遍历过程中是否有错误发生
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

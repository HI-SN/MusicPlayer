package services

import (
	"backend/database"
	"backend/models"
)

type FollowService struct{}

// 创建关注用户信息
func (fs *FollowService) CreateFollowUser(f *models.FollowUser) error {
	query := `INSERT INTO follow_user (follower_id, followed_id) VALUES (?, ?)`
	_, err := database.DB.Exec(query, f.Follower_id, f.Followed_id)
	return err
}

// 删除一条关注用户信息
func (fs *FollowService) DeleteFollowUser(f *models.FollowUser) error {
	query := `DELETE FROM follow_user WHERE follower_id = ? AND followed_id = ?`
	_, err := database.DB.Exec(query, f.Follower_id, f.Followed_id)
	return err
}

// 检查用户A是否关注用户B
func (fs *FollowService) IsAFollowUserB(userAID string, userBID string) (bool, error) {
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM follow_user WHERE follower_id = ? AND followed_id = ?", userAID, userBID).Scan(&count)
	if err != nil {
		return false, err
	}
	if count < 1 {
		return false, nil
	} else {
		return true, nil
	}
}

// 根据用户id获取关注列表（用户）
func (fs *FollowService) GetFollowingUserList(userID string) ([]string, error) {
	// 执行查询
	rows, err := database.DB.Query("SELECT followed_id FROM follow_user WHERE follower_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []string

	// 遍历查询结果
	for rows.Next() {
		var followedID string
		// 使用Scan方法将列值映射到结构体的字段
		err = rows.Scan(&followedID)
		if err != nil {
			return nil, err
		}
		results = append(results, followedID)
	}

	// 检查遍历过程中是否有错误发生
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// 根据用户id获取粉丝列表
func (fs *FollowService) GetFollowerUserList(userID string) ([]string, error) {
	// 执行查询
	rows, err := database.DB.Query("SELECT follower_id FROM follow_user WHERE followed_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []string

	// 遍历查询结果
	for rows.Next() {
		var followerID string
		// 使用Scan方法将列值映射到结构体的字段
		err = rows.Scan(&followerID)
		if err != nil {
			return nil, err
		}
		results = append(results, followerID)
	}

	// 检查遍历过程中是否有错误发生
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// 创建关注歌手信息
func (fs *FollowService) CreateFollowArtist(f *models.FollowArtist) error {
	query := `INSERT INTO follow_artist (follower_id, followed_id) VALUES (?, ?)`
	_, err := database.DB.Exec(query, f.Follower_id, f.Followed_id)
	return err
}

// 删除一条歌手关注信息
func (fs *FollowService) DeleteFollowArtist(f *models.FollowArtist) error {
	query := `DELETE FROM follow_artist WHERE follower_id = ? AND followed_id = ?`
	_, err := database.DB.Exec(query, f.Follower_id, f.Followed_id)
	return err
}

// 根据用户id获取关注列表(歌手)
func (fs *FollowService) GetFollowingArtistList(userID string) ([]int, error) {
	// 执行查询
	rows, err := database.DB.Query("SELECT followed_id FROM follow_artist WHERE follower_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []int

	// 遍历查询结果
	for rows.Next() {
		var followedID int
		// 使用Scan方法将列值映射到结构体的字段
		err = rows.Scan(&followedID)
		if err != nil {
			return nil, err
		}
		results = append(results, followedID)
	}

	// 检查遍历过程中是否有错误发生
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// 根据歌手id获取粉丝列表
func (fs *FollowService) GetFollowerArtistList(artistID int) ([]string, error) {
	// 执行查询
	rows, err := database.DB.Query("SELECT follower_id FROM follow_artist WHERE followed_id = ?", artistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []string

	// 遍历查询结果
	for rows.Next() {
		var followerID string
		// 使用Scan方法将列值映射到结构体的字段
		err = rows.Scan(&followerID)
		if err != nil {
			return nil, err
		}
		results = append(results, followerID)
	}

	// 检查遍历过程中是否有错误发生
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// 根据用户id获取粉丝个数
func (fs *FollowService) GetUserFollowerCount(userID string) (int, error) {
	// 执行查询
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM follow_user WHERE followed_id = ?", userID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 根据用户id获取关注个数
func (fs *FollowService) GetUserFollowingCount(userID string) (int, error) {
	// 执行查询
	var count1, count2 int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM follow_user WHERE follower_id = ?", userID).Scan(&count1)
	if err != nil {
		return 0, err
	}
	err = database.DB.QueryRow("SELECT COUNT(*) FROM follow_artist WHERE follower_id = ?", userID).Scan(&count2)
	if err != nil {
		return 0, err
	}

	return count1 + count2, nil
}

// 根据歌手id获取粉丝个数
func (fs *FollowService) GetArtistFollowerCount(artistID int) (int, error) {
	// 执行查询
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM follow_artist WHERE followed_id = ?", artistID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

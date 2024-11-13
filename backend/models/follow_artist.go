package models

import "database/sql"

// FollowArtist represents the follow_artist table
type FollowArtist struct {
	Follower_id string
	Followed_id int
}

func (FollowArtist) TableName() string {
	return "follow_artist"
}

// 创建关注信息
func CreateFollowArtist(db *sql.DB, f *FollowArtist) error {
	query := `INSERT INTO follow_artist (follower_id, followed_id) VALUES (?, ?)`
	_, err := db.Exec(query, f.Follower_id, f.Followed_id)
	return err
}

// 删除一条关注信息
func DeleteFollowArtist(db *sql.DB, f *FollowArtist) error {
	query := `DELETE FROM follow_artist WHERE follower_id = ? AND followed_id = ?`
	_, err := db.Exec(query, f.Follower_id, f.Followed_id)
	return err
}

// 根据用户id获取关注列表
func GetFollowingArtistList(db *sql.DB, userID string) ([]*FollowArtist, error) {
	// 执行查询
	rows, err := db.Query("SELECT * FROM follow_artist WHERE follower_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*FollowArtist

	// 遍历查询结果
	for rows.Next() {
		f := &FollowArtist{}
		// 使用Scan方法将列值映射到结构体的字段
		err = rows.Scan(&f.Follower_id, &f.Followed_id)
		if err != nil {
			return nil, err
		}
		results = append(results, f)
	}

	// 检查遍历过程中是否有错误发生
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// 根据歌手id获取粉丝列表
func GetFollowerArtistList(db *sql.DB, artistID int) ([]*FollowArtist, error) {
	// 执行查询
	rows, err := db.Query("SELECT * FROM follow_artist WHERE followed_id = ?", artistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*FollowArtist

	// 遍历查询结果
	for rows.Next() {
		f := &FollowArtist{}
		// 使用Scan方法将列值映射到结构体的字段
		err = rows.Scan(&f.Follower_id, &f.Followed_id)
		if err != nil {
			return nil, err
		}
		results = append(results, f)
	}

	// 检查遍历过程中是否有错误发生
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

package models

import "database/sql"

// FollowUser represents the follow_user table
type FollowUser struct {
	Follower_id string
	Followed_id string
}

func (FollowUser) TableName() string {
	return "follow_user"
}

// 创建关注信息
func CreateFollowUser(db *sql.DB, f *FollowUser) error {
	query := `INSERT INTO follow_user (follower_id, followed_id) VALUES (?, ?)`
	_, err := db.Exec(query, f.Follower_id, f.Followed_id)
	return err
}

// 删除一条关注信息
func DeleteFollowUser(db *sql.DB, f *FollowUser) error {
	query := `DELETE FROM follow_user WHERE follower_id = ? AND followed_id = ?`
	_, err := db.Exec(query, f.Follower_id, f.Followed_id)
	return err
}

// 根据用户id获取关注列表
func GetFollowingUserList(db *sql.DB, userID string) ([]*FollowUser, error) {
	// 执行查询
	rows, err := db.Query("SELECT * FROM follow_user WHERE follower_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*FollowUser

	// 遍历查询结果
	for rows.Next() {
		f := &FollowUser{}
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

// 根据用户id获取粉丝列表
func GetFollowerUserList(db *sql.DB, userID string) ([]*FollowUser, error) {
	// 执行查询
	rows, err := db.Query("SELECT * FROM follow_user WHERE followed_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*FollowUser

	// 遍历查询结果
	for rows.Next() {
		f := &FollowUser{}
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

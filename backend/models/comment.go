package models

import (
	"database/sql"
	"time"
)

type Comment struct {
	Comment_id int
	Content    string
	User_id    string
	Created_at time.Time
	Type       string
	Target_id  int
}

func (Comment) TableName() string {
	return "comment_info" // 数据库中的表名
}

// 在数据库中创建新评论
func CreateComment(db *sql.DB, c *Comment) error {
	c.Created_at = time.Now()

	query := `INSERT INTO comment_info (content, created_at, user_id, type, target_id) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(query, c.Content, c.Created_at, c.User_id, c.Type, c.Target_id)
	return err
}

// 在数据库中更新评论（只有评论内容可以更改）
func UpdateComment(db *sql.DB, c *Comment) error {
	query := `UPDATE comment_info SET content = ? WHERE id = ?`
	_, err := db.Exec(query, c.Content, c.Comment_id)
	return err
}

// 根据评论ID删除评论
func DeleteComment(db *sql.DB, commentID int) error {
	query := `DELETE FROM comment_info WHERE id = ?`
	_, err := db.Exec(query, commentID)
	return err
}

// 根据动态ID删除评论
func DeleteCommentByMoment(db *sql.DB, momentID int) error {
	query := `DELETE FROM comment_info WHERE target_id = ? AND type = ?`
	_, err := db.Exec(query, momentID, "moment")
	return err
}

// 根据用户ID删除评论
func DeleteCommentByUser(db *sql.DB, userID int) error {
	query := `DELETE FROM comment_info WHERE target_id = ?`
	_, err := db.Exec(query, userID)
	return err
}

// 根据歌曲ID删除评论
func DeleteCommentBySong(db *sql.DB, songID int) error {
	query := `DELETE FROM comment_info WHERE target_id = ? AND type = ?`
	_, err := db.Exec(query, songID, "song")
	return err
}

// 获取单条评论内容
func GetComment(db *sql.DB, commentID int) (*Comment, error) {
	c := &Comment{}
	query := "SELECT * FROM comment_info WHERE id=?"
	err := db.QueryRow(query, commentID).Scan(&c.Comment_id, &c.Content,
		&c.Created_at, &c.User_id, &c.Type, &c.Target_id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// 根据类型和相关ID获取评论内容
func GetAllComments(db *sql.DB, target_id int, target_type string) ([]*Comment, error) {
	// 执行查询
	rows, err := db.Query("SELECT * FROM comment_info WHERE target_id = ? AND type = ?", target_id, target_type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*Comment

	// 遍历查询结果
	for rows.Next() {
		c := &Comment{}
		// 使用Scan方法将列值映射到结构体的字段
		err = rows.Scan(&c.Comment_id, &c.Content, &c.Created_at, &c.User_id, &c.Type, &c.Target_id)
		if err != nil {
			return nil, err
		}
		results = append(results, c)
	}

	// 检查遍历过程中是否有错误发生
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

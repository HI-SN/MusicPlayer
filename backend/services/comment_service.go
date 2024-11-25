package services

import (
	"backend/database"
	"backend/models"
	"time"
)

type CommentService struct{}

// 在数据库中创建新评论
func (cs *CommentService) CreateComment(c *models.Comment) error {
	c.Created_at = time.Now()

	query := `INSERT INTO comment_info (content, created_at, user_id, type, target_id) VALUES (?, ?, ?, ?, ?)`
	_, err := database.DB.Exec(query, c.Content, c.Created_at, c.User_id, c.Type, c.Target_id)
	return err
}

// 在数据库中更新评论（只有评论内容可以更改）
func (cs *CommentService) UpdateComment(c *models.Comment) error {
	query := `UPDATE comment_info SET content = ? WHERE id = ?`
	_, err := database.DB.Exec(query, c.Content, c.Comment_id)
	return err
}

// 根据评论ID删除评论
func (cs *CommentService) DeleteComment(commentID int) error {
	query := `DELETE FROM comment_info WHERE id = ?`
	_, err := database.DB.Exec(query, commentID)
	return err
}

// 根据动态ID删除评论
func (cs *CommentService) DeleteCommentByMoment(momentID int) error {
	query := `DELETE FROM comment_info WHERE target_id = ? AND type = ?`
	_, err := database.DB.Exec(query, momentID, "moment")
	return err
}

// 根据用户ID删除评论
func (cs *CommentService) DeleteCommentByUser(userID int) error {
	query := `DELETE FROM comment_info WHERE target_id = ?`
	_, err := database.DB.Exec(query, userID)
	return err
}

// 根据歌曲ID删除评论
func (cs *CommentService) DeleteCommentBySong(songID int) error {
	query := `DELETE FROM comment_info WHERE target_id = ? AND type = ?`
	_, err := database.DB.Exec(query, songID, "song")
	return err
}

// 获取单条评论内容
func (cs *CommentService) GetComment(commentID int) (*models.Comment, error) {
	c := &models.Comment{}
	query := "SELECT * FROM comment_info WHERE id=?"
	err := database.DB.QueryRow(query, commentID).Scan(&c.Comment_id, &c.Content,
		&c.Created_at, &c.User_id, &c.Type, &c.Target_id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// 根据类型和相关ID获取评论内容
func (cs *CommentService) GetAllComments(target_id int, target_type string) ([]*models.Comment, error) {
	// 执行查询
	rows, err := database.DB.Query("SELECT * FROM comment_info WHERE target_id = ? AND type = ?", target_id, target_type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*models.Comment

	// 遍历查询结果
	for rows.Next() {
		c := &models.Comment{}
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

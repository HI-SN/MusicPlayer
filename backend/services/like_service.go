package services

import (
	"backend/database"
)

type LikeService struct {
}

// 是否点过赞-动态
func (l *LikeService) HasUserLikedMoment(momentID int, userID string) (bool, error) {
	query := `SELECT COUNT(*) FROM like_info WHERE moment_id = ? AND user_id = ?`
	var count int
	err := database.DB.QueryRow(query, momentID, userID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// 点赞动态
func (l *LikeService) CreateMomentLike(momentID int, userID string) error {
	// 创建点赞记录
	query := `INSERT INTO like_info (moment_id, user_id) VALUES (?, ?)`
	_, err := database.DB.Exec(query, momentID, userID)
	return err
}

// 取消动态点赞
func (l *LikeService) DeleteMomentLike(momentID int, userID string) error {
	// 删除点赞记录
	query := `DELETE FROM like_info WHERE moment_id = ? AND user_id = ?`
	_, err := database.DB.Exec(query, momentID, userID)
	return err
}

// 统计动态点赞数
func (l *LikeService) GetMomentLikeCount(momentID int) (int, error) {
	query := `SELECT COUNT(*) FROM like_info WHERE moment_id = ?`
	var count int
	err := database.DB.QueryRow(query, momentID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 是否点过赞-评论
func (l *LikeService) HasUserLikedComment(comment_id int, userID string) (bool, error) {
	query := `SELECT COUNT(*) FROM like_comment WHERE comment_id = ? AND user_id = ?`
	var count int
	err := database.DB.QueryRow(query, comment_id, userID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// 点赞评论
func (l *LikeService) CreateCommentLike(comment_id int, userID string) error {
	// 创建点赞记录
	query := `INSERT INTO like_comment (comment_id, user_id) VALUES (?, ?)`
	_, err := database.DB.Exec(query, comment_id, userID)
	return err
}

// 取消评论点赞
func (l *LikeService) DeleteCommentLike(comment_id int, userID string) error {
	// 删除点赞记录
	query := `DELETE FROM like_comment WHERE comment_id = ? AND user_id = ?`
	_, err := database.DB.Exec(query, comment_id, userID)
	return err
}

// 统计评论点赞数
func (l *LikeService) GetCommentLikeCount(comment_id int) (int, error) {
	query := `SELECT COUNT(*) FROM like_comment WHERE comment_id = ?`
	var count int
	err := database.DB.QueryRow(query, comment_id).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

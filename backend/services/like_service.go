package services

import (
	"backend/database"
	"errors"
)

type LikeService struct {
}

// 是否点过赞
func (l *LikeService) HasUserLikedMoment(momentID int, userID string) (bool, error) {
	query := `SELECT COUNT(*) FROM like_info WHERE moment_id = ? AND user_id = ?`
	var count int
	err := database.DB.QueryRow(query, momentID, userID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// 点赞
func (l *LikeService) CreateMomentLike(momentID int, userID string) error {
	// 检查用户是否已经点赞
	hasLiked, err := l.HasUserLikedMoment(momentID, userID)
	if err != nil {
		return err
	}
	if hasLiked {
		return errors.New("user has already liked this moment")
	}

	// 创建点赞记录
	query := `INSERT INTO like_info (moment_id, user_id) VALUES (?, ?)`
	_, err = database.DB.Exec(query, momentID, userID)
	return err
}

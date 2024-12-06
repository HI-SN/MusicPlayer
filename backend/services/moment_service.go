package services

import (
	"backend/database"
	"backend/models"
	"time"
)

type MomentService struct {
	CService *CommentService
}

// 在数据库中创建新动态
func (m *MomentService) CreateMoment(moment *models.Moment) error {
	moment.Created_at = time.Now()

	query := `INSERT INTO moment_info (content, created_at, user_id, pic_url) VALUES (?, ?, ?, ?)`
	_, err := database.DB.Exec(query, moment.Content, moment.Created_at, moment.User_id, moment.Pic_url)
	return err
}

// 在数据库中更新动态（只有动态内容和图片路径可以更改）
func (m *MomentService) UpdateMoment(moment *models.Moment) error {
	query := `UPDATE moment_info SET content = ?, pic_url = ? WHERE id = ?`
	_, err := database.DB.Exec(query, moment.Content, moment.Pic_url, moment.Moment_id)
	return err
}

// 根据动态ID删除动态
func (m *MomentService) DeleteMoment(momentID int) error {
	// 删除动态的关联内容（评论等）
	if err := m.CService.DeleteCommentByMoment(momentID); err != nil {
		return err
	}
	// 删除动态
	query := `DELETE FROM moment_info WHERE id = ?`
	_, err := database.DB.Exec(query, momentID)
	return err
}

// 获取单条动态内容
func (m *MomentService) GetMoment(momentID int) (*models.Moment, error) {
	moment := &models.Moment{}
	query := "SELECT * FROM moment_info WHERE id=?"
	err := database.DB.QueryRow(query, momentID).Scan(&moment.Moment_id, &moment.Content,
		&moment.Created_at, &moment.User_id, &moment.Pic_url)
	if err != nil {
		return nil, err
	}
	return moment, nil
}

// 获取某用户的所有动态内容
func (m *MomentService) GetUserMoments(userID string) ([]*models.Moment, error) {
	// 执行查询
	rows, err := database.DB.Query("SELECT * FROM moment_info WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*models.Moment

	// 遍历查询结果
	for rows.Next() {
		m := &models.Moment{}
		// 使用Scan方法将列值映射到结构体的字段
		err = rows.Scan(&m.Moment_id, &m.Content, &m.Created_at, &m.User_id, &m.Pic_url)
		if err != nil {
			return nil, err
		}
		results = append(results, m)
	}

	// 检查遍历过程中是否有错误发生
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// 获取动态数量
func (m *MomentService) GetMomentsCount(userID string) (int, error) {
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM moment_info WHERE user_id = ?", userID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

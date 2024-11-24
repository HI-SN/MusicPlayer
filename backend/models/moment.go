package models

import (
	"database/sql"
	"time"
)

type Moment struct {
	Moment_id  int
	Content    string
	User_id    string
	Created_at time.Time
	Pic_url    string
}

func (Moment) TableName() string {
	return "moment_info" // 数据库中的表名
}

// 在数据库中创建新动态
func CreateMoment(db *sql.DB, m *Moment) error {
	m.Created_at = time.Now()

	query := `INSERT INTO moment_info (content, created_at, user_id, pic_url) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, m.Content, m.Created_at, m.User_id, m.Pic_url)
	return err
}

// 在数据库中更新动态（只有动态内容和图片路径可以更改）
func UpdateMoment(db *sql.DB, m *Moment) error {
	query := `UPDATE moment_info SET content = ?, pic_url = ? WHERE id = ?`
	_, err := db.Exec(query, m.Content, m.Pic_url, m.Moment_id)
	return err
}

// 根据动态ID删除动态
func DeleteMoment(db *sql.DB, momentID int) error {
	// 删除动态的关联内容（评论等）
	if err := DeleteCommentByMoment(db, momentID); err != nil {
		return err
	}
	// 删除动态
	query := `DELETE FROM moment_info WHERE id = ?`
	_, err := db.Exec(query, momentID)
	return err
}

// 获取单条动态内容
func GetMoment(db *sql.DB, momentID int) (*Moment, error) {
	m := &Moment{}
	query := "SELECT * FROM moment_info WHERE id=?"
	err := db.QueryRow(query, momentID).Scan(&m.Moment_id, &m.Content,
		&m.Created_at, &m.User_id, &m.Pic_url)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// 获取某用户的所有动态内容
func GetUserMoments(db *sql.DB, userID string) ([]*Moment, error) {
	// 执行查询
	rows, err := db.Query("SELECT * FROM moment_info WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*Moment

	// 遍历查询结果
	for rows.Next() {
		m := &Moment{}
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

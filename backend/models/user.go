package models

import (
	"database/sql"
	"time"
)

type User struct {
	User_id     string    `json:"user_id"`
	User_name   string    `json:"user_name"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Created_at  time.Time `json:"created_at"`
	Country     string    `json:"country"`
	Region      string    `json:"region"`
	Gender      string    `json:"gender"`
	Bio         string    `json:"bio"`
	Profile_pic string    `json:"profile_pic"`
	Updated_at  time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "user_info" // 数据库中的表名
}

// CreateUser 在数据库中创建新用户
func CreateUser(db *sql.DB, user *User) error {
	// 获取用户创建时间
	user.Created_at = time.Now()
	query := "INSERT INTO user_info (user_id, user_name, password, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, user.User_id, user.User_name, user.Password, user.Email, user.Created_at, user.Created_at)
	return err
}

// UpdateUser 更新现有用户
func UpdateUser(db *sql.DB, user *User) error {
	query := "UPDATE user_info SET user_name=$1, password=$2, email=$3 WHERE user_id=$4"
	_, err := db.Exec(query, user.User_name, user.Password, user.Email, user.User_id)
	return err
}

// DeleteUser 根据用户ID删除用户
func DeleteUser(db *sql.DB, userID int) error {
	query := "DELETE FROM user_info WHERE user_id=$1"
	_, err := db.Exec(query, userID)
	return err
}

// GetUser 根据用户ID获取用户信息
func GetUser(db *sql.DB, userID string) (*User, error) {
	// 先检查是否存在
	var count int
	query := "SELECT COUNT(*) FROM user_info WHERE user_id=?"
	err := db.QueryRow(query, userID).Scan(&count)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		user := &User{}
		query := "SELECT user_id, user_name, password, email FROM user_info WHERE user_id=?"
		err := db.QueryRow(query, userID).Scan(&user.User_id, &user.User_name, &user.Password, &user.Email)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	// 返回值全部为nil，说明用户不存在
	return nil, nil
}

// GetUserByEmail 根据用户邮箱获取用户信息
func GetUserByEmail(db *sql.DB, userEmail string) (*User, error) {
	// 先检查是否存在
	var count int
	query := "SELECT COUNT(*) FROM user_info WHERE email=?"
	err := db.QueryRow(query, userEmail).Scan(&count)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		user := &User{}
		query := "SELECT user_id, user_name, password, email FROM user_info WHERE email=?"
		err := db.QueryRow(query, userEmail).Scan(&user.User_id, &user.User_name, &user.Password, &user.Email)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	// 返回值全部为nil，说明用户不存在
	return nil, nil
}

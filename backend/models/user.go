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
	query := "INSERT INTO user_info (user_name, password, email) VALUES ($1, $2, $3) RETURNING user_id"
	err := db.QueryRow(query, user.User_name, user.Password, user.Email).Scan(&user.User_id)
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
func GetUser(db *sql.DB, userID int) (*User, error) {
	user := &User{}
	query := "SELECT user_id, user_name, password, email FROM user_info WHERE user_id=$1"
	err := db.QueryRow(query, userID).Scan(&user.User_id, &user.User_name, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

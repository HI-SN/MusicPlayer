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
	user.Phone = ""
	user.Country = ""
	user.Region = ""
	user.Gender = ""
	user.Bio = ""
	user.Profile_pic = ""
	query := `INSERT INTO user_info (user_id, user_name, password, email, phone, 
	created_at, country, region, gender, bio, profile_pic, updated_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, user.User_id, user.User_name, user.Password, user.Email, user.Phone,
		user.Created_at, user.Country, user.Region, user.Gender, user.Bio, user.Profile_pic, user.Created_at)
	return err
}

// UpdateUser 更新现有用户
func UpdateUser(db *sql.DB, user *User) error {
	user.Updated_at = time.Now()
	query := `UPDATE user_info SET user_name=?, password=?, email=?, phone=?,
	country=?, region=?, gender=?, bio=?, profile_pic=?, updated_at=? WHERE user_id=?`
	_, err := db.Exec(query, user.User_name, user.Password, user.Email, user.Phone,
		user.Country, user.Region, user.Gender, user.Bio, user.Profile_pic, user.Updated_at, user.User_id)
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
	user := &User{}
	query := "SELECT * FROM user_info WHERE user_id=?"
	err := db.QueryRow(query, userID).Scan(&user.User_id, &user.User_name, &user.Password,
		&user.Email, &user.Phone, &user.Created_at, &user.Country, &user.Region, &user.Gender,
		&user.Bio, &user.Profile_pic, &user.Updated_at)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByEmail 根据用户邮箱获取用户信息
func GetUserByEmail(db *sql.DB, userEmail string) (*User, error) {
	user := &User{}
	query := "SELECT * FROM user_info WHERE email=?"
	err := db.QueryRow(query, userEmail).Scan(&user.User_id, &user.User_name, &user.Password,
		&user.Email, &user.Phone, &user.Created_at, &user.Country, &user.Region, &user.Gender,
		&user.Bio, &user.Profile_pic, &user.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			// 用户不存在
			return nil, nil
		}
		// 发生了其他错误
		return nil, err
	}
	return user, nil
}

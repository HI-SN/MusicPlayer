package services

import (
	"backend/database"
	"backend/models"
	"database/sql"
	"time"
)

// UserService 定义用户相关的服务函数
type UserService struct{}

// CreateUser 在数据库中创建新用户
func (u *UserService) CreateUser(user *models.User) error {
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
	_, err := database.DB.Exec(query, user.User_id, user.User_name, user.Password, user.Email, user.Phone,
		user.Created_at, user.Country, user.Region, user.Gender, user.Bio, user.Profile_pic, user.Created_at)
	return err
}

// UpdateUser 更新现有用户
func (u *UserService) UpdateUser(user *models.User) error {
	user.Updated_at = time.Now()
	query := `UPDATE user_info SET user_name=?, email=?, phone=?,
	country=?, region=?, gender=?, bio=?, profile_pic=?, updated_at=? WHERE user_id=?`
	_, err := database.DB.Exec(query, user.User_name, user.Email, user.Phone,
		user.Country, user.Region, user.Gender, user.Bio, user.Profile_pic, user.Updated_at, user.User_id)
	return err
}

// GetUser 根据用户ID获取用户信息
func (u *UserService) GetUser(userID string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT * FROM user_info WHERE user_id=?"
	err := database.DB.QueryRow(query, userID).Scan(&user.User_id, &user.User_name, &user.Password,
		&user.Email, &user.Phone, &user.Created_at, &user.Country, &user.Region, &user.Gender,
		&user.Bio, &user.Profile_pic, &user.Updated_at)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByEmail 根据用户邮箱获取用户信息
func (u *UserService) GetUserByEmail(userEmail string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT * FROM user_info WHERE email=?"
	err := database.DB.QueryRow(query, userEmail).Scan(&user.User_id, &user.User_name, &user.Password,
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

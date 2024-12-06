package models

import (
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

// 注册时使用的结构体，方便映射；不存入数据库
type UserRegister struct {
	User
	Captcha string `json:"captcha"`
}

func (User) TableName() string {
	return "user_info" // 数据库中的表名
}

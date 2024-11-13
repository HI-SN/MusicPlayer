package controllers

import (
	"backend/database"
	"backend/models"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// UserController 定义用户相关的处理函数
type UserController struct{}

// 用户登录
func (uc *UserController) Login(c *gin.Context) {
	var newUser models.User

	// 绑定 JSON 到结构体
	if err := c.ShouldBind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "ShouldBind"})
		return
	}

	// 检查用户id、密码、邮箱等是否为空
	if len(newUser.User_id) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "账号不能为空",
		})
		return
	}
	if len(newUser.Password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return
	}

	// 检查用户邮箱是否存在
	user, err := models.GetUserByEmail(database.DB, newUser.User_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetUserByEmail failed"})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "账号不存在",
		})
		return
	}
	// }
	// 验证密码是否正确
	// 哈希密码对比
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(newUser.Password))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    500,
			"message": "账号或密码错误",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
}

// 退出登录
func (uc *UserController) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "退出登录成功"})
}

// CreateUser 处理创建用户请求
func (uc *UserController) CreateUser(c *gin.Context) {
	var newUser models.User

	// 绑定 JSON 到结构体
	if err := c.ShouldBind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "ShouldBindJSON"})
		return
	}

	// 检查用户id、密码、邮箱等是否为空
	// if len(newUser.User_id) == 0 {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{
	// 		"code":    422,
	// 		"message": "账号不能为空",
	// 	})
	// 	return
	// }
	if len(newUser.Password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return
	}
	if len(newUser.Email) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "邮箱不能为空",
		})
		return
	}

	// 检查用户邮箱是否存在
	user, err := models.GetUserByEmail(database.DB, newUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if user != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "账号已存在",
		})
		return
	}

	// 生成一个新的UUID
	newUUID := uuid.New()

	// 将UUID转换为Base62编码
	uuidBytes := newUUID[:]
	encodedUUID := base64.URLEncoding.EncodeToString(uuidBytes)
	newUser.User_id = encodedUUID[:15]

	// 哈希密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
		return
	}
	newUser.Password = string(hasedPassword)

	// 创建用户
	if err := models.CreateUser(database.DB, &newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功", "user": newUser, "hasedPassword": string(hasedPassword)})

}

// GetUser 根据ID获取用户信息
func (uc *UserController) GetUser(c *gin.Context) {
	user, err := models.GetUser(database.DB, c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User retrieved", "data": user})
}

// UpdateUser 更新用户信息
func (uc *UserController) UpdateUser(c *gin.Context) {
	var user models.User

	// 绑定 JSON 到结构体
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "ShouldBindJSON"})
		return
	}

	user.User_id = c.Param("user_id")

	err := models.UpdateUser(database.DB, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

// DeleteUser 删除用户
func (uc *UserController) DeleteUser(c *gin.Context) {
	// 这里添加删除用户的逻辑
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

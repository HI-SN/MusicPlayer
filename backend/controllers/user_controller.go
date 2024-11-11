package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
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

	// 检查用户是否存在
	user, err := models.GetUser(database.DB, newUser.User_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if user == nil {
		// 若账号查找不到用户，则用户可能输入的是邮箱，再查找一次邮箱
		user, err = models.GetUserByEmail(database.DB, newUser.User_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if user == nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "账号不存在",
			})
			return
		}
	}
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

	// 登录成功后应该跳转到相应的页面

	c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
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
	if len(newUser.Email) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "邮箱不能为空",
		})
		return
	}

	// 检查用户是否存在
	user, err := models.GetUser(database.DB, newUser.User_id)
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
	// 这里添加根据ID获取用户的逻辑
	c.JSON(http.StatusOK, gin.H{"message": "User retrieved"})
}

// UpdateUser 更新用户信息
func (uc *UserController) UpdateUser(c *gin.Context) {
	// 这里添加更新用户的逻辑
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

// DeleteUser 删除用户
func (uc *UserController) DeleteUser(c *gin.Context) {
	// 这里添加删除用户的逻辑
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

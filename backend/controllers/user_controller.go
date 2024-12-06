package controllers

import (
	"backend/database"
	"backend/models"
	"backend/services"
	"context"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// UserController 定义用户相关的处理函数
type UserController struct {
	Service  *services.UserService
	FService *services.FollowService
}

// 用户登录
func (uc *UserController) Login(c *gin.Context) {
	var a = struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Captcha  string `json:"captcha"`
	}{}
	// 绑定 JSON 到结构体
	if err := c.ShouldBind(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "ShouldBind"})
		return
	}

	// 检查用户邮箱是否存在
	user, err := uc.Service.GetUserByEmail(a.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetUserByEmail failed"})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "账号不存在",
			"data":    a,
		})
		return
	}

	// 验证密码是否正确
	if a.Password != "" {
		// 哈希密码对比
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(a.Password))
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "账号或密码错误",
				"err":     err,
			})
			return
		}
	} else {
		// 验证验证码的逻辑
		// 获取验证码和过期时间
		redisKey := "code:" + a.Email
		code, err := database.RedisClient.Get(context.Background(), redisKey).Result()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "验证码已过期或不存在"})
			return
		}

		// 检查验证码是否匹配以及是否过期
		if a.Captcha == code {
			// c.JSON(http.StatusOK, gin.H{"message": "验证码验证成功"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "验证码错误或已过期"})
		}
	}

	// 生成会话 ID
	sessionID := uuid.New().String()

	// 将会话 ID 存储到 Redis 中
	database.RedisClient.Set(context.Background(), "session:"+sessionID, user.User_id, time.Hour*24)

	// 设置 Cookie 并发送给客户端
	c.SetCookie("sessionID", sessionID, 3600*24, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"user_id": user.User_id,
	})
}

// 退出登录
func (uc *UserController) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "退出登录成功"})
}

// CreateUser 处理创建用户请求
func (uc *UserController) CreateUser(c *gin.Context) {
	var newUser models.UserRegister

	// 绑定 JSON 到结构体
	if err := c.ShouldBind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "ShouldBindJSON"})
		return
	}

	// 验证验证码的逻辑
	// 获取验证码和过期时间
	redisKey := "code:" + newUser.Email
	code, err := database.RedisClient.Get(context.Background(), redisKey).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "验证码已过期或不存在"})
		return
	}

	// 检查验证码是否匹配以及是否过期
	if newUser.Captcha == code {
		// c.JSON(http.StatusOK, gin.H{"message": "验证码验证成功"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "验证码错误或已过期"})
	}

	// 检查用户邮箱是否存在
	user, err := uc.Service.GetUserByEmail(newUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if user != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "密码加密错误",
		})
		return
	}
	newUser.Password = string(hasedPassword)

	// 创建用户
	if err := uc.Service.CreateUser(&newUser.User); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "CreateUser failed", "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功", "user_id": newUser.User_id})

}

// GetUser 根据ID获取用户信息
func (uc *UserController) GetUser(c *gin.Context) {
	user, err := uc.Service.GetUser(c.Param("user_id"))
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

	err := uc.Service.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

// // DeleteUser 删除用户
// func (uc *UserController) DeleteUser(c *gin.Context) {
// 	// 这里添加删除用户的逻辑
// 	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
// }

// 获取关注列表
func (uc *UserController) GetFollows(c *gin.Context) {
	fa, err := uc.FService.GetFollowingArtistList(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetFollowingArtistList failed"})
		return
	}
	fu, err := uc.FService.GetFollowingUserList(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetFollowingUserList failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "following list get", "artistList": fa, "userList": fu})
}

// 获取粉丝列表
func (uc *UserController) GetFollowers(c *gin.Context) {
	// fa, err := models.GetFollowerArtistList(database.DB, c.Param("user_id"))
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 	return
	// }
	fu, err := uc.FService.GetFollowerUserList(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "following list get", "userList": fu})
}

// 验证cookie的功能
func (uc *UserController) SomeProtectedEndpoint(c *gin.Context) {
	// 读取 Cookie
	sessionID, err := c.Cookie("sessionID")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权的访问"})
		return
	}

	// 从 Redis 中获取用户 ID
	userID, err := database.RedisClient.Get(context.Background(), "session:"+sessionID).Result()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "会话已过期或无效"})
		return
	}

	// 继续处理请求
	c.JSON(http.StatusOK, gin.H{
		"message": "访问成功",
		"user_id": userID,
	})
}

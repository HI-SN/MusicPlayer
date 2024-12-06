package controllers

import (
	"backend/database"
	"backend/models"
	"backend/services"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// UserController 定义用户相关的处理函数
type UserController struct {
	Service  *services.UserService
	FService *services.FollowService
	MService *services.MomentService
	Aservice *services.ArtistService
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
		err = verifyCaptcha(a.Email, a.Captcha)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
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
	err := verifyCaptcha(newUser.Email, newUser.Captcha)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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

// 找回密码
func (uc *UserController) ForgetPassword(c *gin.Context) {
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
		})
		return
	}

	// 验证验证码的逻辑
	err = verifyCaptcha(a.Email, a.Captcha)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 修改密码
	user.Password = a.Password
	err = uc.Service.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "UpdateUserPassword failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "密码修改成功"})

}

// 已登录用户修改密码
func (uc *UserController) ChangePassword(c *gin.Context) {
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

	// 绑定 JSON 到结构体
	var a = struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
		Captcha     string `json:"captcha"`
	}{}
	if err := c.ShouldBind(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "ShouldBind"})
		return
	}

	user, err := uc.Service.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetUser failed"})
		return
	}

	if a.OldPassword != "" {
		if user.Password != a.OldPassword {
			c.JSON(http.StatusBadRequest, gin.H{"error": "旧密码错误"})
			return
		}
	} else {
		// 验证验证码的逻辑
		err = verifyCaptcha(user.Email, a.Captcha)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	user.Password = a.NewPassword

	// 更新用户密码
	err = uc.Service.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "UpdateUserPassword failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "密码修改成功",
	})
}

// GetUser 根据ID获取用户信息
func (uc *UserController) GetUser(c *gin.Context) {
	user, err := uc.Service.GetUser(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "User not found"})
		return
	}

	followers_count, err := uc.FService.GetUserFollowerCount(user.User_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "User not found"})
		return
	}
	following_count, err := uc.FService.GetUserFollowingCount(user.User_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "User not found"})
		return
	}
	moments_count, err := uc.MService.GetMomentsCount(user.User_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":         user.User_id,
		"user_name":       user.User_name,
		"profile_pic":     user.Profile_pic,
		"followers_count": followers_count,
		"following_count": following_count,
		"moments_count":   moments_count,
	})
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
func (uc *UserController) GetFollowing(c *gin.Context) {
	// 先获取关注的歌手列表
	faIDs, err := uc.FService.GetFollowingArtistList(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetFollowingArtistList failed"})
		return
	}
	var artistList []*models.UserFollowArtist
	for _, id := range faIDs {
		artist, err := uc.Aservice.GetArtist(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetArtist failed"})
			return
		}
		fCount, err := uc.FService.GetArtistFollowerCount(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetArtistFollowerCount failed"})
			return
		}
		uFollowA := &models.UserFollowArtist{
			Followed_id:     id,
			Name:            artist.Name,
			Profile_pic:     artist.Profile_pic,
			Followers_count: fCount,
		}
		artistList = append(artistList, uFollowA)
	}

	// 再获取关注的用户列表
	fuIDs, err := uc.FService.GetFollowingUserList(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetFollowingUserList failed"})
		return
	}
	var userList []*models.UserFollowUser
	for _, id := range fuIDs {
		uesr, err := uc.Service.GetUser(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetUser failed"})
			return
		}
		mCount, err := uc.MService.GetMomentsCount(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetMomentsCount failed"})
			return
		}
		flerCount, err := uc.FService.GetUserFollowerCount(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetUserFollowerCount failed"})
			return
		}
		flinCount, err := uc.FService.GetUserFollowingCount(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetUserFollowingCount failed"})
			return
		}
		uFollowU := &models.UserFollowUser{
			Followed_id:     id,
			User_name:       uesr.User_name,
			Profile_pic:     uesr.Profile_pic,
			Moments_count:   mCount,
			Followers_count: flerCount,
			Following_count: flinCount,
		}
		userList = append(userList, uFollowU)
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	page_size, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// 合并列表并分页
	startIndex := (page - 1) * page_size
	endIndex := startIndex + page_size

	var pagedArtistList []*models.UserFollowArtist
	var pagedUserList []*models.UserFollowUser

	// 先填充歌手列表
	if startIndex < len(artistList) {
		if endIndex <= len(artistList) {
			pagedArtistList = artistList[startIndex:endIndex]
		} else {
			pagedArtistList = artistList[startIndex:]
			endIndex -= len(artistList)
			startIndex = 0
		}
	} else {
		startIndex -= len(artistList)
		endIndex -= len(artistList)
	}

	// 如果歌手列表不够，再填充用户列表
	if len(pagedArtistList) < page_size && startIndex < len(userList) {
		if endIndex <= len(userList) {
			pagedUserList = userList[startIndex:endIndex]
		} else {
			pagedUserList = userList[startIndex:]
		}
	}
	if pagedArtistList == nil {
		if pagedUserList == nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "超过已有的数据范围", "artistList": []interface{}{}, "userList": []interface{}{}})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "成功获取关注列表", "artistList": []interface{}{}, "userList": pagedUserList})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "成功获取关注列表", "artistList": pagedArtistList, "userList": pagedUserList})
}

// 获取粉丝列表
func (uc *UserController) GetFollowers(c *gin.Context) {
	fuIDs, err := uc.FService.GetFollowerUserList(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	var userList []*models.UserFollowUser
	for _, id := range fuIDs {
		uesr, err := uc.Service.GetUser(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetUser failed"})
			return
		}
		mCount, err := uc.MService.GetMomentsCount(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetMomentsCount failed"})
			return
		}
		flerCount, err := uc.FService.GetUserFollowerCount(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetUserFollowerCount failed"})
			return
		}
		flinCount, err := uc.FService.GetUserFollowingCount(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetUserFollowingCount failed"})
			return
		}
		uFollowU := &models.UserFollowUser{
			Followed_id:     id,
			User_name:       uesr.User_name,
			Profile_pic:     uesr.Profile_pic,
			Moments_count:   mCount,
			Followers_count: flerCount,
			Following_count: flinCount,
		}
		userList = append(userList, uFollowU)
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	page_size, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	// 合并列表并分页
	startIndex := (page - 1) * page_size
	endIndex := startIndex + page_size
	if startIndex >= len(userList) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "超过已有数据范围", "userList": []interface{}{}})
		return
	}
	if endIndex > len(userList) {
		endIndex = len(userList)
	}
	pagedFollowingList := userList[startIndex:endIndex]
	c.JSON(http.StatusOK, gin.H{"message": "成功获取粉丝列表", "userList": pagedFollowingList})
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

// 一些辅助函数
// 校验邮箱验证码
func verifyCaptcha(email, captcha string) error {
	// 获取验证码和过期时间
	redisKey := "code:" + email
	code, err := database.RedisClient.Get(context.Background(), redisKey).Result()
	if err != nil {
		return fmt.Errorf("验证码已过期或不存在")
	}

	// 检查验证码是否匹配以及是否过期
	if captcha != code {
		return fmt.Errorf("验证码错误或已过期")
	}

	return nil
}

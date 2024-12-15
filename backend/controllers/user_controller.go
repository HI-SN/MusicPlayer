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
	Service     *services.UserService
	FService    *services.FollowService
	MService    *services.MomentService
	Aservice    *services.ArtistService
	SetService  *services.SettingService
	USService   *services.UserSongService
	SongService *services.SongService
	ABService   *services.AlbumService
	ASService   *services.ArtistSongService
	PService    *services.PlaylistService
}

// 以下是登录页面相关的代码
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

// CreateUser 注册时处理创建用户请求
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

	// 创建用户默认设置
	setting := &models.Setting{
		UserID: newUser.User_id,
	}
	err = uc.SetService.CreateSetting(setting)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "CreateSetting failed", "error": err})
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

	// 哈希密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "密码加密错误",
		})
		return
	}

	// 修改密码
	user.Password = string(hasedPassword)
	err = uc.Service.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "UpdateUserPassword failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "密码修改成功"})

}

// 已登录用户修改密码
func (uc *UserController) ChangePassword(c *gin.Context) {
	// 从上下文中获取用户名
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取user_id失败"})
		return
	}
	userID := user_id.(string)

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
		// 哈希密码对比
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(a.OldPassword))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "旧密码错误"})
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

	// 哈希密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(a.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "密码加密错误",
			"error":   err,
		})
		return
	}

	// 修改密码
	user.Password = string(hasedPassword)
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
	// 从上下文中获取用户名
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取user_id失败"})
		return
	}

	var user models.User

	// 绑定 JSON 到结构体
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "ShouldBindJSON"})
		return
	}

	user.User_id = user_id.(string)

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

// 以下是我的主页相关的代码
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

// 关注其他用户
func (uc *UserController) FollowUser(c *gin.Context) {
	// 从上下文中获取用户名
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取user_id失败"})
		return
	}
	userID := user_id.(string)
	other_id := c.Param("user_id")
	fu := &models.FollowUser{
		Follower_id: userID,
		Followed_id: other_id,
	}
	err := uc.FService.CreateFollowUser(fu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "CreateFollowUser failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "关注用户成功"})
}

// 取消关注其他用户
func (uc *UserController) UnfollowUser(c *gin.Context) {
	// 从上下文中获取用户名
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取user_id失败"})
		return
	}
	userID := user_id.(string)
	other_id := c.Param("user_id")
	fu := &models.FollowUser{
		Follower_id: userID,
		Followed_id: other_id,
	}
	err := uc.FService.DeleteFollowUser(fu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "DeleteFollowUser failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "取消关注用户成功"})
}

// 关注歌手
func (uc *UserController) FollowArtist(c *gin.Context) {
	// 从上下文中获取用户名
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取user_id失败"})
		return
	}
	userID := user_id.(string)
	other_id, _ := strconv.Atoi(c.Param("artist_id"))
	fa := &models.FollowArtist{
		Follower_id: userID,
		Followed_id: other_id,
	}
	err := uc.FService.CreateFollowArtist(fa)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "CreateFollowArtist failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "关注歌手成功"})
}

// 取消关注歌手
func (uc *UserController) UnfollowArtist(c *gin.Context) {
	// 从上下文中获取用户名
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取user_id失败"})
		return
	}
	userID := user_id.(string)
	other_id, _ := strconv.Atoi(c.Param("artist_id"))
	fa := &models.FollowArtist{
		Follower_id: userID,
		Followed_id: other_id,
	}
	err := uc.FService.DeleteFollowArtist(fa)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "DeleteFollowArtist failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "取消关注歌手成功"})
}

// 以下是设置页面的相关代码
// 获取用户基础信息
func (uc *UserController) GetUserBasic(c *gin.Context) {
	// 从上下文中获取用户名
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取user_id失败"})
		return
	}
	userID := user_id.(string)

	user, err := uc.Service.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetUser failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "成功获取用户基础信息", "user": user})
}

// 获取隐私设置
func (uc *UserController) GetUserSetting(c *gin.Context) {
	// 从上下文中获取用户名
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取user_id失败"})
		return
	}
	userID := user_id.(string)

	setting, err := uc.SetService.GetSetting(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetSetting failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "成功获取用户隐私设置", "setting": setting})
}

// 更新隐私设置
func (uc *UserController) UpdateUserSetting(c *gin.Context) {
	// 从上下文中获取用户名
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取user_id失败"})
		return
	}
	userID := user_id.(string)

	var setting models.Setting
	// 绑定 JSON 到结构体
	if err := c.ShouldBind(&setting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "ShouldBindJSON"})
		return
	}

	setting.UserID = userID
	err := uc.SetService.UpdateSetting(&setting)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "UpdateSetting failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "设置更新成功", "setting": setting})
}

// 以下是我的音乐页面的相关代码
// 获取用户歌手列表
func (uc *UserController) GetUserArtist(c *gin.Context) {
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

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	page_size, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// 合并列表并分页
	startIndex := (page - 1) * page_size
	endIndex := startIndex + page_size
	if startIndex >= len(artistList) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "超过已有数据范围", "artistList": []interface{}{}})
		return
	}
	if endIndex > len(artistList) {
		endIndex = len(artistList)
	}
	pagedFollowingList := artistList[startIndex:endIndex]
	c.JSON(http.StatusOK, gin.H{"message": "成功获取用户歌手列表", "artistList": pagedFollowingList})
}

// 获取用户喜欢的歌曲列表
func (uc *UserController) GetUserLikeSong(c *gin.Context) {
	// 先获取我喜欢的歌曲的id列表
	songIDs, err := uc.USService.GetUserLikeSongList(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetUserLikeSongList failed"})
		return
	}
	var songList []*models.NewUserLikeSong
	for _, id := range songIDs {
		// 根据歌曲id获取歌曲信息、歌手信息、专辑信息
		song, err := uc.SongService.GetSongByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetSongByID failed"})
			return
		}
		album, err := uc.ABService.GetAlbumByID(song.Album_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetAlbumByID failed"})
			return
		}
		artistIDs, err := uc.ASService.GetArtistListBySongID(song.Song_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "GetArtistListBySongID failed"})
			return
		}

		uLikeSong := &models.NewUserLikeSong{
			Song:         *song,
			Album_name:   album.Name,
			Artist_ids:   make([]int, 0),    // 初始化切片
			Artist_names: make([]string, 0), // 初始化切片
		}

		for _, artistID := range artistIDs {
			artist, err := uc.Aservice.GetArtist(artistID)
			if err != nil {
				// 记录错误信息，但不中断流程
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   err,
					"message": "GetArtist failed",
					"id":      artistID,
					"artist":  artist,
					"list":    uLikeSong,
				})
				continue // 跳过当前歌手，继续处理下一个
			}
			// 确保 artist 不为 nil 后再访问其字段
			if artist != nil {
				uLikeSong.Artist_ids = append(uLikeSong.Artist_ids, artistID)
				uLikeSong.Artist_names = append(uLikeSong.Artist_names, artist.Name)
			}
		}

		songList = append(songList, uLikeSong)
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	page_size, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// 合并列表并分页
	startIndex := (page - 1) * page_size
	endIndex := startIndex + page_size
	if startIndex >= len(songList) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "超过已有数据范围", "songList": []interface{}{}})
		return
	}
	if endIndex > len(songList) {
		endIndex = len(songList)
	}
	pagedSongList := songList[startIndex:endIndex]
	c.JSON(http.StatusOK, gin.H{"message": "成功获取用户喜欢的歌曲列表", "songList": pagedSongList})
}

// 用户创建歌单
func (uc *UserController) CreatePlaylist(c *gin.Context) {
	// 从上下文中获取用户名
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取user_id失败"})
		return
	}
	userID := user_id.(string)

	var playlist models.Playlist
	// 绑定 JSON 到结构体
	if err := c.ShouldBind(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "ShouldBindJSON"})
		return
	}
	playlist.Create_at = time.Now()
	playlist.User_id = userID
	err := uc.PService.CreatePlaylist(&playlist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "CreatePlaylist faild"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "创建歌单成功", "playlist_id": playlist.Playlist_id})
}

// 获取用户创建的歌单列表
func (uc *UserController) GetUserPlaylist(c *gin.Context) {
	playLists, err := uc.PService.GetPlaylistByUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "GetPlaylistByUserID failed", "playList": playLists})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "成功获取用户创建的歌单列表", "playlist": playLists})
}

// 获取用户收藏的歌单列表

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

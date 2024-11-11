package controllers

import (
	"backend/database"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

// EmailController 定义用户相关的处理函数
type EmailController struct{}

func (ec *EmailController) SendVerification(c *gin.Context) {
	// 获取目标邮箱
	email := c.PostForm("email")

	// 生成验证码
	randSrc := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSrc)
	code := fmt.Sprintf("%06d", r.Intn(1000000))

	// 将其存储至Redis中，由于Redis为KV键值对存储所以需要定义前缀方便使用
	// 设置过期时间
	codeTimeLimit := 15
	database.RedisClient.Set(context.Background(), "code:", code, time.Minute*time.Duration(codeTimeLimit))

	// 发送验证码
	m := gomail.NewMessage()
	m.SetHeader("From", "musicplayer_sysu@126.com") // 发送者邮箱
	m.SetHeader("To", email)                        // 接收者邮箱
	m.SetHeader("Subject", "邮箱验证码")                 // 邮件主题
	m.SetBody("text/plain", "您的验证码是："+code+"，请在"+strconv.Itoa(codeTimeLimit)+"分钟内使用。")

	// 最后一个字段是网易邮箱授权码，180天更新，这是11月份的
	d := gomail.NewDialer("smtp.126.com", 465, "musicplayer_sysu@126.com", "MFxPC3AgGEXwSzbh")
	if err := d.DialAndSend(m); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "验证码发送失败", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "验证码已发送到您的邮箱"})
}

func (ec *EmailController) VerifyCode(c *gin.Context) {
	// 验证验证码的逻辑
	// 获取验证码和过期时间
	code := database.RedisClient.Get(context.Background(), "code:").Val()
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "验证码已过期或不存在"})
		return
	}
	submittedCode := c.PostForm("captcha")

	// 检查验证码是否匹配以及是否过期
	if submittedCode == code {
		c.JSON(http.StatusOK, gin.H{"message": "验证码验证成功"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "验证码错误或已过期"})
	}
}

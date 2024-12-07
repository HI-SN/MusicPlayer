package controllers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MomentController 定义用户相关的处理函数
type MomentController struct {
	Service  *services.MomentService
	CService *services.CommentService
}

// 获取单条动态及评论
func (mc *MomentController) GetMoment(c *gin.Context) {
	momentID, err := strconv.Atoi(c.Param("moment_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	moment, err := mc.Service.GetMoment(momentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "moment not found"})
		return
	}
	comments, err := mc.CService.GetAllComments(momentID, "moment")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "comment not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Moment get", "data": gin.H{"moment": moment, "comments": comments}})
}

// 发布动态
func (mc *MomentController) CreateMoment(c *gin.Context) {
	// 从上下文中获取用户名
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取user_id失败"})
		return
	}
	userID := user_id.(string)

	var moment models.Moment
	// 绑定 JSON 到结构体
	if err := c.ShouldBind(&moment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "ShouldBindJSON"})
		return
	}

	moment.User_id = userID

	// 创建动态
	if err := mc.Service.CreateMoment(&moment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Moment created", "moment_id": moment.Moment_id, "created_at": moment.Created_at})
}

// 获取某用户的所有动态
func (mc *MomentController) GetAllMoments(c *gin.Context) {
	user_id := c.Param("user_id")
	results, err := mc.Service.GetUserMoments(user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	page_size, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// 合并列表并分页
	startIndex := (page - 1) * page_size
	endIndex := startIndex + page_size

	if startIndex >= len(results) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "超过已有数据范围", "momentList": []interface{}{}})
		return
	}
	if endIndex > len(results) {
		endIndex = len(results)
	}
	pagedMomentList := results[startIndex:endIndex]
	c.JSON(http.StatusOK, gin.H{"message": "Moments get", "momentList": pagedMomentList})
}

// 修改动态
func (mc *MomentController) UpdateMoment(c *gin.Context) {
	var moment models.Moment
	// 绑定 JSON 到结构体
	err := c.ShouldBind(&moment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "ShouldBindJSON"})
		return
	}
	moment.Moment_id, err = strconv.Atoi(c.Param("moment_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Atoi"})
		return
	}
	// 创建用户
	if err := mc.Service.UpdateMoment(&moment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Moment updated"})
}

// 删除某条动态
func (mc *MomentController) DeleteMoment(c *gin.Context) {
	moment_id, err := strconv.Atoi(c.Param("moment_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = mc.Service.DeleteMoment(moment_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Moment deleted"})
}

// // 删除某用户的所有动态
// func (mc *MomentController) DeleteAllMoment(c *gin.Context) {
// 	moment_id, err := strconv.Atoi(c.Param("moment_id"))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
// 		return
// 	}
// 	err = mc.Service.DeleteMoment(moment_id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Moment deleted"})
// }

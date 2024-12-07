package controllers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	Service  *services.CommentService
	UService *services.UserService
}

// 创建动态评论
func (cc *CommentController) CreateMomentComment(c *gin.Context) {
	// 从上下文中获取用户名
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取user_id失败"})
		return
	}
	userID := user_id.(string)

	momentID, err := strconv.Atoi(c.Param("moment_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var comment models.Comment
	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "ShouldBindJSON"})
		return
	}

	comment.Type = "moment"
	comment.Target_id = momentID
	comment.User_id = userID
	// 创建评论
	if err := cc.Service.CreateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "评论成功", "comment_id": comment.Comment_id})
}

// 获取动态评论
func (cc *CommentController) GetMomentComment(c *gin.Context) {
	momentID, err := strconv.Atoi(c.Param("moment_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	comments, err := cc.Service.GetAllComments(momentID, "moment")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取评论信息失败", "error": err})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	page_size, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// 合并列表并分页
	startIndex := (page - 1) * page_size
	endIndex := startIndex + page_size

	if startIndex >= len(comments) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "超过已有数据范围", "commentList": []interface{}{}})
		return
	}
	if endIndex > len(comments) {
		endIndex = len(comments)
	}

	var results []*models.MomentComment
	for i := startIndex; i < endIndex; i++ {
		user, err := cc.UService.GetUser(comments[i].User_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "获取用户信息失败", "error": err})
			return
		}
		results = append(results, &models.MomentComment{
			Comment_id: comments[i].Comment_id,
			Content:    comments[i].Content,
			User_id:    comments[i].User_id,
			User_name:  user.User_name,
			Created_at: comments[i].Created_at,
		})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comments get", "commentList": results})
}

// 删除动态评论
func (cc *CommentController) DeleteMomentComment(c *gin.Context) {
	// 从上下文中获取用户名
	_, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取user_id失败"})
		return
	}
	// userID := user_id.(string)
	comment_id, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = cc.Service.DeleteComment(comment_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": "删除评论失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "评论删除"})
}

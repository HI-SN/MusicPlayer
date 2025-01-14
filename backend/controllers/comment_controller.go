package controllers

import (
	"backend/database"
	"backend/models"
	"backend/services"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	Service  *services.CommentService
	UService *services.UserService
	LService *services.LikeService
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

// 点赞评论
func (cc *CommentController) LikeComment(c *gin.Context) {
	// 从上下文中获取用户名
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取user_id失败"})
		return
	}
	userID := user_id.(string)

	comment_id, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	hasLiked, err := cc.LService.HasUserLikedComment(comment_id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if hasLiked {
		c.JSON(http.StatusBadRequest, gin.H{"message": "已点过赞"})
		return
	}

	err = cc.LService.CreateCommentLike(comment_id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "评论点赞成功"})
}

// 取消点赞评论
func (cc *CommentController) UnLikeComment(c *gin.Context) {
	// 从上下文中获取用户名
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取user_id失败"})
		return
	}
	userID := user_id.(string)

	comment_id, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	hasLiked, err := cc.LService.HasUserLikedComment(comment_id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if !hasLiked {
		c.JSON(http.StatusBadRequest, gin.H{"message": "未点过赞"})
		return
	}

	err = cc.LService.DeleteCommentLike(comment_id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "评论取消点赞成功"})
}

// 统计评论点赞数
func (cc *CommentController) GetCommentLikeCount(c *gin.Context) {
	comment_id, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	count, err := cc.LService.GetCommentLikeCount(comment_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get like count", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count, "message": "成功获取点赞数"})
}

type UserComment struct {
	Comment_id  int
	Content     string
	User_id     string
	Create_at   time.Time
	Type        string
	Target_id   int
	Profile_pic string // 根据user_id从user_info表中获得
}

func GetMomentComments(c *gin.Context) {
	userID := c.Param("user_id")
	var usercomments []UserComment
	db := database.DB

	// 查询动态评论的SQL语句，根据user_id获取相关评论且只筛选出类型为动态评论的记录
	query := `
        SELECT 
            ci.id, ci.content, ci.created_at, ci.user_id, ci.type, ci.target_id, ui.profile_pic
        FROM 
            comment_info ci
        JOIN 
            user_info ui ON ci.user_id = ui.user_id
		JOIN
		    moment_info mi ON mi.id = ci.target_id
        WHERE 
            mi.user_id =? AND ci.type = 'moment'
    `
	rows, err := db.Query(query, userID)
	if err != nil {
		log.Printf("查询动态评论失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取动态评论失败，请稍后再试"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var usercomment UserComment
		err := rows.Scan(&usercomment.Comment_id, &usercomment.Content, &usercomment.Create_at, &usercomment.User_id, &usercomment.Type, &usercomment.Target_id, &usercomment.Profile_pic)
		if err != nil {
			log.Printf("扫描评论数据失败: %v", err)
			continue
		}

		usercomments = append(usercomments, usercomment)
	}

	c.JSON(http.StatusOK, gin.H{"moment_comments": usercomments})
}

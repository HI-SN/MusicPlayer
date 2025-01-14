package controllers

import (
	"backend/database"
	"backend/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type CurrentChatMessages struct {
	messages     []models.Message
	sender_pic   string
	receiver_pic string
}

func GetCurrentChatMessages(c *gin.Context) {
	// 获取路径参数
	senderID := c.Param("sender_id")
	receiverID := c.Param("receiver_id")
	var currentchatmessages CurrentChatMessages
	db := database.DB

	// 查询聊天记录的SQL语句，按照时间顺序（created_at字段）升序排列
	query := `
        SELECT 
            id, created_at, sender_id, receiver_id, content, is_read
        FROM 
            message_info
        WHERE 
            (sender_id =? AND receiver_id =?) OR (sender_id =? AND receiver_id =?)
        ORDER BY 
            created_at ASC
    `
	rows, err := db.Query(query, senderID, receiverID, receiverID, senderID)
	if err != nil {
		log.Printf("查询聊天记录失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取聊天记录失败，请稍后再试"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.ID, &msg.CreatedAt, &msg.SenderID, &msg.ReceiverID, &msg.Content, &msg.IsRead)
		if err != nil {
			log.Printf("扫描聊天记录数据失败: %v", err)
			continue
		}
		currentchatmessages.messages = append(currentchatmessages.messages, msg)
	}

	// 查询当前登录用户头像的SQL语句
	query = `
        SELECT 
            profile_pic
        FROM 
            user_info
        WHERE 
            user_id = ?
    `
	rows, err = db.Query(query, senderID)
	if err != nil {
		log.Printf("查询当前登录用户头像失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取当前登录用户头像失败，请稍后再试"})
		return
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&currentchatmessages.sender_pic)
		if err != nil {
			log.Printf("扫描当前登录用户头像失败: %v", err)
		}
	}

	// 查询对方用户头像的SQL语句
	query = `
        SELECT 
            profile_pic
        FROM 
            user_info
        WHERE 
            user_id = ?
    `
	rows, err = db.Query(query, receiverID)
	if err != nil {
		log.Printf("查询当前登录用户头像失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取对方用户头像失败，请稍后再试"})
		return
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&currentchatmessages.receiver_pic)
		if err != nil {
			log.Printf("扫描对方用户头像失败: %v", err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"messages": currentchatmessages})
}

type PrivateMessage struct {
	ID          int
	CreatedAt   time.Time
	SenderID    string
	ReceiverID  string
	Content     string
	IsRead      int
	Profile_pic string
}

func GetPrivateMessageList(c *gin.Context) {
	userID := c.Param("user_id")
	var pml []PrivateMessage
	db := database.DB

	// 使用子查询和分组来获取每个聊天对象（sender或receiver）的最近一条消息
	query := `
        SELECT 
            m.id, m.created_at, m.sender_id, m.receiver_id, m.content, m.is_read
        FROM 
            (
                SELECT 
                    *, 
                    ROW_NUMBER() OVER (PARTITION BY LEAST(sender_id, receiver_id), GREATEST(sender_id, receiver_id) ORDER BY created_at DESC) AS row_num
                FROM 
                    message_info
                WHERE 
                    sender_id =? OR receiver_id =?
            ) m
        WHERE 
            m.row_num = 1
        ORDER BY 
            m.created_at DESC
    `
	rows, err := db.Query(query, userID, userID)
	if err != nil {
		log.Printf("查询私信列表失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取私信列表失败，请稍后再试"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var msg PrivateMessage
		// var created string
		err := rows.Scan(&msg.ID, &msg.CreatedAt, &msg.SenderID, &msg.ReceiverID, &msg.Content, &msg.IsRead)
		// log.Print(created)
		// msg.CreatedAt, err = time.Parse("2006-01-02 15:04:05", created)
		// if err != nil {
		// 	log.Printf("日期转换失败: %v", err)
		// 	continue
		// }
		// } else {
		// 	// 处理 created_at 为 NULL 的情况
		// 	msg.CreatedAt = time.Time{}
		// }

		if err != nil {
			log.Printf("扫描消息数据失败: %v", err)
			continue
		}

		// 查询对方用户头像的SQL语句
		query = `
			SELECT 
				profile_pic
			FROM 
				user_info
			WHERE 
				user_id = ?
		`
		if userID == msg.SenderID {
			rows, err = db.Query(query, msg.ReceiverID)
		} else {
			rows, err = db.Query(query, msg.SenderID)
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取对方用户头像失败，请稍后再试"})
			return
		}
		defer rows.Close()
		if rows.Next() {
			err := rows.Scan(&msg.Profile_pic)
			if err != nil {
				log.Printf("扫描头像失败: %v", err)
			}
		}

		pml = append(pml, msg)
	}

	c.JSON(http.StatusOK, gin.H{"messages": pml})
}

// 发送消息的接口处理函数
func SendMessage(c *gin.Context) {
	var req models.SendMessageRequest
	// 绑定JSON数据到请求结构体，如果绑定失败返回400错误
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求参数格式错误，请检查后重新发送"})
		return
	}
	// 简单验证参数，比如发送者和接收者ID不能为空等，这里可根据实际业务需求完善验证逻辑
	if req.SenderID == "" || req.ReceiverID == "" || req.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sender_id、receiver_id和content为必填参数，不能为空"})
		return
	}
	db := database.DB
	// 插入消息到数据库的SQL语句，这里假设messages表有相应的字段来存储这些信息
	query := `
        INSERT INTO message_info (sender_id, receiver_id, content, message_type, is_read)
        VALUES (?,?,?,?,0)
    `
	result, err := db.Exec(query, req.SenderID, req.ReceiverID, req.Content, req.MessageType, req.IsRead)
	if err != nil {
		log.Printf("消息插入数据库失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "消息发送失败，请稍后再试"})
		return
	}
	// 获取插入消息后生成的自增ID（如果数据库表messages的ID字段是自增类型的话），用于返回给前端
	messageID, err := result.LastInsertId()
	if err != nil {
		log.Printf("获取消息ID失败: %v", err)
		// 这里虽然获取ID失败，但消息其实已经插入成功了，所以可以选择给前端返回一个通用的成功提示而不包含ID
		c.JSON(http.StatusCreated, gin.H{"message": "消息发送成功"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "消息发送成功", "message_id": messageID})
}

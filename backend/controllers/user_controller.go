package controllers

import (
	"backend/models" // 替换为您的项目路径
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	DB *sql.DB
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 使用模型层的 CreateUser 方法
	if err := models.CreateUser(ctrl.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// 其他控制器方法...

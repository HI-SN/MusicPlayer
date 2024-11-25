package controllers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AlbumController 定义专辑相关的处理函数
type AlbumController struct {
	Service *services.AlbumService
}

// CreateAlbum 处理创建专辑请求
func (ac *AlbumController) CreateAlbum(c *gin.Context) {
	var album models.Album

	// 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用服务层函数
	err := ac.Service.CreateAlbum(&album)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album created", "album": album})
}

// GetAlbum 处理获取专辑请求
func (ac *AlbumController) GetAlbum(c *gin.Context) {
	albumID, err := strconv.Atoi(c.Param("album_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}

	// 调用服务层函数
	album, err := ac.Service.GetAlbumByID(albumID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album retrieved", "album": album})
}

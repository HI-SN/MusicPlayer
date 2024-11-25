package controllers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SongController 定义歌曲相关的处理函数
type SongController struct {
	Service *services.SongService
}

// CreateSong 处理创建歌曲请求
func (sc *SongController) CreateSong(c *gin.Context) {
	var song models.Song

	// 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用服务层函数
	err := sc.Service.CreateSong(&song)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song created", "song": song})
}

// GetSong 处理获取歌曲请求
func (sc *SongController) GetSong(c *gin.Context) {
	songID, err := strconv.Atoi(c.Param("song_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	// 调用服务层函数
	song, err := sc.Service.GetSongByID(songID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song retrieved", "song": song})
}

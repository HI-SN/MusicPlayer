package controllers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ArtistController 封装艺术家相关的控制器
type ArtistController struct {
	ArtistSongService *services.ArtistSongService
}

// NewArtistController 创建一个新的 ArtistController
func NewArtistController() *ArtistController {
	return &ArtistController{
		ArtistSongService: &services.ArtistSongService{},
	}
}

// AddArtistToSong 添加艺术家与歌曲的关系
func (ctrl *ArtistController) AddArtistToSong(c *gin.Context) {
	var relation models.ArtistSongRelation
	if err := c.ShouldBindJSON(&relation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用 service 层添加艺术家与歌曲的关系
	if err := ctrl.ArtistSongService.CreateArtistSong(&relation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Artist added to song successfully"})
}

// GetSongsByArtistID 根据艺术家 ID 获取所有相关的歌曲
func (ctrl *ArtistController) GetSongsByArtistID(c *gin.Context) {
	artistID, err := strconv.Atoi(c.Param("artistID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid artist ID"})
		return
	}

	// 调用 service 层获取歌曲信息
	songIDs, err := ctrl.ArtistSongService.GetSongListByArtistID(artistID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"song_ids": songIDs})
}

// GetArtistsBySongID 根据歌曲 ID 获取所有相关的艺术家
func (ctrl *ArtistController) GetArtistsBySongID(c *gin.Context) {
	songID, err := strconv.Atoi(c.Param("songID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	// 调用 service 层获取艺术家信息
	artistIDs, err := ctrl.ArtistSongService.GetArtistListBySongID(songID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"artist_ids": artistIDs})
}

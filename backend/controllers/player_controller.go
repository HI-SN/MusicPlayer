package controllers

import (
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PlayerController 定义播放器相关的处理函数
type PlayerController struct {
	Service *services.PlayerService
}

// PlaySong 处理播放歌曲请求
func (pc *PlayerController) PlaySong(c *gin.Context) {
	songID, err := strconv.Atoi(c.Param("song_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	// 调用服务层函数
	songURL, err := pc.Service.PlaySong(songID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Playing song", "song_url": songURL})
}

// PauseSong 处理暂停歌曲请求
func (pc *PlayerController) PauseSong(c *gin.Context) {
	songID, err := strconv.Atoi(c.Param("song_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	// 调用服务层函数
	err = pc.Service.PauseSong(songID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song paused"})
}

// ResumeSong 处理继续播放歌曲请求
func (pc *PlayerController) ResumeSong(c *gin.Context) {
	songID, err := strconv.Atoi(c.Param("song_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	// 调用服务层函数
	err = pc.Service.ResumeSong(songID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song resumed"})
}

// AdjustVolume 处理调整音量请求
func (pc *PlayerController) AdjustVolume(c *gin.Context) {
	songID, err := strconv.Atoi(c.Param("song_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	volume, err := strconv.Atoi(c.Param("volume"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid volume"})
		return
	}

	// 调用服务层函数
	err = pc.Service.AdjustVolume(songID, volume)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Volume adjusted"})
}

// ShowLyrics 处理显示歌词文件路径请求
func (pc *PlayerController) ShowLyrics(c *gin.Context) {
	songID, err := strconv.Atoi(c.Param("song_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	// 调用服务层函数
	lyricsPath, err := pc.Service.ShowLyrics(songID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lyrics path retrieved", "lyrics_path": lyricsPath})
}

package controllers

import (
	"backend/services"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadController 定义上传相关的处理函数
type UploadController struct {
	Service *services.UploadService
}

// UploadAudio 处理上传音频文件请求
func (uc *UploadController) UploadAudio(c *gin.Context) {
	file, err := c.FormFile("audio")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 打开文件
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	// 调用服务层函数
	songID, err := uc.Service.UploadAudio(src.(*os.File), file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Audio uploaded", "song_id": songID})
}

// UploadSongInfo 处理上传歌曲信息请求
func (uc *UploadController) UploadSongInfo(c *gin.Context) {
	var songInfo struct {
		Title       string    `json:"title"`
		Duration    int       `json:"duration"`
		AlbumID     int       `json:"album_id"`
		Genre       string    `json:"genre"`
		ReleaseDate time.Time `json:"release_date"`
		SongURL     string    `json:"song_url"`
		Lyrics      string    `json:"lyrics"`
		SongHit     int       `json:"song_hit"`
	}

	// 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(&songInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用服务层函数
	songID, err := uc.Service.UploadSongInfo(songInfo.Title, songInfo.Duration, songInfo.AlbumID, songInfo.Genre, songInfo.ReleaseDate, songInfo.SongURL, songInfo.Lyrics, songInfo.SongHit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song info uploaded", "song_id": songID})
}

// UploadLyrics 处理上传歌词文件请求
func (uc *UploadController) UploadLyrics(c *gin.Context) {
	songID, err := strconv.Atoi(c.Param("song_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	file, err := c.FormFile("lyrics")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 打开文件
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	// 调用服务层函数
	err = uc.Service.UploadLyrics(songID, src.(*os.File), file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lyrics uploaded"})
}

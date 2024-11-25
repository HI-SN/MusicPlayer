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

// CreatePlaylist 处理创建播放列表请求
func (pc *PlayerController) CreatePlaylist(c *gin.Context) {
	var playlist struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	// 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用服务层函数
	playlistID, err := pc.Service.CreatePlaylist(playlist.Title, playlist.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Playlist created", "playlist_id": playlistID})
}

// AddSongToPlaylist 处理添加歌曲到播放列表请求
func (pc *PlayerController) AddSongToPlaylist(c *gin.Context) {
	playlistID, err := strconv.Atoi(c.Param("playlist_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	songID, err := strconv.Atoi(c.Param("song_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	// 调用服务层函数
	err = pc.Service.AddSongToPlaylist(playlistID, songID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song added to playlist"})
}

// RemoveSongFromPlaylist 处理从播放列表移除歌曲请求
func (pc *PlayerController) RemoveSongFromPlaylist(c *gin.Context) {
	playlistID, err := strconv.Atoi(c.Param("playlist_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	songID, err := strconv.Atoi(c.Param("song_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	// 调用服务层函数
	err = pc.Service.RemoveSongFromPlaylist(playlistID, songID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song removed from playlist"})
}

// GetSongsByPlaylistID 处理获取播放列表中的所有歌曲请求
func (pc *PlayerController) GetSongsByPlaylistID(c *gin.Context) {
	playlistID, err := strconv.Atoi(c.Param("playlist_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	// 调用服务层函数
	songIDs, err := pc.Service.GetSongsByPlaylistID(playlistID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Songs retrieved", "song_ids": songIDs})
}

// ShowLyrics 处理显示歌词请求
func (pc *PlayerController) ShowLyrics(c *gin.Context) {
	songID, err := strconv.Atoi(c.Param("song_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	// 调用服务层函数
	lyrics, err := pc.Service.ShowLyrics(songID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lyrics retrieved", "lyrics": lyrics})
}

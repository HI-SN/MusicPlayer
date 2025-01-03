package controllers

import (
	"backend/models"
	"backend/services"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PlaylistController 定义播放列表相关的处理函数
type PlaylistController struct {
	Service *services.PlaylistService
}

// CreatePlaylist 处理创建播放列表请求
func (pc *PlaylistController) CreatePlaylist(c *gin.Context) {
	var playlist models.Playlist

	// 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用服务层函数
	err := pc.Service.CreatePlaylist(&playlist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Playlist created", "playlist": playlist})
}

// GetPlaylist 处理获取歌单详细信息请求
func (pc *PlaylistController) GetPlaylist(c *gin.Context) {
	playlistID, err := strconv.Atoi(c.Param("playlist_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	// 调用服务层函数
	playlist, songs, isLiked, err := pc.Service.GetPlaylistByID(playlistID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Playlist retrieved",
		"playlist": playlist,
		"songs":    songs,
		"is_liked": isLiked,
	})
}

// UpdatePlaylist 处理更新播放列表请求
func (pc *PlaylistController) UpdatePlaylist(c *gin.Context) {
	var playlist models.Playlist

	// 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取 URL 中的 playlist_id
	playlistID, err := strconv.Atoi(c.Param("playlist_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	// 设置 playlist.Playlist_id
	playlist.Playlist_id = playlistID

	// 调用服务层函数
	err = pc.Service.UpdatePlaylist(&playlist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Playlist updated", "playlist": playlist})
}

// DeletePlaylist 处理删除播放列表请求
func (pc *PlaylistController) DeletePlaylist(c *gin.Context) {
	// 获取 URL 中的 playlist_id
	playlistID, err := strconv.Atoi(c.Param("playlist_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	// 调用服务层函数
	err = pc.Service.DeletePlaylist(playlistID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Playlist deleted"})
}

// AddSongToPlaylist 处理添加歌曲到播放列表请求
func (pc *PlaylistController) AddSongToPlaylist(c *gin.Context) {
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
func (pc *PlaylistController) RemoveSongFromPlaylist(c *gin.Context) {
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
func (pc *PlaylistController) GetSongsByPlaylistID(c *gin.Context) {
	// 获取 URL 中的 playlist_id
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

// UploadPlaylistCover 处理上传歌单封面请求
func (pc *PlaylistController) UploadPlaylistCover(c *gin.Context) {
	// 获取 playlist_id
	playlistID, err := strconv.Atoi(c.Param("playlist_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	// 获取上传的文件
	file, err := c.FormFile("cover")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// 检查文件类型
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
	}
	fileHeader, _ := file.Open()
	defer fileHeader.Close()
	buffer := make([]byte, 512)
	_, err = fileHeader.Read(buffer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read file"})
		return
	}
	fileType := http.DetectContentType(buffer)
	if !allowedTypes[fileType] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type. Only JPEG, PNG, and GIF are allowed"})
		return
	}

	// 检查文件大小（限制为 5MB）
	if file.Size > 5<<20 { // 5 MB
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds the limit of 5MB"})
		return
	}

	// 确保上传目录存在
	uploadDir := "./uploads/playlist_cover"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}
	}

	// 生成文件名（每个歌单只有一个封面文件）
	ext := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("playlist_%d%s", playlistID, ext)
	filePath := filepath.Join(uploadDir, fileName)

	// 删除旧封面文件（如果存在）
	if _, err := os.Stat(filePath); err == nil {
		if err := os.Remove(filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete old cover file"})
			return
		}
	}

	// 保存新封面文件到本地
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// 生成文件的访问 URL
	fileURL := fmt.Sprintf("/uploads/playlist_cover/%s", fileName)

	// 更新歌单的 cover_url
	err = pc.Service.UpdatePlaylistCover(playlistID, fileURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cover uploaded successfully", "cover_url": fileURL})
}

// GetPlaylistsByType 处理获取推荐歌单请求
func (pc *PlaylistController) GetPlaylistsByType(c *gin.Context) {
	// 获取类型参数
	playlistType := c.Query("type")
	if playlistType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type parameter is required"})
		return
	}

	// 获取限制参数（可选，默认 10）
	limitStr := c.DefaultQuery("limit", "10") // 如果 limit 未提供，默认值为 "10"
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	// 调用服务层函数
	playlists, err := pc.Service.GetPlaylistsByType(playlistType, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Playlists retrieved", "playlists": playlists})
}

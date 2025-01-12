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

	// 从上下文中获取用户 ID，判断用户是否登录
	userID := c.GetString("user_id")
	isLoggedIn := userID != ""

	// 调用服务层函数
	songs, err := pc.Service.GetSongsByPlaylistID(playlistID, userID, isLoggedIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 构造返回的 JSON 结构
	response := gin.H{
		"data": songs,
	}

	c.JSON(http.StatusOK, response)
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

// GetSongIDsByPlaylistID 处理获取歌单下的所有歌曲ID请求
func (pc *PlaylistController) GetSongIDsByPlaylistID(c *gin.Context) {
	// 获取 URL 中的 playlist_id
	playlistID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	// 调用服务层函数
	songIDs, err := pc.Service.GetSongIDsByPlaylistID(playlistID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回歌曲ID列表
	c.JSON(http.StatusOK, songIDs)
}

// GetPlaylistsByType 处理根据歌单类型获取歌单列表的请求
func (pc *PlaylistController) GetPlaylistsByType(c *gin.Context) {
	// 获取 URL 中的 type 参数
	playlistType := c.Param("type")
	if playlistType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type parameter is required"})
		return
	}

	// 调用服务层函数
	playlists, err := pc.Service.GetPlaylistsByType(playlistType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 构造返回的 JSON 结构
	response := gin.H{
		"data": playlists,
	}

	c.JSON(http.StatusOK, response)
}

// GetPlaylistsBySearch 获取搜索结果相关的歌单
func (c *PlaylistController) GetPlaylistsBySearch(ctx *gin.Context) {
	searchKeyword := ctx.Param("keyword")

	// 调用服务层获取搜索结果
	playlists, err := c.Service.GetPlaylistsBySearch(searchKeyword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 构造返回的JSON结构
	var response struct {
		Lists []PlaylistInfo `json:"lists"`
	}

	response.Lists = make([]PlaylistInfo, 0)

	for _, playlist := range playlists {
		playlistInfo := PlaylistInfo{
			ListID: strconv.Itoa(playlist.Playlist_id),
			Title:  playlist.Title,
			Sum:    playlist.Hits,
		}
		response.Lists = append(response.Lists, playlistInfo)
	}

	ctx.JSON(http.StatusOK, response)
}

// PlaylistInfo 用于返回歌单信息的结构体
type PlaylistInfo struct {
	ListID string `json:"list_id"`
	Title  string `json:"title"`
	Sum    int    `json:"sum"`
}

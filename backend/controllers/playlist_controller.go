package controllers

import (
	"backend/database"
	"backend/models"
	"backend/services"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

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

func (pc *PlaylistController) GetPlaylist(c *gin.Context) {
	playlistID, err := strconv.Atoi(c.Param("playlist_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	// 从上下文中获取 user_id
	userID, exists := c.Get("user_id")
	if !exists {
		userID = "" // 如果未登录，设置 userID 为空字符串
	}
	isLoggedIn := userID != ""

	// 将 userID 断言为 string 类型
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	// 调用服务层函数
	playlist, songs, err := pc.Service.GetPlaylistByID(playlistID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 检查用户是否喜欢该歌单
	var isLiked bool
	if isLoggedIn {
		isLiked, err = pc.Service.IsPlaylistLikedByUser(playlistID, userIDStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		isLiked = false // 用户未登录，默认设置为 false
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

	// 从上下文中获取 user_id
	userID, exists := c.Get("user_id")
	if !exists {
		userID = "" // 如果未登录，设置 userID 为空字符串
	}
	isLoggedIn := userID != ""

	// 将 userID 断言为 string 类型
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	// 调用服务层函数
	songs, err := pc.Service.GetSongsByPlaylistID(playlistID, userIDStr, isLoggedIn)
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
		// 从 user_info 表中获取 user_name
		var userName string
		userQuery := `
			SELECT user_name
			FROM user_info
			WHERE user_id = ?
		`
		err := database.DB.QueryRow(userQuery, playlist.User_id).Scan(&userName)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				userName = "" // 如果没有找到用户，设置为空字符串
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		// 从 song_playlist_relation 表中统计 song_id 的数量
		var sum int
		sumQuery := `
			SELECT COUNT(song_id) AS sum
			FROM song_playlist_relation
			WHERE playlist_id = ?
		`
		err = database.DB.QueryRow(sumQuery, playlist.Playlist_id).Scan(&sum)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				sum = 0 // 如果没有找到歌曲，设置为 0
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		// 构造歌单信息
		playlistInfo := PlaylistInfo{
			ListID:    strconv.Itoa(playlist.Playlist_id),
			Title:     playlist.Title,
			Cover_url: playlist.Cover_url,
			User_name: userName,
			Create_at: playlist.Create_at,
			Sum:       sum,
		}
		response.Lists = append(response.Lists, playlistInfo)
	}

	ctx.JSON(http.StatusOK, response)
}

// PlaylistInfo 用于返回歌单信息的结构体
type PlaylistInfo struct {
	ListID    string    `json:"list_id"`
	Title     string    `json:"title"`
	User_name string    `json:"user_name"`
	Cover_url string    `json:"cover_url"`
	Create_at time.Time `json:"created_at"`
	Sum       int       `json:"sum"`
}

func randomPlaylists(playlists []models.PlaylistResponse) []models.PlaylistResponse {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(playlists) <= 4 {
		return playlists
	}
	// 随机打乱歌单顺序
	for i := range playlists {
		j := r.Intn(len(playlists))
		playlists[i], playlists[j] = playlists[j], playlists[i]
	}
	return playlists[:4]
}
func getPlaylistsFromDB(db *sql.DB) []models.PlaylistResponse {
	var playlists []models.PlaylistResponse
	query := "SELECT id, title, description, type, hits, cover_url FROM playlist_info"
	if db == nil {
		return playlists
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return playlists
	}
	defer rows.Close()
	for rows.Next() {
		var p models.PlaylistResponse
		err := rows.Scan(&p.Playlist_id, &p.Title, &p.Description, &p.Type, &p.Hits, &p.CoverUrl)
		if err != nil {
			log.Println(err)
			continue
		}
		playlists = append(playlists, p)
	}
	return playlists
}

func GetHomePlaylists(c *gin.Context) {
	db := database.DB
	playlists := getPlaylistsFromDB(db)
	randomPlaylists := randomPlaylists(playlists)
	c.JSON(http.StatusOK, gin.H{"playlists": randomPlaylists})
}

// 获取我喜欢的歌曲
func GetSongsOfLikelist(c *gin.Context) {
	// 获取当前用户的 ID
	userID := c.GetString("user_id") // 假设用户 ID 存储在上下文中
	isLoggedIn := userID != ""

	// 查询播放列表中的歌曲 ID
	query := `
		SELECT song_id
		FROM song_info
		JOIN user_like_song
		WHERE user_id = ?
	`
	rows, err := database.DB.Query(query, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取我喜欢的歌曲失败"})
		return
	}
	defer rows.Close()

	var songs []gin.H

	for rows.Next() {
		var songID int
		if err := rows.Scan(&songID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "获取我喜欢的歌曲失败"})
			return
		}

		// 获取歌曲详细信息（复现 GetSongsBySearch 逻辑）
		var song struct {
			ID          int
			Title       string
			Duration    int
			AlbumID     int
			Genre       string
			ReleaseDate string
			SongURL     string
			Lyrics      string
			CreatedAt   time.Time
			UpdatedAt   time.Time
			SongHit     int
		}
		songQuery := `
			SELECT id, title, duration, album_id, genre, release_date, song_url, lyrics, created_at, updated_at, song_hit
			FROM song_info
			WHERE id = ?
		`
		err := database.DB.QueryRow(songQuery, songID).Scan(
			&song.ID, &song.Title, &song.Duration, &song.AlbumID, &song.Genre, &song.ReleaseDate,
			&song.SongURL, &song.Lyrics, &song.CreatedAt, &song.UpdatedAt, &song.SongHit,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				continue // 如果没有找到歌曲，跳过
			}
			return
		}

		// 获取歌手名称（复现 GetArtistNameBySongID 逻辑）
		var artistName string
		artistQuery := `
			SELECT ai.name
			FROM artist_info ai
			JOIN artist_song_relation asr ON ai.id = asr.artist_id
			WHERE asr.song_id = ?
		`
		err = database.DB.QueryRow(artistQuery, songID).Scan(&artistName)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "获取歌手名称失败"})
			return
		}

		// 获取专辑名称（复现 GetAlbumNameByID 逻辑）
		var albumName string
		var Cover_url string
		albumQuery := `
			SELECT name, cover_url
			FROM album_info
			WHERE id = ?
		`
		err = database.DB.QueryRow(albumQuery, song.AlbumID).Scan(&albumName, &Cover_url)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "获取专辑名称失败"})
			return
		}

		// 检查用户是否喜欢该歌曲（复现 IsSongLikedByUser 逻辑）
		var isLiked bool
		if isLoggedIn {
			isLiked = true
		} else {
			isLiked = false
		}

		// 格式化时长
		minutes := song.Duration / 60
		seconds := song.Duration % 60
		formattedDuration := fmt.Sprintf("%02d:%02d", minutes, seconds)

		// 构造歌曲信息
		// 构造歌曲信息
		songInfo := gin.H{
			"id":        strconv.Itoa(song.ID),
			"title":     song.Title,
			"singer":    artistName,
			"album":     albumName,
			"album_id":  strconv.Itoa(song.AlbumID),
			"cover_url": Cover_url,
			"duration":  formattedDuration,
			"liked":     strconv.FormatBool(isLiked), // 动态设置 liked 字段
		}

		songs = append(songs, songInfo)
	}

	c.JSON(http.StatusOK, songs)
}

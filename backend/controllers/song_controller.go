package controllers

import (
	"backend/database"
	"backend/models"
	"backend/services"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"path/filepath"

	"time"

	"github.com/gin-gonic/gin"
)

type SongController struct {
	SongService *services.SongService
}

func NewSongController(songService *services.SongService) *SongController {
	return &SongController{
		SongService: songService,
	}
}

// getFFmpegPath 获取项目内的 FFmpeg 二进制文件路径
func getFFmpegPath() string {
	// 获取当前执行文件的目录
	exePath, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Printf("Failed to get executable path: %v\n", err)
		return ""
	}
	exeDir := filepath.Dir(exePath)

	// 构建 FFmpeg 二进制文件的路径
	ffmpegPath := filepath.Join(exeDir, "backend", "ffmpeg", "bin", "ffmpeg")
	if runtime.GOOS == "windows" {
		ffmpegPath += ".exe"
	}
	return ffmpegPath
}

// getAudioDuration 获取音频文件的时长（秒）
func getAudioDuration(filePath string) (int, error) {
	// 获取项目内的 FFmpeg 二进制文件路径
	ffmpegPath := getFFmpegPath()

	// 构建 ffmpeg 命令
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		// Windows 系统
		cmd = exec.Command("cmd.exe", "/C", fmt.Sprintf("%s -i \"%s\" 2>&1 | findstr \"Duration\"", ffmpegPath, filePath))
	} else {
		// Unix-like 系统
		cmd = exec.Command("bash", "-c", fmt.Sprintf("%s -i \"%s\" 2>&1 | grep 'Duration' | cut -d ' ' -f 4 | sed s/,//", ffmpegPath, filePath))
	}

	// 执行命令
	res, err := cmd.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("failed to execute command: %v, output: %s", err, string(res))
	}

	body := string(res)
	if !strings.Contains(body, ":") {
		return 0, fmt.Errorf("invalid duration format in output: %s", body)
	}

	// 提取时长信息
	timeArr := strings.Split(strings.TrimSpace(body), ":")
	if len(timeArr) != 3 {
		return 0, fmt.Errorf("invalid duration format in output: %s", body)
	}

	hour, err := strconv.ParseFloat(timeArr[0], 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse hours: %v", err)
	}
	min, err := strconv.ParseFloat(timeArr[1], 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse minutes: %v", err)
	}
	second, err := strconv.ParseFloat(strings.Split(timeArr[2], ".")[0], 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse seconds: %v", err)
	}

	duration := int(3600*hour + 60*min + second)
	return duration, nil
}

// CreateSong 创建歌曲
func (c *SongController) CreateSong(ctx *gin.Context) {
	var request struct {
		Title       string `form:"title" binding:"required"`
		ArtistID    int    `form:"artist_id" binding:"required"`
		Duration    int    `form:"duration"`
		AlbumID     int    `form:"album_id"`
		Genre       string `form:"genre"`
		ReleaseDate string `form:"release_date" binding:"required"`
	}

	// 绑定请求体
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 将字符串日期转换为 time.Time
	parsedDate, err := time.Parse("2006-01-02", request.ReleaseDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid release_date format, expected YYYY-MM-DD"})
		return
	}

	// 处理音频文件上传
	audioFile, err := ctx.FormFile("audio")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "audio file is required"})
		return
	}

	// 保存音频文件到 /uploads/audio
	audioDir := "uploads/audio"
	if err := os.MkdirAll(audioDir, 0755); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create audio directory"})
		return
	}

	// 调用服务层创建歌曲
	songID, err := c.SongService.CreateSong(request.Title, request.ArtistID, request.Duration, request.AlbumID, request.Genre, parsedDate, "", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 生成唯一的音频文件名
	audioExt := filepath.Ext(audioFile.Filename)
	audioFileName := fmt.Sprintf("audio_%d%s", songID, audioExt)
	audioFilePath := filepath.Join(audioDir, audioFileName)
	if err := ctx.SaveUploadedFile(audioFile, audioFilePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save audio file"})
		return
	}

	// 获取音频文件的时长
	duration, err := getAudioDuration(audioFilePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get audio duration"})
		return
	}

	// 更新歌曲的音频文件路径和时长
	err = c.SongService.UpdateSongInfo(int(songID), request.Title, duration, request.AlbumID, request.Genre, parsedDate, audioFilePath, "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 处理歌词文件上传（可选）
	lyricsFile, err := ctx.FormFile("lyrics")
	if err == nil {
		// 保存歌词文件到 /uploads/lyrics
		lyricsDir := "uploads/lyrics"
		if err := os.MkdirAll(lyricsDir, 0755); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create lyrics directory"})
			return
		}

		// 生成唯一的歌词文件名
		lyricsExt := filepath.Ext(lyricsFile.Filename)
		lyricsFileName := fmt.Sprintf("lyrics_%d%s", songID, lyricsExt)
		lyricsFilePath := filepath.Join(lyricsDir, lyricsFileName)
		if err := ctx.SaveUploadedFile(lyricsFile, lyricsFilePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save lyrics file"})
			return
		}

		// 更新歌曲的歌词文件路径
		err = c.SongService.UploadLyricsBySongID(int(songID), lyricsFilePath)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "song created successfully", "song_id": songID})
}

// GetSongByID retrieves a song by its ID
func (c *SongController) GetSongByID(ctx *gin.Context) {
	songID, err := strconv.Atoi(ctx.Param("song_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid song ID"})
		return
	}

	// 从上下文中获取 user_id
	userID, exists := ctx.Get("user_id")
	if !exists {
		userID = "" // 如果未登录，设置 userID 为空字符串
	}
	isLoggedIn := userID != ""

	// 将 userID 断言为 string 类型
	userIDStr, ok := userID.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	// 调用服务层函数
	song, artistName, albumName, Cover_url, liked, err := c.SongService.GetSongByID(songID, userIDStr, isLoggedIn)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// 构造返回的 JSON 结构
	ctx.JSON(http.StatusOK, gin.H{
		"id":           song.Song_id,
		"title":        song.Title,
		"artist_name":  artistName,
		"duration":     song.Duration,
		"album_name":   albumName,
		"cover_url":    Cover_url,
		"album_id":     song.Album_id,
		"genre":        song.Genre,
		"release_date": song.Release_date,
		"song_url":     song.Song_url,
		"lyrics":       song.Lyrics,
		"created_at":   song.Created_at,
		"updated_at":   song.Updated_at,
		"song_hit":     song.Song_hit,
		"liked":        liked, // 添加 liked 字段
	})
}

// UpdateSongInfo 更新歌曲信息
func (c *SongController) UpdateSongInfo(ctx *gin.Context) {
	songID, err := strconv.Atoi(ctx.Param("song_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid song ID"})
		return
	}

	var request struct {
		Title       string `json:"title"`
		Duration    int    `json:"duration"`
		AlbumID     int    `json:"album_id"`
		Genre       string `json:"genre"`
		ReleaseDate string `json:"release_date"`
		SongURL     string `json:"song_url"`
		Lyrics      string `json:"lyrics"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 将字符串日期转换为 time.Time
	parsedDate, err := time.Parse("2006-01-02", request.ReleaseDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid release_date format, expected YYYY-MM-DD"})
		return
	}

	// 处理音频文件上传（可选）
	audioFile, err := ctx.FormFile("audio")
	if err == nil {
		// 保存音频文件到 /uploads/audio
		audioDir := "uploads/audio"
		if err := os.MkdirAll(audioDir, 0755); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create audio directory"})
			return
		}

		// 生成唯一的音频文件名
		audioExt := filepath.Ext(audioFile.Filename)
		audioFileName := fmt.Sprintf("audio_%d%s", songID, audioExt)
		audioFilePath := filepath.Join(audioDir, audioFileName)
		if err := ctx.SaveUploadedFile(audioFile, audioFilePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save audio file"})
			return
		}

		// 获取音频文件的时长
		duration, err := getAudioDuration(audioFilePath)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get audio duration"})
			return
		}

		// 更新音频文件路径和时长
		request.SongURL = audioFilePath
		request.Duration = duration
	}

	// 处理歌词文件上传（可选）
	lyricsFile, err := ctx.FormFile("lyrics")
	if err == nil {
		// 保存歌词文件到 /uploads/lyrics
		lyricsDir := "uploads/lyrics"
		if err := os.MkdirAll(lyricsDir, 0755); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create lyrics directory"})
			return
		}

		// 生成唯一的歌词文件名
		lyricsExt := filepath.Ext(lyricsFile.Filename)
		lyricsFileName := fmt.Sprintf("lyrics_%d%s", songID, lyricsExt)
		lyricsFilePath := filepath.Join(lyricsDir, lyricsFileName)

		// 删除旧歌词文件（如果存在）
		if _, err := os.Stat(lyricsFilePath); err == nil {
			if err := os.Remove(lyricsFilePath); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete old lyrics file"})
				return
			}
		}

		// 保存新歌词文件
		if err := ctx.SaveUploadedFile(lyricsFile, lyricsFilePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save lyrics file"})
			return
		}

		// 更新歌词文件路径
		request.Lyrics = lyricsFilePath
	}

	// 调用服务层更新歌曲信息
	err = c.SongService.UpdateSongInfo(songID, request.Title, request.Duration, request.AlbumID, request.Genre, parsedDate, request.SongURL, request.Lyrics)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "song updated successfully"})
}

// UploadLyricsBySongID uploads lyrics for a song
func (c *SongController) UploadLyricsBySongID(ctx *gin.Context) {
	songID, err := strconv.Atoi(ctx.Param("song_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid song ID"})
		return
	}

	// 获取上传的歌词文件
	lyricsFile, err := ctx.FormFile("lyrics")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "lyrics file is required"})
		return
	}

	// 保存歌词文件到 /uploads/lyrics
	lyricsDir := "uploads/lyrics"
	if err := os.MkdirAll(lyricsDir, 0755); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create lyrics directory"})
		return
	}

	// 生成唯一的歌词文件名
	lyricsExt := filepath.Ext(lyricsFile.Filename)
	lyricsFileName := fmt.Sprintf("lyrics_%d%s", songID, lyricsExt)
	lyricsFilePath := filepath.Join(lyricsDir, lyricsFileName)
	if err := ctx.SaveUploadedFile(lyricsFile, lyricsFilePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save lyrics file"})
		return
	}

	// 更新歌曲的歌词文件路径
	err = c.SongService.UploadLyricsBySongID(int(songID), lyricsFilePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "lyrics uploaded successfully"})
}

// DownloadAudioBySongID downloads the audio file of a song
func (c *SongController) DownloadAudioBySongID(ctx *gin.Context) {
	songID, err := strconv.Atoi(ctx.Param("song_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid song ID"})
		return
	}

	filePath, err := c.SongService.DownloadAudioBySongID(songID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "audio file not found"})
		return
	}

	ctx.File(filePath)
}

// DownloadLyricsBySongID downloads the lyrics of a song
func (c *SongController) DownloadLyricsBySongID(ctx *gin.Context) {
	songID, err := strconv.Atoi(ctx.Param("song_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid song ID"})
		return
	}

	lyrics, err := c.SongService.DownloadLyricsBySongID(songID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.String(http.StatusOK, lyrics)
}

// DeleteSongByID 删除歌曲
func (c *SongController) DeleteSongByID(ctx *gin.Context) {
	songID, err := strconv.Atoi(ctx.Param("song_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid song ID"})
		return
	}

	// 调用服务层删除歌曲
	err = c.SongService.DeleteSongByID(songID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "song deleted successfully"})
}

// GetCommentsBySongID 获取歌曲相关评论
func (c *SongController) GetCommentsBySongID(ctx *gin.Context) {
	songID, err := strconv.Atoi(ctx.Param("song_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid song ID"})
		return
	}

	// 调用服务层获取评论
	comments, err := c.SongService.GetCommentsBySongID(songID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"comments": comments})
}

// GetSongsBySearch 获取搜索结果的歌曲信息
func (c *SongController) GetSongsBySearch(ctx *gin.Context) {
	searchKeyword := ctx.Param("search")

	// 从上下文中获取 user_id
	userID, exists := ctx.Get("user_id")
	if !exists {
		userID = "" // 如果未登录，设置 userID 为空字符串
	}
	isLoggedIn := userID != ""

	// 将 userID 断言为 string 类型
	userIDStr, ok := userID.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	// 调用服务层获取搜索结果
	songs, err := c.SongService.GetSongsBySearch(searchKeyword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 如果没有找到歌曲，返回空列表
	if len(songs) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 49,
			"data": []interface{}{},
		})
		return
	}

	// 构造返回的 JSON 结构
	var response struct {
		Code int        `json:"code"`
		Data []SongInfo `json:"data"`
	}

	response.Code = 49
	response.Data = make([]SongInfo, 0)

	for _, song := range songs {
		// 获取歌手名称
		artistName, err := c.SongService.GetArtistNameBySongID(song.Song_id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 获取专辑名称
		albumName, Cover_url, err := c.SongService.GetAlbumByID(song.Album_id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 检查用户是否喜欢该歌曲
		var isLiked bool
		if isLoggedIn {
			// 用户已登录，查询是否喜欢该歌曲
			isLiked, err = c.SongService.IsSongLikedByUser(song.Song_id, userIDStr)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			// 用户未登录，默认设置为 false
			isLiked = false
		}

		// 构造歌曲信息
		songInfo := SongInfo{
			ID:        strconv.Itoa(song.Song_id),
			Title:     song.Title,
			Singer:    artistName,
			Album:     albumName,
			Cover_url: Cover_url,
			Album_id:  song.Album_id,
			IfLike:    strconv.FormatBool(isLiked),
			Time:      formatDuration(song.Duration),
		}
		response.Data = append(response.Data, songInfo)
	}

	ctx.JSON(http.StatusOK, response)
}

// SongInfo 用于返回歌曲信息的结构体
type SongInfo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Singer    string `json:"singer"`
	Album     string `json:"album"`
	Cover_url string `json:"cover_url"`
	Album_id  int    `json:"album_id"`
	IfLike    string `json:"liked"`
	Time      string `json:"duration"`
}

// formatDuration 将秒数转换为 "mm:ss" 格式
func formatDuration(duration int) string {
	minutes := duration / 60
	seconds := duration % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

func GetSongsByArtistID(artistID int, userID string, isLoggedIn bool) ([]models.Song_ranking_detail, error) {
	var songs []models.Song_ranking_detail
	db := database.DB
	query := "SELECT id, title, duration, album_id, genre, release_date, song_url, lyrics, song_hit FROM song_info si join artist_song_relation asr on si.id = asr.song_id WHERE artist_id =?"
	rows, err := db.Query(query, artistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var song models.Song_ranking_detail
		err := rows.Scan(&song.ID, &song.Title, &song.Duration, &song.AlbumID, &song.Genre, &song.ReleaseDate, &song.SongUrl, &song.Lyrics, &song.SongHit)
		if err != nil {
			return nil, err
		}
		// 检查用户是否喜欢该歌曲
		if isLoggedIn {
			var count int
			likeQuery := `
				SELECT COUNT(*)
				FROM user_like_song
				WHERE user_id = ? AND song_id = ?
			`
			err := db.QueryRow(likeQuery, userID, song.ID).Scan(&count)
			if err != nil {
				// log.Printf("检查用户是否喜欢该歌曲失败: %v", err)
				continue
			}
			if count > 0 {
				song.Liked = "true"
			} else {
				song.Liked = "false"
			}
		} else {
			song.Liked = "false" // 用户未登录，默认设置为 false
		}
		songs = append(songs, song)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return songs, nil
}

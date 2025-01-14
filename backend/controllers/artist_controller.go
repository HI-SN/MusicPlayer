package controllers

import (
	"backend/database"
	"backend/models"
	"backend/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ArtistController 封装艺术家相关的控制器
type ArtistController struct {
	ArtistSongService *services.ArtistSongService
	ArtistService     *services.ArtistService
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

// GetArtistsBySearch 获取搜索结果相关的歌手
func (c *ArtistController) GetArtistsBySearch(ctx *gin.Context) {
	searchKeyword := ctx.Param("keyword")

	// 调用服务层获取搜索结果
	artists, err := c.ArtistService.GetArtistsBySearch(searchKeyword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 构造返回的JSON结构
	var response struct {
		Singers []SingerInfo `json:"singers"`
	}

	response.Singers = make([]SingerInfo, 0)

	for _, artist := range artists {
		singerInfo := SingerInfo{
			SingerID:    strconv.Itoa(artist.Artist_id),
			Name:        artist.Name,
			Profile_pic: artist.Profile_pic,
		}
		response.Singers = append(response.Singers, singerInfo)
	}

	ctx.JSON(http.StatusOK, response)
}

// SingerInfo 用于返回歌手信息的结构体
type SingerInfo struct {
	SingerID    string `json:"singer_id"`
	Name        string `json:"name"`
	Profile_pic string `json:"url"`
}

// Artist 歌手信息结构体
type ArtistDetail struct {
	Artist      models.Artist
	Is_followed string
	Songs       []models.Song_ranking_detail `json:"songs"`
}

// GetArtistByID 从数据库根据 id 获取歌手信息
func GetArtistDetailByID(c *gin.Context) {
	id := c.Param("id")
	// 获取当前用户的 ID
	userID := c.GetString("user_id")
	isLoggedIn := userID != ""
	var response ArtistDetail
	var artist models.Artist
	db := database.DB
	// 查询歌手信息
	query := "SELECT id, name, bio, profile_pic, type, nation FROM artist_info WHERE id =?"
	err := db.QueryRow(query, id).Scan(&artist.Artist_id, &artist.Name, &artist.Bio, &artist.Profile_pic, &artist.Type, &artist.Nation)
	if err != nil {
		log.Printf("查询歌手信息失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取歌手信息失败，请稍后再试"})
		return
	}
	// 检查用户是否关注该歌手
	if isLoggedIn {
		var count int
		followQuery := `
			SELECT COUNT(*)
			FROM follow_artist
			WHERE follower_id = ? AND followed_id = ?
		`
		err := db.QueryRow(followQuery, userID, id).Scan(&count)
		if err != nil {
			log.Printf("获取歌手的歌曲信息失败: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取歌手的歌曲信息失败！"})
			return
		}
		if count > 0 {
			response.Is_followed = "true"
		} else {
			response.Is_followed = "false"
		}
	} else {
		response.Is_followed = "false" // 用户未登录，默认设置为 false
	}
	// 获取该歌手的歌曲信息
	songs, err := GetSongsByArtistID(artist.Artist_id, userID, isLoggedIn)
	if err != nil {
		log.Printf("获取歌手的歌曲信息失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取歌手的歌曲信息失败！"})
		return
	}

	response.Artist = artist
	response.Songs = songs
	c.JSON(http.StatusOK, response)
}

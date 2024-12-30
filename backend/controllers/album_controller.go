package controllers

import (
	"backend/database"
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// AlbumController 定义专辑相关的处理函数
type AlbumController struct {
	Service *services.AlbumService
}

// CreateAlbum 处理创建专辑请求
func (ac *AlbumController) CreateAlbum(c *gin.Context) {
	var request struct {
		Name         string `json:"name" binding:"required"`
		Description  string `json:"description"`
		Release_date string `json:"release_date" binding:"required"` // 接收字符串
	}

	// 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 将 Release_date 字符串解析为 time.Time
	releaseDate, err := time.Parse("2006-01-02", request.Release_date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid release_date format, expected YYYY-MM-DD"})
		return
	}

	// 创建专辑对象
	album := &models.Album{
		Name:         request.Name,
		Description:  request.Description,
		Release_date: releaseDate, // 使用解析后的 time.Time
	}

	// 调用服务层函数创建专辑
	err = ac.Service.CreateAlbum(album)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 获取上传的封面文件（非必须）
	coverFile, err := c.FormFile("cover")
	if err == nil {
		// 上传封面文件
		coverURL, err := ac.Service.UploadAlbumCover(album.Album_id, coverFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 更新专辑的封面路径
		album.Cover_url = coverURL
		err = ac.Service.UpdateAlbum(album.Album_id, album.Name, album.Description, album.Release_date, coverFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album created", "album": album})
}

// UploadAlbumCover 处理上传专辑封面请求
func (ac *AlbumController) UploadAlbumCover(c *gin.Context) {
	// 获取专辑 ID
	albumID, err := strconv.Atoi(c.Param("album_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}

	// 获取上传的文件
	file, err := c.FormFile("cover")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cover file is required"})
		return
	}

	// 调用服务层函数上传封面
	coverURL, err := ac.Service.UploadAlbumCover(albumID, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 更新专辑的封面 URL
	query := "UPDATE album_info SET cover_url=? WHERE id=?"
	_, err = database.DB.Exec(query, coverURL, albumID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cover uploaded", "cover_url": coverURL})
}

// UpdateAlbum 处理更新专辑请求
func (ac *AlbumController) UpdateAlbum(c *gin.Context) {
	albumID, err := strconv.Atoi(c.Param("album_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}

	var request struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		Release_date string `json:"release_date"` // 接收字符串
	}

	// 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 将 Release_date 字符串解析为 time.Time
	var releaseDate time.Time
	if request.Release_date != "" {
		releaseDate, err = time.Parse("2006-01-02", request.Release_date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid release_date format, expected YYYY-MM-DD"})
			return
		}
	}

	// 获取上传的封面文件（非必须）
	coverFile, _ := c.FormFile("cover")

	// 调用服务层函数更新专辑信息
	err = ac.Service.UpdateAlbum(albumID, request.Name, request.Description, releaseDate, coverFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album updated"})
}

// GetAlbum 处理获取专辑请求
func (ac *AlbumController) GetAlbum(c *gin.Context) {
	albumID, err := strconv.Atoi(c.Param("album_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}

	// 调用服务层函数获取专辑信息
	album, artists, songs, err := ac.Service.GetAlbumByID(albumID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Album retrieved",
		"album":   album,
		"artists": artists,
		"songs":   songs,
	})
}

// DeleteAlbum 处理删除专辑请求
func (ac *AlbumController) DeleteAlbum(c *gin.Context) {
	albumID, err := strconv.Atoi(c.Param("album_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}

	// 调用服务层函数删除专辑
	err = ac.Service.DeleteAlbum(albumID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album deleted"})
}

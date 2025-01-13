package controllers

import (
	"backend/database"
	"backend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// SearchResults结构体用于统一封装所有类型的搜索结果
type SearchResults struct {
	Songs     []models.Song     `json:"songs"`
	Artists   []models.Artist   `json:"artists"`
	Playlists []models.Playlist `json:"playlists"`
	Albums    []models.Album    `json:"albums"`
}

func Search(c *gin.Context) {
	db := database.DB
	query := c.Param("str")
	var results SearchResults
	// 查询歌曲
	songQuery := `SELECT * FROM song_info WHERE title LIKE ?`
	log.Print(songQuery)
	songRows, err := db.Query(songQuery, "%"+query+"%")
	if err != nil {
		log.Printf("查询歌曲失败: %v", err)
	}
	defer songRows.Close()
	for songRows.Next() {
		var song models.Song
		err := songRows.Scan(&song.Song_id, &song.Title, &song.Duration, &song.Album_id, &song.Genre, &song.Release_date, &song.Song_url, &song.Lyrics, &song.Created_at, &song.Updated_at, &song.Song_hit)
		if err != nil {
			log.Printf("扫描歌曲数据失败: %v", err)
			continue
		}
		results.Songs = append(results.Songs, song)
	}

	// 查询歌手
	artistQuery := `SELECT * FROM artist_info WHERE name LIKE ?`
	artistRows, err := db.Query(artistQuery, "%"+query+"%")
	if err != nil {
		log.Printf("查询歌手失败: %v", err)
	}
	defer artistRows.Close()
	for artistRows.Next() {
		var artist models.Artist
		err := artistRows.Scan(&artist.Artist_id, &artist.Name, &artist.Bio, &artist.Profile_pic, &artist.Type, &artist.Nation)
		if err != nil {
			log.Printf("扫描歌手数据失败: %v", err)
			continue
		}
		results.Artists = append(results.Artists, artist)
	}

	// 查询歌单
	playlistQuery := `SELECT * FROM playlist_info WHERE title LIKE ?`
	playlistRows, err := db.Query(playlistQuery, "%"+query+"%")
	if err != nil {
		log.Printf("查询歌单失败: %v", err)
	}
	defer playlistRows.Close()
	for playlistRows.Next() {
		var playlist models.Playlist
		err := playlistRows.Scan(&playlist.Playlist_id, &playlist.Title, &playlist.User_id, &playlist.Create_at, &playlist.Description, &playlist.Type, &playlist.Hits, &playlist.Cover_url)
		if err != nil {
			log.Printf("扫描歌单数据失败: %v", err)
			continue
		}
		results.Playlists = append(results.Playlists, playlist)
	}

	// 查询专辑
	albumQuery := `SELECT * FROM album_info WHERE name LIKE?`
	albumRows, err := db.Query(albumQuery, "%"+query+"%")
	if err != nil {
		log.Printf("查询专辑失败: %v", err)
	}
	defer albumRows.Close()
	for albumRows.Next() {
		var album models.Album
		err := albumRows.Scan(&album.Album_id, &album.Name, &album.Description, &album.Release_date, &album.Cover_url)
		if err != nil {
			log.Printf("扫描专辑数据失败: %v", err)
			continue
		}
		results.Albums = append(results.Albums, album)
	}

	c.JSON(http.StatusOK, results)
}

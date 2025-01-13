package controllers

import (
	"backend/database"
	"backend/models"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func getRankingInfoByName(db *sql.DB, rankingName string) (models.Ranking_home, error) {
	var ranking models.Ranking_home
	// 查询排行榜基本信息（名称、封面URL）
	// rankingQuery := "SELECT ranking_name, cover_url FROM ranking_info where name=?"
	// err := db.QueryRow(rankingQuery, rankingName).Scan(&rankingName, &ranking.CoverUrl)
	// if err != nil {
	// 	return ranking, err
	// }
	// 查询该排行榜前三首歌曲信息
	// songsQuery := "SELECT id, title FROM song WHERE ranking_name =? ORDER BY ranking_order LIMIT 3"
	songsQuery := `
        SELECT 
            s.id, s.title
        FROM 
            song_info s
        JOIN 
            ranking_info r ON s.id = r.song_id
        WHERE 
            r.name =?
		ORDER BY 
			'rank'
		LIMIT 3
    `
	rows, err := db.Query(songsQuery, rankingName)
	if err != nil {
		return ranking, err
	}
	defer rows.Close()
	for rows.Next() {
		var song models.Song_rank_home
		err := rows.Scan(&song.ID, &song.Title)
		if err != nil {
			return ranking, err
		}
		ranking.Songs = append(ranking.Songs, song)
	}
	return ranking, nil
}

func GetHomeRanking(c *gin.Context) {
	rankingName := c.Param("name")
	db := database.DB
	rankingInfo, err := getRankingInfoByName(db, rankingName)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取排行榜信息失败"})
		return
	}
	c.JSON(http.StatusOK, rankingInfo)
}

// 获取指定排行榜详情
func GetRankDetailsByName(c *gin.Context) {
	rankingName := c.Param("name")
	// 获取当前用户的 ID
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取user_id失败"})
		return
	}
	userID := user_id.(string)
	isLoggedIn := userID != ""
	var songs []models.Song_ranking_detail
	db := database.DB
	if db == nil {
		log.Println("数据库连接未初始化")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "数据库连接异常，请稍后再试"})
		return
	}
	query := `
        SELECT 
            s.id, s.title, s.duration, s.album_id, s.genre, s.release_date, s.song_url, s.lyrics, s.song_hit
        FROM 
            song_info s
        JOIN 
            ranking_info r ON s.id = r.song_id
        WHERE 
            r.name =?
    `
	rows, err := db.Query(query, rankingName)
	if err != nil {
		log.Printf("查询排行榜详情失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取排行榜详情失败，请稍后再试"})
		return
	}
	defer rows.Close()
	// 检查结果
	if !rows.Next() {
		log.Println("结果为空")
		return // 或其他处理空结果的代码
	}

	for rows.Next() {
		var song models.Song_ranking_detail
		err := rows.Scan(&song.ID, &song.Title, &song.Duration, &song.AlbumID, &song.Genre, &song.ReleaseDate, &song.SongUrl, &song.Lyrics, &song.SongHit)
		if err != nil {
			log.Printf("扫描歌曲详情数据失败: %v", err)
			continue
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
				log.Printf("检查用户是否喜欢该歌曲失败: %v", err)
				continue
			}
			song.Liked = strconv.FormatBool(count > 0)
		} else {
			song.Liked = "false" // 用户未登录，默认设置为 false
		}
		songs = append(songs, song)
	}

	c.JSON(http.StatusOK, gin.H{"rankings": songs})
}

// AddSongToPlaylist 将歌曲添加到歌单
func AddSongToPlaylist(c *gin.Context) {
	db := database.DB
	if db == nil {
		log.Println("数据库连接未初始化")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "数据库连接异常，请稍后再试"})
		return
	}
	songID := c.Param("song_id")
	playlistID := c.Param("playlist_id")
	// 插入歌曲到歌单的 SQL 语句
	query := "INSERT INTO song_playlist_relation (playlist_id, song_id) VALUES (?,?)"
	_, err := db.Exec(query, playlistID, songID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "500"})
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "将歌曲加入歌单失败！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "200"})
}

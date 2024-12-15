package models

import (
	"time"
)

type Song struct {
	Song_id      int       `json:"song_id"`
	Title        string    `json:"title"`
	Duration     int       `json:"duration"`
	Album_id     int       `json:"album_id"`
	Genre        string    `json:"genre"`
	Release_date time.Time `json:"release_date"`
	Song_url     string    `json:"song_url"`
	Lyrics       string    `json:"lyrics"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
	Song_hit     int       `json:"song_hit"`
}

// func (Song) TableName() string {
// 	return "song_info" // 数据库中的表名
// }

// // CreateUser 在数据库中创建新歌曲
// // func CreateSong(db *sql.DB, song *Song) error {
// // 	query := "INSERT INTO song_info (title, duration, album_id, genre, release_date, song_url, lyrics, created_at, updated_at, song_hit) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING song_id"
// // 	err := db.QueryRow(query, song.Title, song.Duration, song.Album_id, song.Genre, song.Release_date, song.Song_url, song.Lyrics, song.Created_at, song.Updated_at, song.Song_hit).Scan(&song.Song_id)
// // 	return err
// // }

// // CreateSong 在数据库中创建新歌曲
// func CreateSong(db *sql.DB, song *Song) error {
// 	// 构建插入 SQL 语句和参数
// 	query := "INSERT INTO song_info ("
// 	values := []interface{}{}
// 	placeholders := []string{}

// 	if song.Title != nil {
// 		query += "title,"
// 		placeholders = append(placeholders, "$"+strconv.Itoa(len(values)+1))
// 		values = append(values, *song.Title)
// 	}
// 	if song.Duration != nil {
// 		query += "duration,"
// 		placeholders = append(placeholders, "$"+strconv.Itoa(len(values)+1))
// 		values = append(values, *song.Duration)
// 	}
// 	if song.Album_id != nil {
// 		query += "album_id,"
// 		placeholders = append(placeholders, "$"+strconv.Itoa(len(values)+1))
// 		values = append(values, *song.Album_id)
// 	}
// 	if song.Genre != nil {
// 		query += "genre,"
// 		placeholders = append(placeholders, "$"+strconv.Itoa(len(values)+1))
// 		values = append(values, *song.Genre)
// 	}
// 	if song.Release_date != nil {
// 		query += "release_date,"
// 		placeholders = append(placeholders, "$"+strconv.Itoa(len(values)+1))
// 		values = append(values, *song.Release_date)
// 	}
// 	if song.Song_url != nil {
// 		query += "song_url,"
// 		placeholders = append(placeholders, "$"+strconv.Itoa(len(values)+1))
// 		values = append(values, *song.Song_url)
// 	}
// 	if song.Lyrics != nil {
// 		query += "lyrics,"
// 		placeholders = append(placeholders, "$"+strconv.Itoa(len(values)+1))
// 		values = append(values, *song.Lyrics)
// 	}
// 	if song.Song_hit != nil {
// 		query += "song_hit,"
// 		placeholders = append(placeholders, "$"+strconv.Itoa(len(values)+1))
// 		values = append(values, *song.Song_hit)
// 	}

// 	// 添加时间戳信息
// 	song.Created_at = time.Now()
// 	song.Updated_at = time.Now()
// 	query += "created_at, updated_at) VALUES ("
// 	query += strings.Join(placeholders, ", ")
// 	query += ", $" + strconv.Itoa(len(values)+1) + ", $" + strconv.Itoa(len(values)+2) + ") RETURNING song_id"
// 	values = append(values, song.Created_at, song.Updated_at)

// 	// 执行插入操作
// 	return db.QueryRow(query, values...).Scan(&song.Song_id)
// }

// // UpdateUser 更新现有用户
// func UpdateSong(db *sql.DB, song *Song) error {
// 	query := "UPDATE user_info SET user_name=$1, password=$2, email=$3 WHERE user_id=$4"
// 	_, err := db.Exec(query, user.User_name, user.Password, user.Email, user.User_id)
// 	return err
// }

// // DeleteUser 根据用户ID删除用户
// func DeleteUser(db *sql.DB, userID int) error {
// 	query := "DELETE FROM user_info WHERE user_id=$1"
// 	_, err := db.Exec(query, userID)
// 	return err
// }

// // GetUser 根据用户ID获取用户信息
// func GetUser(db *sql.DB, userID int) (*User, error) {
// 	user := &User{}
// 	query := "SELECT user_id, user_name, password, email FROM user_info WHERE user_id=$1"
// 	err := db.QueryRow(query, userID).Scan(&user.User_id, &user.User_name, &user.Password, &user.Email)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

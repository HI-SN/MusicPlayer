package routers

import (
	"backend/controllers"
	"backend/middleware"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// 实例化控制器和服务
	songController := &controllers.SongController{
		SongService: &services.SongService{},
	}
	albumController := &controllers.AlbumController{
		Service: &services.AlbumService{},
	}
	playerController := &controllers.PlayerController{
		Service: &services.PlayerService{},
	}
	playlistController := &controllers.PlaylistController{
		Service: &services.PlaylistService{},
	}
	artistController := &controllers.ArtistController{
		ArtistSongService: &services.ArtistSongService{},
	}

	userController := &controllers.UserController{
		Service:     &services.UserService{},
		FService:    &services.FollowService{},
		MService:    &services.MomentService{},
		Aservice:    &services.ArtistService{},
		SetService:  &services.SettingService{},
		USService:   &services.UserSongService{},
		SongService: &services.SongService{},
		ABService:   &services.AlbumService{},
		// ASService:   &services.ArtistSongService{},
		PService:  &services.PlaylistService{},
		UPService: &services.UserPlaylistService{},
	}
	emailController := &controllers.EmailController{}
	momentController := &controllers.MomentController{
		Service:  &services.MomentService{},
		CService: &services.CommentService{},
		LService: &services.LikeService{},
	}
	commentController := &controllers.CommentController{
		Service:  &services.CommentService{},
		UService: &services.UserService{},
		LService: &services.LikeService{},
	}

	// r.Use(middleware.AuthMiddleware())

	// 用户认证相关路由
	authGroup := r.Group("/v1")
	{
		authGroup.POST("/send-captcha", emailController.SendVerification)
		authGroup.POST("/register", userController.CreateUser)
		authGroup.POST("/login", userController.Login)
		authGroup.POST("/forgot-password", userController.ForgetPassword)
		// authGroup.POST("/change-password", userController.ChangePassword)
		// // 退出登录
		// authGroup.POST("/logout", userController.Logout)
	}

	// 需要身份验证的路由
	authRequiredGroup := r.Group("")
	authRequiredGroup.Use(middleware.AuthMiddleware())
	{
		// 用户相关
		authRequiredGroup.POST("/v1/change-password", userController.ChangePassword)
		authRequiredGroup.GET("/user/basic", userController.GetUserBasic)
		authRequiredGroup.PUT("/user/basic", userController.UpdateUser)
		authRequiredGroup.GET("/user/setting", userController.GetUserSetting)
		authRequiredGroup.PUT("/user/setting", userController.UpdateUserSetting)
		authRequiredGroup.POST("/user/follow/:user_id", userController.FollowUser)
		authRequiredGroup.DELETE("/user/unfollow/:user_id", userController.UnfollowUser)
		authRequiredGroup.POST("/user/follow/artist/:artist_id", userController.FollowArtist)
		authRequiredGroup.DELETE("/user/unfollow/artist/:artist_id", userController.UnfollowArtist)
		authRequiredGroup.POST("/user/playlist/create", userController.CreatePlaylist)
		authRequiredGroup.DELETE("/user/playlist/:playlist_id", userController.DeletePlaylist)
		authRequiredGroup.POST("/user/like/playlist/:playlist_id", userController.LikePlaylist)
		authRequiredGroup.DELETE("/user/unlike/playlist/:playlist_id", userController.UnlikePlaylist)
		authRequiredGroup.POST("/user/like/song/:song_id", userController.LikeSong)
		authRequiredGroup.DELETE("/user/unlike/song/:song_id", userController.UnlikeSong)
		authRequiredGroup.GET("/user/:user_id/followers", userController.GetFollowers)
		authRequiredGroup.GET("/user/:user_id/following", userController.GetFollowing)
		authRequiredGroup.GET("/user/:user_id/artist", userController.GetUserArtist)
		// 动态相关
		authRequiredGroup.POST("/comment/moment/:moment_id", commentController.CreateMomentComment)
		authRequiredGroup.POST("/moment", momentController.CreateMoment)
		authRequiredGroup.DELETE("/comment/moment/:comment_id", commentController.DeleteMomentComment)
		authRequiredGroup.POST("/moment/:moment_id/like", momentController.LikeMoment)
		authRequiredGroup.POST("/moment/:moment_id/unlike", momentController.UnLikeMoment)
		//评论相关
		authRequiredGroup.POST("/comment/:comment_id/like", commentController.LikeComment)
		authRequiredGroup.POST("/comment/:comment_id/unlike", commentController.UnLikeComment)
		// 其他需要身份验证的路由
	}

	// 用户相关路由
	userGroup := r.Group("/user")
	{
		userGroup.GET("/:user_id", userController.GetUser)
		userGroup.GET("/:user_id/like/song", userController.GetUserLikeSong)
		userGroup.GET("/:user_id/playlist", userController.GetUserPlaylist)
		userGroup.GET("/:user_id/like/playlist", userController.GetUserLikePlaylist)
	}

	// 动态相关路由
	momentGroup := r.Group("/moment")
	{
		momentGroup.GET("/all/:user_id", momentController.GetAllMoments)
		momentGroup.GET("/:moment_id/like/count", momentController.GetMomentLikeCount)
		momentGroup.PUT("/:moment_id", momentController.UpdateMoment)
		momentGroup.DELETE("/:moment_id", momentController.DeleteMoment)
	}

	// 评论相关路由
	commentGroup := r.Group("/comment")
	{
		commentGroup.GET("/moment/all/:moment_id", commentController.GetMomentComment)
		commentGroup.GET("/:comment_id/like/count", commentController.GetCommentLikeCount)
	}

	// 其他路由...
	// 歌曲相关路由
	songGroup := r.Group("/songs")
	{
		songGroup.POST("/create", songController.CreateSong)
		songGroup.GET("/:song_id", songController.GetSongByID)
		songGroup.PUT("/update/:song_id", songController.UpdateSongInfo)
		songGroup.POST("/:song_id/upload/lyrics", songController.UploadLyricsBySongID)
		songGroup.GET("/:song_id/download/audio", songController.DownloadAudioBySongID)
		songGroup.GET("/:song_id/download/lyrics", songController.DownloadLyricsBySongID)
		songGroup.DELETE("/delete/:song_id", songController.DeleteSongByID)
		songGroup.GET("/:song_id/comments", songController.GetCommentsBySongID)
		songGroup.GET("/:song_id/:playlist_id", controllers.AddSongToPlaylist)
	}
	r.GET("/res/songs/:search", songController.GetSongsBySearch)

	// 专辑相关路由
	albumGroup := r.Group("/albums")
	{
		albumGroup.POST("/create", albumController.CreateAlbum)
		albumGroup.GET("/:album_id", albumController.GetAlbum)
		albumGroup.PUT("/update/:album_id", albumController.UpdateAlbum)
		albumGroup.DELETE("/delete/:album_id", albumController.DeleteAlbum)
		albumGroup.POST("/:album_id/cover", albumController.UploadAlbumCover)
	}

	// 艺术家相关路由
	artistGroup := r.Group("/artists")
	{
		artistGroup.POST("/add-to-song", artistController.AddArtistToSong)
		artistGroup.GET("/:artist_id/songs", artistController.GetSongsByArtistID)
		artistGroup.GET("/detail/:id", controllers.GetArtistDetailByID)
	}
	r.GET("/res/singer/:keyword", artistController.GetArtistsBySearch)
	// 播放器相关路由
	playerGroup := r.Group("/player")
	{
		playerGroup.GET("/play/:song_id", playerController.PlaySong)
		playerGroup.GET("/pause/:song_id", playerController.PauseSong)
		playerGroup.GET("/resume/:song_id", playerController.ResumeSong)
		playerGroup.GET("/volume/:song_id/:volume", playerController.AdjustVolume)
		playerGroup.GET("/lyric/:song_id", playerController.ShowLyrics)
	}

	// 播放列表相关路由
	playlistGroup := r.Group("/playlists")
	{
		playlistGroup.POST("/", playlistController.CreatePlaylist)
		playlistGroup.GET("/:playlist_id", playlistController.GetPlaylist)
		playlistGroup.PUT("/update/:playlist_id", playlistController.UpdatePlaylist)
		playlistGroup.DELETE("/:playlist_id", playlistController.DeletePlaylist)
		playlistGroup.POST("/:playlist_id/add/:song_id", playlistController.AddSongToPlaylist)
		playlistGroup.DELETE("/:playlist_id/remove/:song_id", playlistController.RemoveSongFromPlaylist)
		playlistGroup.GET("/:playlist_id/songs", playlistController.GetSongsByPlaylistID)
		playlistGroup.POST("/:playlist_id/updatecover", playlistController.UploadPlaylistCover)
		playlistGroup.GET("/recommend/:type", playlistController.GetPlaylistsByType)
		playlistGroup.GET("/allsongs/:id", playlistController.GetSongIDsByPlaylistID)
	}
	r.GET("/res/playlist/:keyword", playlistController.GetPlaylistsBySearch)

	// 首页相关路由注册
	homeGroup := r.Group("/home")
	{
		homeGroup.GET("/playlist", controllers.GetHomePlaylists)
		homeGroup.GET("/ranking/:name", controllers.GetHomeRanking)
	}

	// 消息中心相关路由注册
	messageGroup := r.Group("/message")
	{
		messageGroup.GET("/records/:sender_id/:receiver_id", controllers.GetCurrentChatMessages)
		messageGroup.GET("/comment/:user_id", controllers.GetMomentComments)
		messageGroup.GET("/private/:user_id", controllers.GetPrivateMessageList)
	}

	// 排行榜页面相关路由注册
	rankingGroup := r.Group("/ranking")
	{
		rankingGroup.GET("/:name", controllers.GetRankDetailsByName)
	}
	r.GET("/search/:str", controllers.Search)
}

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
		Service: &services.SongService{},
	}
	albumController := &controllers.AlbumController{
		Service: &services.AlbumService{},
	}
	playerController := &controllers.PlayerController{
		Service: &services.PlayerService{},
	}
	uploadController := &controllers.UploadController{
		Service: &services.UploadService{},
	}
	playlistController := &controllers.PlaylistController{
		Service: &services.PlaylistService{},
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
		ASService:   &services.ArtistSongService{},
		PService:    &services.PlaylistService{},
		UPService:   &services.UserPlaylistService{},
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
	authGroup := r.Group("/api/v1")
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
	authRequiredGroup := r.Group("/api")
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
	userGroup := r.Group("/api/user")
	{
		userGroup.GET("/:user_id", userController.GetUser)
		userGroup.GET("/:user_id/following", userController.GetFollowing)
		userGroup.GET("/:user_id/followers", userController.GetFollowers)
		userGroup.GET("/:user_id/artist", userController.GetUserArtist)
		userGroup.GET("/:user_id/like/song", userController.GetUserLikeSong)
		userGroup.GET("/:user_id/playlist", userController.GetUserPlaylist)
		userGroup.GET("/:user_id/like/playlist", userController.GetUserLikePlaylist)
	}

	// 动态相关路由
	momentGroup := r.Group("/api/moment")
	{
		momentGroup.GET("/all/:user_id", momentController.GetAllMoments)
		momentGroup.GET("/:moment_id/like/count", momentController.GetMomentLikeCount)
		momentGroup.PUT("/:moment_id", momentController.UpdateMoment)
		momentGroup.DELETE("/:moment_id", momentController.DeleteMoment)
	}

	// 评论相关路由
	commentGroup := r.Group("/api/comment")
	{
		commentGroup.GET("/moment/all/:moment_id", commentController.GetMomentComment)
		commentGroup.GET("/:comment_id/like/count", commentController.GetCommentLikeCount)
	}

	// 其他路由...
	// 歌曲相关路由
	songGroup := r.Group("/songs")
	{
		songGroup.POST("/create", songController.CreateSong)
		songGroup.GET("/:song_id", songController.GetSong)
	}

	// 专辑相关路由
	albumGroup := r.Group("/albums")
	{
		albumGroup.POST("/create", albumController.CreateAlbum)
		albumGroup.GET("/:album_id", albumController.GetAlbum)
	}

	// 播放器相关路由
	playerGroup := r.Group("/player")
	{
		playerGroup.GET("/play/:song_id", playerController.PlaySong)
		playerGroup.GET("/pause/:song_id", playerController.PauseSong)
		playerGroup.GET("/resume/:song_id", playerController.ResumeSong)
		playerGroup.GET("/volume/:song_id/:volume", playerController.AdjustVolume)
	}

	// 上传相关路由
	uploadGroup := r.Group("/upload")
	{
		uploadGroup.POST("/audio", uploadController.UploadAudio)
		uploadGroup.POST("/song-info", uploadController.UploadSongInfo)
		uploadGroup.POST("/lyrics/:song_id", uploadController.UploadLyrics)
	}
	// 播放列表相关路由
	playlistGroup := r.Group("/playlists")
	{
		playlistGroup.POST("/", playlistController.CreatePlaylist)
		playlistGroup.GET("/:playlist_id", playlistController.GetPlaylist)
		playlistGroup.PUT("/:playlist_id", playlistController.UpdatePlaylist)
		playlistGroup.DELETE("/:playlist_id", playlistController.DeletePlaylist)
		playlistGroup.POST("/:playlist_id/add/:song_id", playlistController.AddSongToPlaylist)
		playlistGroup.DELETE("/:playlist_id/remove/:song_id", playlistController.RemoveSongFromPlaylist)
		playlistGroup.GET("/:playlist_id/songs", playlistController.GetSongsByPlaylistID)
	}
}

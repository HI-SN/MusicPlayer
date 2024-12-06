package routers

import (
	"backend/controllers"
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
		Service:  &services.UserService{},
		FService: &services.FollowService{},
		MService: &services.MomentService{},
		Aservice: &services.ArtistService{},
	}
	emailController := &controllers.EmailController{}
	momentController := &controllers.MomentController{
		Service: &services.MomentService{},
	}

	// 用户认证相关路由
	authGroup := r.Group("/api/v1")
	{
		// 发送邮箱验证码
		authGroup.POST("/send-captcha", emailController.SendVerification)
		// 注册
		authGroup.POST("/register", userController.CreateUser)
		// 登录
		authGroup.POST("/login", userController.Login)
		// 找回密码
		authGroup.POST("/forgot-password", userController.ForgetPassword)
		// 已登录用户修改密码
		authGroup.POST("/change-password", userController.ChangePassword)
		// // 退出登录
		// authGroup.POST("/logout", userController.Logout)
	}

	// 用户相关路由
	userGroup := r.Group("/api/user")
	{
		userGroup.GET("/:user_id", userController.GetUser)
		userGroup.PUT("/:user_id", userController.UpdateUser)
		userGroup.GET("/:user_id/following", userController.GetFollowing)
		userGroup.GET("/:user_id/followers", userController.GetFollowers)
	}

	// 动态相关路由
	momentGroup := r.Group("/moment")
	{
		momentGroup.POST("/", momentController.CreateMoment)
		momentGroup.GET("/:moment_id", momentController.GetMoment)
		momentGroup.GET("/all/:user_id", momentController.GetAllMoments)
		momentGroup.PUT("/:moment_id", momentController.UpdateMoment)
		momentGroup.DELETE("/:moment_id", momentController.DeleteMoment)
	}

	// productGroup := r.Group("/products")
	// {
	// 	productGroup.POST("/", productController.CreateProduct)
	// 	productGroup.GET("/:id", productController.GetProduct)
	// 	productGroup.PUT("/:id", productController.UpdateProduct)
	// 	productGroup.DELETE("/:id", productController.DeleteProduct)
	// }

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

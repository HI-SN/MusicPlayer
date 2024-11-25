package routers

import (
	"backend/controllers" // 替换为您的项目路径

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController *controllers.UserController,
	emailController *controllers.EmailController, momentController *controllers.MomentController) {

	// 用户认证相关路由

	// 发送邮箱验证码
	r.POST("/captcha/send", emailController.SendVerification)
	// 验证邮箱验证码
	r.POST("/captcha/verify", emailController.VerifyCode)
	// 注册
	r.POST("/register", userController.CreateUser)
	// 登录
	r.POST("/login", userController.Login)
	// 退出登录
	r.POST("/logout", userController.Logout)

	// 用户相关路由
	userGroup := r.Group("/users")
	{
		userGroup.GET("/:user_id", userController.GetUser)
		userGroup.PUT("/:user_id", userController.UpdateUser)
		userGroup.POST("/:user_id/moments", userController.CreateMoment)
		userGroup.GET("/:user_id/moments", userController.GetAllMoments)
		// userGroup.DELETE("/:user_id/:moment_id", userController.DeleteMoment)
		// userGroup.DELETE("/:user_id", userController.DeleteUser)
		userGroup.GET("/:user_id/follows", userController.GetFollows)
		userGroup.GET("/:user_id/followers", userController.GetFollowers)
	}

	// 动态相关路由
	momentGroup := r.Group("/moments")
	{
		momentGroup.GET("/:moment_id", momentController.GetMoment)

	}

	// productGroup := r.Group("/products")
	// {
	// 	productGroup.POST("/", productController.CreateProduct)
	// 	productGroup.GET("/:id", productController.GetProduct)
	// 	productGroup.PUT("/:id", productController.UpdateProduct)
	// 	productGroup.DELETE("/:id", productController.DeleteProduct)
	// }

	// 其他路由...
}

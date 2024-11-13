package routers

import (
	"backend/controllers" // 替换为您的项目路径

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController *controllers.UserController,
	emailController *controllers.EmailController) {

	// 用户认证相关路由
	authGroup := r.Group("/auth")
	{
		// 发送邮箱验证码
		authGroup.POST("/send-verification-email", emailController.SendVerification)
		// 验证邮箱验证码
		authGroup.POST("/verify-email", emailController.VerifyCode)
		// 注册
		authGroup.POST("/register", userController.CreateUser)
		// 登录
		authGroup.POST("/login", userController.Login)
		// 退出登录
		authGroup.POST("/logout", userController.Logout)
	}

	// 用户相关路由
	userGroup := r.Group("/users")
	{
		userGroup.GET("/:user_id", userController.GetUser)
		userGroup.PUT("/:user_id", userController.UpdateUser)
		// userGroup.DELETE("/:user_id", userController.DeleteUser)
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

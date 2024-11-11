package routers

import (
	"backend/controllers" // 替换为您的项目路径

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController *controllers.UserController,
	emailController *controllers.EmailController) {

	// 发送邮箱验证码
	r.POST("/send-verification-email", emailController.SendVerification)

	// 验证邮箱验证码
	r.POST("/verify-email", emailController.VerifyCode)

	// 用户相关路由
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", userController.CreateUser)
		userGroup.POST("/login", userController.Login)
		// userGroup.GET("/:user_id", userController.GetUser)
		// userGroup.PUT("/:user_id", userController.UpdateUser)
		// userGroup.DELETE("/:id", userController.DeleteUser)
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

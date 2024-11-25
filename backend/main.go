package main

import (
	"backend/configs"
	"backend/controllers"
	"backend/database"
	"backend/routers"

	// "fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 初始化配置
	configs.LoadConfig("configs/config.json")

	// 用于保存临时验证码的专用数据库
	database.Redis()

	// 获取初始化的数据库
	database.InitDB()
	// 延迟关闭数据库
	// defer database.CloseDB()

	//创建一个默认的路由引擎
	r := gin.Default()

	// 实例化控制器，并注册到路由中
	userController := controllers.UserController{}
	emailController := controllers.EmailController{}
	momentController := controllers.MomentController{}

	routers.SetupRoutes(r, &userController, &emailController, &momentController)

	//在9090端口启动服务
	panic(r.Run(":8080"))

	// c := &models.Comment{}

	// c.User_id = "e9nRUN7ZRB6pDw6"
	// c.Content = "zxcasd   sdd"
	// c.Type = "song"
	// c.Target_id = 3

	// if c1, err := models.GetComment(database.DB, 5); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(c1)
	// }

	// if err := models.CreateComment(database.DB, c); err != nil {
	// 	fmt.Println(err)
	// }

	// mID := 4
	// if err := models.DeleteMoment(database.DB, mID); err != nil {
	// 	fmt.Println(err)
	// }

	// if results, err := models.GetAllComments(database.DB, 6, "moment"); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	for _, r := range results {
	// 		fmt.Println(r)
	// 	}
	// }
}

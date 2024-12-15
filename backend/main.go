package main

import (
	"backend/configs"
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

	// Aservice := &services.ArtistService{}
	// artist, err := Aservice.GetArtist(8)
	// if err != nil {
	// fmt.Println("err:", err)
	// } else {
	// fmt.Println(artist)
	// }

	// 延迟关闭数据库
	// defer database.CloseDB()

	//创建一个默认的路由引擎
	r := gin.Default()

	routers.SetupRoutes(r)

	//在9090端口启动服务
	panic(r.Run(":54212"))
}

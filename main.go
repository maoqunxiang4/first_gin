package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goPro/config"
)

func main() {
	//初始化配置,凡是使用了yml文件注入的地方都要使用InitConfig文件配置的方法。
	//注意：在一个方法中调用了InitCnnfig，一个方法A调用另外一个方法B，这个方法B就不需要再调用InitConfig了
	config.InitConfig()

	r := gin.Default()
	r = collectRoutes(r)
	port := viper.GetString("server.port")
	//如果有用户自定义的port就用用户自定义的port
	if port != "" {
		panic(r.Run(":" + port))
	}
	//否则就用默认的port
	r.Run(":8080")
}

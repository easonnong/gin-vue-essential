package main

import (
	"os"

	"github.com/easonnong/gin-vue-essential/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	//读取配置信息
	InitConfig()

	//连接mysql
	common.InitDB()

	db := common.GetDB()
	defer db.Close()

	router := gin.Default()
	router = CollectRoutes(router)

	//从配置文件获取监听端口
	portListen := viper.GetString("server.port")
	if portListen == "" {
		panic("监听端口错误")
	}
	portAddr := viper.GetString("server.host") + ":" + portListen

	//运行路由
	panic(router.Run(portAddr))
}

//从配置文件中读取
func InitConfig() {
	//获取当前工作目录
	workDir, err := os.Getwd()
	if err != nil {
		panic("获取当前工作目录失败")
	}
	//设置读取文件的类型
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err = viper.ReadInConfig()
	if err != nil {
		panic("读取文件失败")
	}
}

package main

import (
	"github.com/easonnong/gin-vue-essential/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//连接mysql
	common.InitDB()

	db := common.GetDB()
	defer db.Close()

	router := gin.Default()
	router = CollectRoutes(router)

	//运行路由
	router.Run()
}

package common

import (
	"fmt"

	"github.com/easonnong/gin-vue-essential/model"
	"github.com/jinzhu/gorm"
)

//数据库
var DB *gorm.DB

//初始化开启连接池
func InitDB() {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "gin-vue-essential"
	username := "root"
	password := "2222"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	dbTemp, err := gorm.Open(driverName, args)
	DB = dbTemp
	if err != nil {
		panic("failed to connect database, err=" + err.Error())
	}

	//创建数据表
	DB.AutoMigrate(&model.User{})

}

//获取DB
func GetDB() *gorm.DB {
	return DB
}

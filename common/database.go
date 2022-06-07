package common

import (
	"fmt"
	"net/url"

	"github.com/easonnong/gin-vue-essential/model"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

//数据库
var DB *gorm.DB

//初始化开启连接池
func InitDB() {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
	)

	dbTemp, err := gorm.Open(driverName, args)
	DB = dbTemp
	if err != nil {
		panic("failed to connect database, err=" + err.Error())
	}

	//创建数据表 自动迁移
	DB.AutoMigrate(&model.User{})

}

//获取DB
func GetDB() *gorm.DB {
	return DB
}

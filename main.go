package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//定义gorm model
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

func main() {
	//连接mysql
	db := InitDB()
	defer db.Close()

	router := gin.Default()

	router.POST("/api/auth/register", func(ctx *gin.Context) {
		//获取参数(name, telephone, password)
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		//数据验证
		//如果telephone长度不为11位
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "手机号必须为11位",
			})
			return
		}

		//如果password长度小于6位
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "密码长度不能少于6位",
			})
			return
		}

		//如果name为空，给一个10位的随机字符串
		if len(name) == 0 {
			name = randomString(10)
		}

		fmt.Printf("name=%v tel=%v pwd=%v", name, telephone, password)

		//判断手机号是否存在
		if isTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "该用户已经存在",
			})
			return
		}

		//创建用户(用户不存在时)
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		//返回结果
		ctx.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})

	//运行路由
	router.Run()
}

//随机生成10位字符串
func randomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	//随机数种子
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

//判断手机号是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}

//开启连接池
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "gin_vue_essential"
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

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err=" + err.Error())
	}

	//创建数据表
	db.AutoMigrate(&User{})

	return db
}

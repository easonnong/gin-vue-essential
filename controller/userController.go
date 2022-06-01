package controller

import (
	"log"
	"net/http"

	"github.com/easonnong/gin-vue-essential/common"
	"github.com/easonnong/gin-vue-essential/model"
	"github.com/easonnong/gin-vue-essential/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(ctx *gin.Context) {
	//获取DB
	db := common.GetDB()

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
		name = util.RandomString(10)
	}

	log.Printf("name=%v tel=%v pwd=%v", name, telephone, password)

	//判断手机号是否存在
	if isTelephoneExist(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "该用户已经存在",
		})
		return
	}

	//创建用户(用户不存在时)
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	common.DB.Create(&newUser)

	//返回结果
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})

}

//判断手机号是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}

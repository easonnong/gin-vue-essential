package controller

import (
	"log"
	"net/http"

	"github.com/easonnong/gin-vue-essential/common"
	"github.com/easonnong/gin-vue-essential/dto"
	"github.com/easonnong/gin-vue-essential/model"
	"github.com/easonnong/gin-vue-essential/response"
	"github.com/easonnong/gin-vue-essential/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//用户注册
func Register(ctx *gin.Context) {
	//获取DB
	db := common.GetDB()

	//获取参数(name, telephone, password)
	/*方法一：
	var requestMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&requestMap)
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//方法二：
	var requestUser = model.User{}
	json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password") */
	//方法三：
	var requestUser = model.User{}
	ctx.Bind(&requestUser)
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	//数据验证
	//如果telephone长度不为11位
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	//如果password长度小于6位
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码长度不能少于6位")
		return
	}

	//如果name为空，给一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Printf("name=%v tel=%v pwd=%v", name, telephone, password)

	//判断手机号是否存在
	if isTelephoneExist(db, telephone) {

		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "该用户已经存在")
		return
	}

	//创建用户(用户不存在时)
	//加密用户的密码
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {

		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密用户密码失败")
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	common.DB.Create(&newUser)

	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {

		response.Response(
			ctx,
			http.StatusInternalServerError,
			500,
			nil,
			"系统异常",
		)

		//打日志
		log.Printf("token generate error = %v\n", err)
		return
	}

	//返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "注册成功",
	})

}

//用户登录
func Login(ctx *gin.Context) {
	//获取DB
	db := common.GetDB()

	//获取参数(tel,pwd)
	var requestUser = model.User{}
	ctx.Bind(&requestUser)
	telephone := requestUser.Telephone
	password := requestUser.Password

	//数据验证
	//如果telephone长度不为11位
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	//如果password长度小于6位
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码长度不能少于6位")
		return
	}

	//判断手机号是否存在
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "该用户不存在")
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {

		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")

		//打日志
		log.Printf("token generate error = %v\n", err)
		return
	}

	//返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功",
	})
}

//用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	response.Response(
		ctx, http.StatusOK, 200,
		gin.H{
			"user": dto.ToUserDto(user.(model.User)),
		},
		"密码错误",
	)

}

//判断手机号是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}

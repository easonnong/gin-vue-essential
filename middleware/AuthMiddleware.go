package middleware

import (
	"net/http"
	"strings"

	"github.com/easonnong/gin-vue-essential/common"
	"github.com/easonnong/gin-vue-essential/model"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足，验证token失败",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		//解析token
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足,解析token失败",
			})
			ctx.Abort()
			return
		}

		//验证通过后获取claim中的UserId
		userId := claims.UserId
		db := common.GetDB()
		var user model.User
		db.First(&user, userId)

		//如果用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "用户不存在",
			})
			ctx.Abort()
			return
		}

		//如果用户存在,将user的信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}

package middleware

import (
	"fmt"

	"github.com/easonnong/gin-vue-essential/response"
	"github.com/gin-gonic/gin"
)

//自定义recover获取panic信息

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(ctx, fmt.Sprint(err), nil)
			}
		}()

		ctx.Next()
	}
}

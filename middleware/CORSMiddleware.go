package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//可访问域名
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//缓存时间
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		//可通过访问的方法
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		//指定Header
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		//判断请求的方法是否为option
		if ctx.Request.Method == http.MethodOptions {
			//直接返回200
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}

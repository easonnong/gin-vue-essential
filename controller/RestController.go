package controller

import "github.com/gin-gonic/gin"

//定义增删改查的接口
type IRestController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Show(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type RestController struct {
}

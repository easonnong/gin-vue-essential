package main

import (
	"github.com/easonnong/gin-vue-essential/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(router *gin.Engine) *gin.Engine {
	router.POST("/api/auth/register", controller.Register)
	return router
}

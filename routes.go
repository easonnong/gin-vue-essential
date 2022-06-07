package main

import (
	"github.com/easonnong/gin-vue-essential/controller"
	"github.com/easonnong/gin-vue-essential/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(router *gin.Engine) *gin.Engine {
	//解决同源策略,自定义panic处理
	router.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	//用户注册
	router.POST("/api/auth/register", controller.Register)
	//用户登录
	router.POST("/api/auth/login", controller.Login)
	//用户信息
	//中间件用于保护用户信息
	router.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	categoryRoutes := router.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT("/:id", categoryController.Update)
	categoryRoutes.GET("/:id", categoryController.Show)
	categoryRoutes.DELETE("/:id", categoryController.Delete)

	return router
}

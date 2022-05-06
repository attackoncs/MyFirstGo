package main

import (
	"github.com/gin-gonic/gin"
	"yl/ginessential/controller"
	"yl/ginessential/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info) //中间件保护路由并用户认证
	return r
}

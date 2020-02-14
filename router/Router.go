package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"winter_holiday/controller"
)

func Router() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("loginuser"))
	router.Use(sessions.Sessions("session", store))


	//注册
	router.POST("/register",controller.Regist)
	//登录
	router.POST("/login",controller.Login)
	//永久注销账号

	//网页首页
	router.GET("/", controller.FirstPage)

	//注销登录
	router.GET("/exit", controller.Exit)

	return router
}



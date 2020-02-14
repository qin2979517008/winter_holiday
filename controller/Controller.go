package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"winter_holiday/data"
	"winter_holiday/model"
)

func Regist(c *gin.Context) {
	var user data.User
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	if model.SearchbyName(user.Name) {
		if model.InserUser(&user) {
			c.JSON(http.StatusOK, gin.H{"code": 200, "message": "注册成功"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 200, "message": "注册失败"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "该用户已经纯在，注册失败"})
	}
}

func Login(c *gin.Context) {
	var user data.User
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	if model.SearchUser(&user) {
		session := sessions.Default(c)
		session.Set("loginuser", user.Name)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "登录失败,用户名或者密码错误"})
	}
}

func Exit(c *gin.Context) {
	//清除该用户登录状态的数据
	session := sessions.Default(c)
	session.Delete("loginuser")
	session.Save()
	fmt.Println("delete session...", session.Get("loginuser"))
	c.Redirect(http.StatusMovedPermanently, "/")
}

//获取session，查看用户是否登录
func Session(c *gin.Context) bool {
	session := sessions.Default(c)
	loginuser := session.Get("loginuser")
	if loginuser != nil {
		return true
	} else {
		return false
	}
}

func FirstPage(c *gin.Context) {
	islogin := Session(c)
	c.JSON(http.StatusOK, gin.H{"IsLogin": islogin})
}

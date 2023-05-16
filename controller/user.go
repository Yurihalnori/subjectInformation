package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"subjectInformation/model"
	"subjectInformation/service"
)

type UserController struct {
}

func (UserController) Logincheck(c *gin.Context) {
	session := sessions.Default(c)
	num := session.Get("id")
	if num != nil {
		c.Next()
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "请登录"})
	}
}

func (UserController) CreateUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	isExist := service.CheckUsername(user.Username)
	if isExist == false {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户已存在",
			"code":    service.OpErr,
		})
		return
	}
	err := service.CreateAUser(&user)
	session := sessions.Default(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
			"code":    service.OpErr,
		})
	} else {
		session.Set("id", user.Id)
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"username": user.Username,
				"password": user.Password,
				"id":       user.Id,
			},
		})
	}
}

func (UserController) Login(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	pass, err := service.CheckPassword(user.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
			"code":    service.ParamErr,
		})
		return
	}
	if pass.Password == user.Password {
		session := sessions.Default(c)
		session.Set("id", user.Id)
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    pass,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户名或密码错误",
			"code":    service.ParamErr,
		})
	}
}

func (UserController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	num := session.Get("id")
	if num != nil {
		session.Clear()
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"login": true,
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "您还未登录",
			"code":    service.AuthErr,
		})
	}
}

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
	num := session.Get("number")
	if num != nil {
		c.Next()
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "请登录"})
	}
}

func (UserController) CreateUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	err := service.CreateAUser(&user)
	session := sessions.Default(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
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
	pass := service.CheckPassword(user.Username)
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
			"error":   "密码错误",
		})
	}
}

func (UserController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	num := session.Get("number")
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
			"success": true,
			"data": gin.H{
				"login": false,
			},
		})
	}
}

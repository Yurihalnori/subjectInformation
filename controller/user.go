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

func (UserController) LoginCheck(c *gin.Context) {
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
	bindErr := c.BindJSON(&user)
	if bindErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": bindErr,
			"code":    service.OpErr,
		})
		return
	}
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
		sessionErr := session.Save()
		if sessionErr != nil {
			return
		}
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
	bindErr := c.BindJSON(&user)
	if bindErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": bindErr,
			"code":    service.OpErr,
		})
		return
	}
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
		sessionErr := session.Save()
		if sessionErr != nil {
			return
		}
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
		err := session.Save()
		if err != nil {
			return
		}
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

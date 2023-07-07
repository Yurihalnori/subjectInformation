package controller

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		c.Abort()
	}
}

func (UserController) AdminCheck(c *gin.Context) {
	session := sessions.Default(c)
	num := session.Get("id")
	user, err := service.UserService{}.CheckInfo(num.(string))
	if err != nil {
		_ = c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	}
	if user.Usertype == 1 {
		c.Next()
	} else {
		c.Abort()
		c.JSON(http.StatusOK, gin.H{"error": "您不是管理员"})
	}
}

func (UserController) CreateUser(c *gin.Context) {
	var user model.User
	bindErr := c.BindJSON(&user)
	if bindErr != nil {
		_ = c.Error(&gin.Error{
			Err:  bindErr,
			Type: service.ParamErr,
		})
		return
	}
	isExist := service.UserService{}.CheckUsername(user.Username)
	if isExist == false {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户名已存在",
			"code":    service.OpErr,
		})
		return
	}
	err := service.UserService{}.CreateAUser(&user)
	session := sessions.Default(c)
	if err != nil {
		_ = c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
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
			"message": "用户名或密码错误",
			"code":    service.ParamErr,
		})
		return
	}
	res, err := service.UserService{}.CheckPassword(user.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
			"code":    service.ParamErr,
		})
		return
	}
	if res.Password == user.Password {
		if user.Usertype == 1 {
			if res.Usertype != 1 {
				c.JSON(http.StatusOK, gin.H{
					"success": false,
					"message": "您不是管理员",
					"code":    service.AuthErr,
				})
				return
			}
		}
		session := sessions.Default(c)
		session.Set("id", res.Id)
		sessionErr := session.Save()
		if sessionErr != nil {
			_ = c.Error(&gin.Error{
				Err:  sessionErr,
				Type: service.ParamErr,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"userId": res.Id,
				"role":   res.Usertype,
			},
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

func (UserController) CheckUserStatus(c *gin.Context) {
	id := c.Param("id")
	validate := validator.New()
	if err := validate.Var(id, "numeric"); err != nil {
		_ = c.Error(&gin.Error{
			Err:  errors.New("请返回正确的id"),
			Type: service.ParamErr,
		})
		return
	}
	user, err := service.UserService{}.CheckInfo(id)
	if user == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户不存在",
			"code":    service.ParamErr,
		})
		return
	}
	if err != nil {
		_ = c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"userId": user.Id,
			"role":   user.Usertype,
		},
	})
}

func (UserController) ChangeUserInfo(c *gin.Context) {
	var form model.UserChangeRequest
	bindErr := c.ShouldBindJSON(&form)
	if bindErr != nil {
		_ = c.Error(&gin.Error{
			Err:  bindErr,
			Type: service.ParamErr,
		})
	}
	err := service.UserService{}.ChangeInfo(form)
	if err != nil {
		_ = c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"userId": form.Id,
			"role":   form.UserType,
		},
	})
}

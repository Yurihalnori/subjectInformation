package middleware

import (
	"net/http"
	"strconv"

	"subjectInformation/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginCheck(c *gin.Context) {
	session := sessions.Default(c)
	num := session.Get("id")
	if num != nil {
		c.Next()
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "请登录"})
		c.Abort()
	}
}

func AdminCheck(c *gin.Context) {
	session := sessions.Default(c)
	var num = session.Get("id")
	println(session.Get("id"))
	var stringId = strconv.Itoa(num.(int))
	user, err := service.UserService{}.CheckInfo(stringId)
	if err != nil {
		_ = c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		c.Abort()
		return
	}
	if user.Usertype == 1 {
		c.Next()
	} else {
		c.Abort()
		c.JSON(http.StatusOK, gin.H{"error": "您不是管理员"})
	}
}

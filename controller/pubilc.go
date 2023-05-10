package controller

import (
	"fmt"
	"net/http"
	"time"

	"subjectInformation/model"
	"subjectInformation/service"

	"github.com/gin-gonic/gin"
)

type PublicDatabaseController struct {
}

func (s PublicDatabaseController) Hello(c *gin.Context) {
	var form struct {
		Msg string `form:"msg" binding:"required,msg"`
	}
	if err := c.ShouldBindQuery(&form); err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	}

	helloService := service.HelloService{}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    helloService.Hello(form.Msg),
	})
}

func (s PublicDatabaseController) HelloTime(c *gin.Context) {
	var form struct {
		Date time.Time `form:"date" binding:"required,timing" time_format:"2006-01-02"`
	}
	if err := c.ShouldBindQuery(&form); err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	}

	helloService := service.HelloService{}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    helloService.HelloTime(form.Date),
	})
}

func (s PublicDatabaseController) Add(c *gin.Context) {
	var form model.Form
	if err := c.ShouldBind(&form); err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	}
	databaseService := service.DatabaseService{}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    databaseService.Add(form),
	})
}

package controller

import (
	"fmt"
	"net/http"
	"time"

	"subjectInformation/service"

	"github.com/gin-gonic/gin"
)

type PublicDatabaseController struct {
}

type msg struct {
	Module         string `form:"module" json:"module" binding:"binary"`
	Category       string `form:"category" json:"category" binding:"len=9"`
	Title          string `form:"title" json:"title" binding:"omitempty"`
	Text           string `form:"text" json:"text" binding:"omitempty"`
	Name           string `form:"name" json:"name" binding:"omitempty"`
	Classification string `form:"classification" json:"classification" binding:"omitempty"` //项目类别 学位点类别 导师类型
	Sponsor        string `form:"sponsor" json:"" binding:"omitempty"`
	ApprovalNumber string `form:"approvalNumber" json:"approvalNumber" biding:"omitempty"`
	CreateDate     string `form:"createDate" json:"createDate" biding:"omitempty"`
	Superintendent string `form:"superintendent" json:"superintendent" biding:"omitempty"`
	Organization   string `form:"organization" json:"organization" biding:"omitempty"`
	Nation         string `form:"nation" json:"nation" biding:"omitempty"`
	Province       string `form:"province" json:"province" biding:"omitempty"`
	City           string `form:"city" json:"city" biding:"omitempty"`
	University     string `form:"university" json:"university" biding:"omitempty"`
	College        string `form:"college" json:"college" biding:"omitempty"`
	Direction      string `form:"direction" json:"direction" biding:"omitempty"`
	Contact        string `form:"contact" json:"contact" biding:"omitempty"`
	Tutor          string `form:"tutor" json:"tutor" biding:"omitempty"`
	Year           string `form:"year" json:"year" biding:"omitempty"`
	Technique      string `form:"technique" json:"technique" biding:"omitempty"`
	KeyWord        string `form:"keyWord" json:"keyWord" biding:"omitempty"`
	Periodical     string `form:"periodical" json:"periodical" biding:"omitempty"`
	Digest         string `form:"digest" json:"digest" biding:"omitempty"`
	Data           string `form:"data" json:"data" biding:"omitempty"`
	Language       string `form:"language" json:"language" biding:"omitempty"`
	Publisher      string `form:"publisher" json:"publisher" biding:"omitempty"`
	Click          string `form:"click" json:"click" biding:"omitempty"`
	Download       string `form:"download" json:"download" biding:"omitempty"`
}

type form struct {
	Total int   `form:"total" json:"total" binding:"required"`
	List  []msg `form:"list" json:"list" binding:"required"`
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
	var data form
	if err := c.ShouldBind(&data); err != nil {
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
		Data:    databaseService.Add(data),
	})
}

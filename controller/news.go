package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"subjectInformation/model"
	"subjectInformation/service"
)

type NewsController struct {
}

func (NewsController) AddNews(c *gin.Context) {
	var news model.News

	info, err := service.AddOneNews(news)
	if err != nil {
		fmt.Printf("controller %v", err)
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err,
			"code":    service.ParamErr,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    info,
	})
}

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
	var form model.NewsForm
	bindErr := c.ShouldBind(&form)
	if bindErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": bindErr,
			"code":    service.OpErr,
		})
		return
	}
	allNews := form.NewsList
	fmt.Println(allNews)
	var infoList []interface{}
	var errList []interface{}
	for _, aNews := range allNews {
		info, err := service.AddOneNews(aNews)
		if err != nil {
			errList = append(errList, err)
			fmt.Printf("controller %v", err)
		} else {
			infoList = append(infoList, info)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"Add":  infoList,
			"fail": errList},
	})
}

package controller

import (
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
	var infoList []interface{}
	var errList []interface{}
	for _, aNews := range allNews {
		info, err := service.AddOneNews(aNews)
		if err != nil {
			errList = append(errList, err)
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

func (NewsController) GetNews(c *gin.Context) {
	var form model.GetSomeNews
	bindErr := c.ShouldBindQuery(&form)
	if bindErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": bindErr,
			"code":    service.OpErr,
		})
		return
	}
	res, err := service.GetSomeNews(form)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err,
			"code":    service.ParamErr,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    res,
		})
	}
}

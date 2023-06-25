package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	res, total, err := service.GetSomeNews(form)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err,
			"code":    service.ParamErr,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"category": form.Category,
				"module":   form.Module,
				"total":    total,
				"list":     res,
			},
		})
	}
}

func (NewsController) ApplyOneNews(c *gin.Context) {
	id := c.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err,
			"code":    service.ParamErr,
		})
		return
	}
	IntId, _ := strconv.Atoi(id)
	aPieceOfNews, FindErr := service.GetOneNew(IntId)
	if FindErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": FindErr,
			"code":    service.ParamErr,
		})
	} else {
		err := service.ClickNewsOnce(IntId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": err,
				"code":    service.ParamErr,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"data":    aPieceOfNews,
			})
		}
	}
}

func (NewsController) EditOneNews(c *gin.Context) {
	var form model.News
	bindErr := c.ShouldBindJSON(&form)
	if bindErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": bindErr,
			"code":    service.OpErr,
		})
		return
	}
	id := c.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err,
			"code":    service.ParamErr,
		})
		return
	}
	form.Id, _ = strconv.Atoi(id)
	err := service.EditNews(form)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err,
			"code":    service.ParamErr,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    form.Id,
		})
	}
}

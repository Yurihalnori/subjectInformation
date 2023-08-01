package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"subjectInformation/model"
	"subjectInformation/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type NewsController struct {
}

func (NewsController) AddNews(c *gin.Context) {
	var form model.NewsForm
	bindErr := c.ShouldBind(&form)
	if bindErr != nil {
		_ = c.Error(&gin.Error{
			Err:  bindErr,
			Type: service.ParamErr,
		})
		return
	}
	allNews := form.NewsList
	var infoList []interface{}
	var errList []interface{}
	for _, aNews := range allNews {
		info, err := service.NewsService{}.AddOneNews(aNews)
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
			"fail": errList,
		},
	})
}

func (NewsController) GetNews(c *gin.Context) {
	var form model.GetSomeNews
	bindErr := c.ShouldBindQuery(&form)
	if bindErr != nil {
		_ = c.Error(&gin.Error{
			Err:  bindErr,
			Type: service.ParamErr,
		})
		return
	}
	res, total, err := service.NewsService{}.GetSomeNews(form)
	fmt.Println(res)
	if err != nil {
		_ = c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
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
	validate := validator.New()
	if err := validate.Var(id, "numeric"); err != nil {
		_ = c.Error(&gin.Error{
			Err:  errors.New("请返回正确的id"),
			Type: service.ParamErr,
		})
		return
	}
	if _, err := strconv.Atoi(id); err != nil {
		_ = c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	}
	IntId, _ := strconv.Atoi(id)
	aPieceOfNews, FindErr := service.NewsService{}.GetOneNew(IntId)
	if FindErr != nil {
		_ = c.Error(&gin.Error{
			Err:  errors.New("请返回正确的id"),
			Type: service.ParamErr,
		})
	} else {
		err := service.NewsService{}.ClickNewsOnce(IntId)
		if err != nil {
			_ = c.Error(&gin.Error{
				Err:  err,
				Type: service.ParamErr,
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
		_ = c.Error(&gin.Error{
			Err:  bindErr,
			Type: service.ParamErr,
		})
		return
	}
	id := c.Param("id")
	validate := validator.New()
	if err := validate.Var(id, "numeric"); err != nil {
		_ = c.Error(&gin.Error{
			Err:  errors.New("请返回正确的id"),
			Type: service.ParamErr,
		})
		return
	}
	form.Id, _ = strconv.Atoi(id)
	err := service.NewsService{}.EditNews(form)
	if err != nil {
		_ = c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"Id": form.Id,
			},
		})
	}
}

func (NewsController) DeleteOneNews(c *gin.Context) {
	id := c.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		_ = c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	}
	IntId, _ := strconv.Atoi(id)
	err := service.NewsService{}.DeleteNews(IntId)
	if err != nil {
		_ = c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"id": IntId,
			},
		})
	}
}

func (NewsController) SearchNews(c *gin.Context) {
	var form model.NewsSearchRequest
	bindErr := c.ShouldBindJSON(&form)
	if bindErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": bindErr,
			"code":    service.OpErr,
		})
		return
	}
	total, res, err := service.NewsService{}.SearchNews(form)
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
				"total":    total,
				"module":   form.Module,
				"category": form.Category,
				"list":     res,
			},
		})
	}
}

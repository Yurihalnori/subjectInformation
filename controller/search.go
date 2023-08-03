package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"subjectInformation/model"
	"subjectInformation/service"
)

type SearchController struct {
}

func (SearchController) SearchCommonDB(c *gin.Context) {
	var form model.SearchCommonDBRequest
	bindErr := c.BindJSON(&form)
	if bindErr != nil {
		_ = c.Error(&gin.Error{
			Err:  bindErr,
			Type: service.ParamErr,
		})
		return
	}
	var res []model.SearchCommonDBPreview
	var err error
	var total = 0
	switch form.Module {
	case 5:
		var dissertationCount = 0
		var bookCount = 0
		var projectCount = 0
		var articleCount = 0
		res, err = service.SearchService{}.SearchInCommonDB(form)
		if err != nil {
			_ = c.Error(&gin.Error{
				Err:  err,
				Type: service.ParamErr,
			})
			return
		}
		bookCount, dissertationCount, articleCount, projectCount, err = service.SearchService{}.CountModuleInCommonDB(form)
		if err != nil {
			_ = c.Error(&gin.Error{
				Err:  err,
				Type: service.ParamErr,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"list":               res,
			"dissertationsCount": dissertationCount,
			"booksCount":         bookCount,
			"projectsCount":      projectCount,
			"articlesCount":      articleCount,
		})
		return
	case 1:
		res, total, err = service.SearchService{}.SearchCommonDBProject(form)
	case 2:
		res, total, err = service.SearchService{}.SearchCommonDBBook(form)
	case 3:
		res, total, err = service.SearchService{}.SearchCommonDBDissertation(form)
	case 4:
		res, total, err = service.SearchService{}.SearchCommonDBArticle(form)
	}
	if err != nil {
		_ = c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"list":  res,
		"total": total,
	})
}

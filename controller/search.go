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
	switch form.Module {
	case 0:
		res, err = service.SearchService{}.SearchInCommonDB(form)
	case 1:
		res, err = service.SearchService{}.SearchCommonDBProject(form)
	case 2:
		res, err = service.SearchService{}.SearchCommonDBBook(form)
	case 3:
		res, err = service.SearchService{}.SearchCommonDBDissertation(form)
	case 4:
		res, err = service.SearchService{}.SearchCommonDBArticle(form)
	}
	if err != nil {
		_ = c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"list": res,
	})
}

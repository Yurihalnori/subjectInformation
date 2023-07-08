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
	res, err := service.SearchService{}.SearchInCommonDB(form)
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

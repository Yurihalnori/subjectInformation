package controller

import (
	"github.com/gin-gonic/gin"
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
	service.SearchService{}.SearchInCommonDB(form)
}

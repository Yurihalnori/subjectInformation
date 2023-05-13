package controller

import (
	"fmt"
	"net/http"

	"subjectInformation/model"
	"subjectInformation/service"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
}

func (s ProjectController) Add(c *gin.Context) {
	var form model.ProjectForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	}

	databaseService := service.ProjectService{}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    databaseService.Add(form),
	})
}

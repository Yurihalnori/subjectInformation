package controller

import (
	"fmt"
	"net/http"

	"subjectInformation/model"
	"subjectInformation/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UniqueController struct {
}

func (s UniqueController) Add(c *gin.Context) {
	var form model.UniqueForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(form); err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	}

	UniqueService := service.UniqueService{}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    UniqueService.Add(form),
	})
}

// func (s UniqueController) Change(c *gin.Context) {
// 	var form model.UniqueForm
// 	if err := c.ShouldBind(&form); err != nil {
// 		fmt.Printf("controller %v", err)
// 		c.Error(&gin.Error{
// 			Err:  err,
// 			Type: service.ParamErr,
// 		})
// 		return
// 	}

// 	UniqueService := service.UniqueService{}

// 	c.JSON(http.StatusOK, Response{
// 		Success: true,
// 		Data:    UniqueService.Change(form),
// 	})
// }

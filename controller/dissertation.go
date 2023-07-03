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

type DissertationController struct {
}

// 添加条目
func (s DissertationController) Add(c *gin.Context) {
	var form model.DissertationForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  errors.New("存在字段错误，请检查输入信息"),
			Type: service.ParamErr,
		})
		return
	}

	DissertationService := service.DissertationService{}
	data := DissertationService.Add(form)

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

// 修改条目
func (s DissertationController) Change(c *gin.Context) {
	var form model.DissertationOmitempty
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBind(&form); err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	}

	validate := validator.New()
	if err := validate.Var(id, "numeric"); err != nil {
		c.Error(&gin.Error{
			Err:  errors.New("请返回正确的id?"),
			Type: service.ParamErr,
		})
		return
	}

	DissertationService := service.DissertationService{}
	data, err := DissertationService.Change(form, id)

	if err == nil {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Data:    data,
		})
	} else {
		c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
	}
}

// 删除条目
func (s DissertationController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	validate := validator.New()
	if err := validate.Var(id, "numeric"); err != nil {
		c.Error(&gin.Error{
			Err:  errors.New("请返回正确的id?"),
			Type: service.ParamErr,
		})
		return
	}

	DissertationController := service.DissertationService{}
	err := DissertationController.Delete(id)

	if err == nil {
		c.JSON(http.StatusOK, Response{
			Success: true,
		})
	} else {
		c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
	}
}

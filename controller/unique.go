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

type UniqueController struct {
}

// 添加条目
func (s UniqueController) Add(c *gin.Context) {
	var form model.UniqueForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  errors.New("存在字段为空，请检查输入信息"),
			Type: service.ParamErr,
		})
		return
	}

	UniqueService := service.UniqueService{}
	data := UniqueService.Add(form)

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

// 修改条目
func (s UniqueController) Change(c *gin.Context) {
	var form model.UniqueDatabaseOmitempty
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

	UniqueService := service.UniqueService{}
	data, err := UniqueService.Change(form, id)

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
func (s UniqueController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	validate := validator.New()
	if err := validate.Var(id, "numeric"); err != nil {
		c.Error(&gin.Error{
			Err:  errors.New("请返回正确的id?"),
			Type: service.ParamErr,
		})
		return
	}

	UniqueService := service.UniqueService{}
	err := UniqueService.Delete(id)

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

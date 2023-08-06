package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"subjectInformation/global"
	"subjectInformation/model"
	"subjectInformation/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type InstituteController struct {
}

// 添加条目
func (s InstituteController) Add(c *gin.Context) {
	var form model.InstituteForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Printf("controller %v", err)
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.Error(&gin.Error{
				Err:  errors.New("非数据格式错误,maybe数据类型或其他错误"),
				Type: service.ParamErr,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   errs.Translate(global.Trans),
		})
		return
	}

	InstituteService := service.InstituteService{}
	data := InstituteService.Add(form)

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

// 修改条目
func (s InstituteController) Change(c *gin.Context) {
	var form model.InstituteOmitempty
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBind(&form); err != nil {
		fmt.Printf("controller %v", err)
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.Error(&gin.Error{
				Err:  errors.New("非数据格式错误,maybe数据类型或其他错误"),
				Type: service.ParamErr,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   errs.Translate(global.Trans),
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

	InstituteService := service.InstituteService{}
	data, err := InstituteService.Change(form, id)

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
func (s InstituteController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	validate := validator.New()
	if err := validate.Var(id, "numeric"); err != nil {
		c.Error(&gin.Error{
			Err:  errors.New("请返回正确的id?"),
			Type: service.ParamErr,
		})
		return
	}

	InstituteService := service.InstituteService{}
	err := InstituteService.Delete(id)

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

// 获取条目
func (s InstituteController) GetInfo(c *gin.Context) {
	var form model.GetInfoForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Printf("controller %v", err)
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.Error(&gin.Error{
				Err:  errors.New("非数据格式错误,maybe数据类型或其他错误"),
				Type: service.ParamErr,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   errs.Translate(global.Trans),
		})
		return
	}

	ListService := service.ListService{}
	data := ListService.GetList(form, "institute")

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

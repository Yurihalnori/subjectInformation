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

type ProjectController struct {
}

// 添加条目
func (s ProjectController) Add(c *gin.Context) {
	var form model.ProjectForm
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

	ProjectService := service.ProjectService{}
	data := ProjectService.Add(form)

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

// 修改条目
func (s ProjectController) Change(c *gin.Context) {
	var form model.ProjectOmitempty
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

	ProjectService := service.ProjectService{}
	data, err := ProjectService.Change(form, id)

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
func (s ProjectController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	validate := validator.New()
	if err := validate.Var(id, "numeric"); err != nil {
		c.Error(&gin.Error{
			Err:  errors.New("请返回正确的id?"),
			Type: service.ParamErr,
		})
		return
	}

	ProjectService := service.ProjectService{}
	err := ProjectService.Delete(id)

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
func (s ProjectController) GetInfo(c *gin.Context) {
	var form model.GetInfoForm
	if err := c.BindJSON(&form); err != nil {
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
	data := ListService.GetList(form, "project")

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

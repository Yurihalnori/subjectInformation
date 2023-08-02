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

type ArticleController struct {
}

// 添加条目
func (s ArticleController) Add(c *gin.Context) {
	var form model.ArticleForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  errors.New("存在字段错误，请检查输入信息"),
			Type: service.ParamErr,
		})
		return
	}

	ArticleService := service.ArticleService{}
	data := ArticleService.Add(form)

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

// 修改条目
func (s ArticleController) Change(c *gin.Context) {
	var form model.ArticleOmitempty
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

	ArticleService := service.ArticleService{}
	data, err := ArticleService.Change(form, id)

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
func (s ArticleController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	validate := validator.New()
	if err := validate.Var(id, "numeric"); err != nil {
		c.Error(&gin.Error{
			Err:  errors.New("请返回正确的id?"),
			Type: service.ParamErr,
		})
		return
	}

	ArticleService := service.ArticleService{}
	err := ArticleService.Delete(id)

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
func (s ArticleController) GetInfo(c *gin.Context) {
	var form model.GetInfoForm
	if err := c.ShouldBindJSON(&form); err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	}

	ListService := service.ListService{}
	data := ListService.GetList(form, "article")

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

package service

import (
	"errors"
	"subjectInformation/model"
)

type ArticleService struct {
}

func (h ArticleService) Add(form model.ArticleForm) interface{} {

	var Response []model.ArticleResponseMsg
	categoryService := Category{}

	for i := 0; i < form.Total; i++ {
		data := form.List[i]
		model.DB.Create(&data)
		categoryService.AddCategory(data.Category, data.Id, "articles")
		s := model.ArticleResponseMsg{
			Id: data.Id,
		}
		Response = append(Response, s)
	}
	return Response
}

func (h ArticleService) Change(form model.ArticleOmitempty, id int) (interface{}, error) {

	var data model.Article
	if result := model.DB.First(&data, id); result.Error != nil {
		return nil, errors.New("未找到对应id数据,请检查id是否正确")
	}

	model.DB.Model(&data).Updates(&form)
	model.DB.Save(&data)

	if form.Category != "" {
		categoryService := Category{}
		cateErr := categoryService.ChangeCategory(form.Category, data.Id, "articles")
		if cateErr != nil {
			return nil, cateErr
		}
	}

	Response := model.ArticleResponseMsg{
		Id: data.Id,
	}

	return Response, nil
}

func (h ArticleService) Delete(id int) error {
	var data model.Article
	if result := model.DB.First(&data, id); result.Error != nil {
		return errors.New("未找到对应id信息,请检查id是否正确")
	}
	model.DB.Delete(&model.Article{}, id)
	categoryService := Category{}
	cateErr := categoryService.DeleteCategory(id, "articles")
	return cateErr
}

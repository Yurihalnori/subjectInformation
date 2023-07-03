package service

import (
	"errors"
	"subjectInformation/model"
)

type ProjectService struct {
}

func (h ProjectService) Add(form model.ProjectForm) interface{} {

	var Response []model.ProjectResponseMsg
	categoryService := Category{}

	for i := 0; i < form.Total; i++ {
		data := form.List[i]
		model.DB.Create(&data)
		categoryService.AddCategory(data.Category, data.Id, "projects")
		s := model.ProjectResponseMsg{
			Id: data.Id,
		}
		Response = append(Response, s)
	}
	return Response
}

func (h ProjectService) Change(form model.ProjectOmitempty, id int) (interface{}, error) {

	var data model.Project
	if result := model.DB.First(&data, id); result.Error != nil {
		return nil, errors.New("未找到对应id数据,请检查id是否正确")
	}

	model.DB.Model(&data).Updates(&form)
	model.DB.Save(&data)

	if form.Category != "" {
		categoryService := Category{}
		cateErr := categoryService.ChangeCategory(form.Category, data.Id, "projects")
		if cateErr != nil {
			return nil, cateErr
		}
	}

	Response := model.ProjectResponseMsg{
		Id: data.Id,
	}

	return Response, nil
}

func (h ProjectService) Delete(id int) error {
	var data model.Project
	if result := model.DB.First(&data, id); result.Error != nil {
		return errors.New("未找到对应id信息,请检查id是否正确")
	}
	model.DB.Delete(&model.Project{}, id)
	categoryService := Category{}
	cateErr := categoryService.DeleteCategory(id, "projects")
	return cateErr
}

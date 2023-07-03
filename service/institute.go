package service

import (
	"errors"
	"subjectInformation/model"
)

type InstituteService struct {
}

func (h InstituteService) Add(form model.InstituteForm) interface{} {

	var Response []model.InstituteResponseMsg
	categoryService := Category{}

	for i := 0; i < form.Total; i++ {
		data := form.List[i]
		model.DB.Create(&data)
		categoryService.AddCategory(data.Category, data.Id, "institutes")
		s := model.InstituteResponseMsg{
			Id: data.Id,
		}
		Response = append(Response, s)
	}
	return Response
}

func (h InstituteService) Change(form model.InstituteOmitempty, id int) (interface{}, error) {

	var data model.Institute
	if result := model.DB.First(&data, id); result.Error != nil {
		return nil, errors.New("未找到对应id数据,请检查id是否正确")
	}

	model.DB.Model(&data).Updates(&form)
	model.DB.Save(&data)

	if form.Category != "" {
		categoryService := Category{}
		cateErr := categoryService.ChangeCategory(form.Category, data.Id, "institutes")
		if cateErr != nil {
			return nil, cateErr
		}
	}

	Response := model.ProjectResponseMsg{
		Id: data.Id,
	}

	return Response, nil
}

func (h InstituteService) Delete(id int) error {
	var data model.Institute
	if result := model.DB.First(&data, id); result.Error != nil {
		return errors.New("未找到对应id信息,请检查id是否正确")
	}
	model.DB.Delete(&model.Institute{}, id)
	categoryService := Category{}
	cateErr := categoryService.DeleteCategory(id, "institutes")
	return cateErr
}

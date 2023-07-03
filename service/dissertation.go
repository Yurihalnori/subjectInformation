package service

import (
	"errors"
	"subjectInformation/model"
)

type DissertationService struct {
}

func (h DissertationService) Add(form model.DissertationForm) interface{} {

	var Response []model.DissertationResponseMsg
	categoryService := Category{}

	for i := 0; i < form.Total; i++ {
		data := form.List[i]
		model.DB.Create(&data)
		categoryService.AddCategory(data.Category, data.Id, "dissertations")
		s := model.DissertationResponseMsg{
			Id: data.Id,
		}
		Response = append(Response, s)
	}
	return Response
}

func (h DissertationService) Change(form model.DissertationOmitempty, id int) (interface{}, error) {

	var data model.Dissertation
	if result := model.DB.First(&data, id); result.Error != nil {
		return nil, errors.New("未找到对应id数据,请检查id是否正确")
	}

	model.DB.Model(&data).Updates(&form)
	model.DB.Save(&data)

	if form.Category != "" {
		categoryService := Category{}
		cateErr := categoryService.ChangeCategory(form.Category, data.Id, "dissertations")
		if cateErr != nil {
			return nil, cateErr
		}
	}

	Response := model.DissertationResponseMsg{
		Id: data.Id,
	}

	return Response, nil
}

func (h DissertationService) Delete(id int) error {
	var data model.Dissertation
	if result := model.DB.First(&data, id); result.Error != nil {
		return errors.New("未找到对应id信息,请检查id是否正确")
	}
	model.DB.Delete(&model.Dissertation{}, id)
	categoryService := Category{}
	cateErr := categoryService.DeleteCategory(id, "dissertations")
	return cateErr
}

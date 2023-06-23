package service

import (
	"errors"
	"subjectInformation/model"
)

type UniqueService struct {
}

func (h UniqueService) Add(form model.UniqueForm) interface{} {

	var Response []model.UniqueResponseMsg
	categoryService := Category{}

	for i := 0; i < form.Total; i++ {
		data := form.List[i]
		model.DB.Create(&data)
		categoryService.AddCategory(data.Category, data.Id, "unique_databases")
		s := model.UniqueResponseMsg{
			Id: data.Id,
		}
		Response = append(Response, s)
	}
	return Response
}

func (h UniqueService) Change(form model.UniqueDatabaseOmitempty, id int) (interface{}, error) {

	var data model.UniqueDatabase
	if result := model.DB.First(&data, id); result.Error != nil {
		return nil, errors.New("未找到对应id数据,请检查id是否正确")
	}

	model.DB.Model(&data).Updates(&form)
	model.DB.Save(&data)

	categoryService := Category{}
	cateErr := categoryService.ChangeCategory(form.Category, data.Id, "unique_databases")

	Response := model.UniqueResponseMsg{
		Id: data.Id,
	}

	return Response, cateErr
}

func (h UniqueService) Delete(id int) error {
	var data model.UniqueDatabase
	if result := model.DB.First(&data, id); result.Error != nil {
		return errors.New("未找到对应id信息,请检查id是否正确")
	}
	model.DB.Delete(&model.UniqueDatabase{}, id)
	categoryService := Category{}
	cateErr := categoryService.DeleteCategory(id, "unique_databases")
	return cateErr
}

package service

import (
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

func (h UniqueService) Change(form model.UniqueDatabaseOmitempty, id int) interface{} {
	var data model.UniqueDatabase
	model.DB.First(&data, id)
	model.DB.Model(&data).Updates(&form)
	model.DB.Save(&data)

	categoryService := Category{}
	categoryService.ChangeCategory(form.Category, data.Id, "unique_databases")

	Response := model.UniqueResponseMsg{
		Id: data.Id,
	}

	return Response
}

func (h UniqueService) Delete(id int) bool {
	var data model.UniqueDatabase
	if result := model.DB.First(&data, id); result.Error != nil {
		return false
	}
	model.DB.Delete(&model.UniqueDatabase{}, id)
	return true
}

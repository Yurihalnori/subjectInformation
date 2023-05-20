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

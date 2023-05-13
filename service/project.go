package service

import (
	"subjectInformation/model"
)

type ProjectService struct {
}

func (h ProjectService) Add(form model.ProjectForm) interface{} {
	var Response []model.Message
	for i := 0; i < form.Total; i++ {
		model.DB.Create(&form.List[i]).Table("projects")
		s := model.Message{
			Check: "",
			Id:    form.List[i].Id,
			Title: form.List[i].Name,
		}
		Response = append(Response, s)
	}
	return Response
}

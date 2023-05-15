package service

import (
	"fmt"
	"subjectInformation/model"
)

type ProjectService struct {
}

func (h ProjectService) Add(form model.ProjectForm) interface{} {

	var Response []model.Message
	categoryService := Category{}

	for i := 0; i < form.Total; i++ {
		data := form.List[i]
		model.DB.Create(&data)
		categoryService.AddCategory(data.Category, data.Id, "projects")
		s := model.Message{
			Check: "",
			Id:    data.Id,
			Title: data.Name,
		}
		Response = append(Response, s)
	}
	return Response
}

func (h ProjectService) Change(form model.Project) interface{} {
	return fmt.Sprint(1)
}

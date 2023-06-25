package service

import (
	"errors"
	"subjectInformation/model"
)

type TeamworkService struct {
}

func (h TeamworkService) Add(form model.TeamworkForm) interface{} {

	var Response []model.TeamworkResponseMsg
	categoryService := Category{}

	for i := 0; i < form.Total; i++ {
		data := form.List[i]
		model.DB.Create(&data)
		categoryService.AddCategory(data.Category, data.Id, "teamworks")
		s := model.TeamworkResponseMsg{
			Id: data.Id,
		}
		Response = append(Response, s)
	}
	return Response
}

func (h TeamworkService) Change(form model.TeamworkOmitempty, id int) (interface{}, error) {

	var data model.Teamwork
	if result := model.DB.First(&data, id); result.Error != nil {
		return nil, errors.New("未找到对应id数据,请检查id是否正确")
	}

	model.DB.Model(&data).Updates(&form)
	model.DB.Save(&data)

	if form.Category != "" {
		categoryService := Category{}
		cateErr := categoryService.ChangeCategory(form.Category, data.Id, "teamworks")
		if cateErr != nil {
			return nil, cateErr
		}
	}

	Response := model.TeamworkResponseMsg{
		Id: data.Id,
	}

	return Response, nil
}

func (h TeamworkService) Delete(id int) error {
	var data model.Teamwork
	if result := model.DB.First(&data, id); result.Error != nil {
		return errors.New("未找到对应id信息,请检查id是否正确")
	}
	model.DB.Delete(&model.Teamwork{}, id)
	categoryService := Category{}
	cateErr := categoryService.DeleteCategory(id, "teamworks")
	return cateErr
}

package service

import (
	"errors"
	"fmt"
	"subjectInformation/model"
)

type TutorService struct {
}

func (h TutorService) Add(form model.TutorForm) interface{} {

	var Response []model.TutorResponseMsg
	categoryService := Category{}

	for i := 0; i < form.Total; i++ {
		data := form.List[i]

		if data.InstituteID != 0 {
			var Institute model.Institute
			result := model.DB.Preload("Tutors").Find(&Institute, int(data.InstituteID))
			fmt.Println(Institute)
			if result.RowsAffected == 0 {
				return errors.New("未找到对应学位点信息")
			} else {
				model.DB.Create(&data)
				model.DB.Model(&Institute).Association("Tutors").Append([]model.Tutor{data})
			}
		}

		categoryService.AddCategory(data.Category, data.Id, "tutors")
		s := model.TutorResponseMsg{
			Id: data.Id,
		}
		Response = append(Response, s)
	}
	return Response
}

func (h TutorService) Change(form model.TutorOmitempty, id int) (interface{}, error) {

	var data model.Tutor
	if result := model.DB.First(&data, id); result.Error != nil {
		return nil, errors.New("未找到对应id数据,请检查id是否正确")
	}

	model.DB.Model(&data).Updates(&form)
	model.DB.Save(&data)

	if form.Category != "" {
		categoryService := Category{}
		cateErr := categoryService.ChangeCategory(form.Category, data.Id, "tutors")
		if cateErr != nil {
			return nil, cateErr
		}
	}

	Response := model.TutorResponseMsg{
		Id: data.Id,
	}

	return Response, nil
}

func (h TutorService) Delete(id int) error {
	var data model.Tutor
	if result := model.DB.First(&data, id); result.Error != nil {
		return errors.New("未找到对应id信息,请检查id是否正确")
	}
	model.DB.Delete(&model.Tutor{}, id)
	categoryService := Category{}
	cateErr := categoryService.DeleteCategory(id, "tutors")
	return cateErr
}

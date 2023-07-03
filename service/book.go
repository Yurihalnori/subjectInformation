package service

import (
	"errors"
	"subjectInformation/model"
)

type BookService struct {
}

func (h BookService) Add(form model.BookForm) interface{} {

	var Response []model.BookResponseMsg
	categoryService := Category{}

	for i := 0; i < form.Total; i++ {
		data := form.List[i]
		model.DB.Create(&data)
		categoryService.AddCategory(data.Category, data.Id, "books")
		s := model.BookResponseMsg{
			Id: data.Id,
		}
		Response = append(Response, s)
	}
	return Response
}

func (h BookService) Change(form model.BookOmitempty, id int) (interface{}, error) {

	var data model.Book
	if result := model.DB.First(&data, id); result.Error != nil {
		return nil, errors.New("未找到对应id数据,请检查id是否正确")
	}

	model.DB.Model(&data).Updates(&form)
	model.DB.Save(&data)

	if form.Category != "" {
		categoryService := Category{}
		cateErr := categoryService.ChangeCategory(form.Category, data.Id, "books")
		if cateErr != nil {
			return nil, cateErr
		}
	}

	Response := model.BookResponseMsg{
		Id: data.Id,
	}

	return Response, nil
}

func (h BookService) Delete(id int) error {
	var data model.Book
	if result := model.DB.First(&data, id); result.Error != nil {
		return errors.New("未找到对应id信息,请检查id是否正确")
	}
	model.DB.Delete(&model.Book{}, id)
	categoryService := Category{}
	cateErr := categoryService.DeleteCategory(id, "books")
	return cateErr
}

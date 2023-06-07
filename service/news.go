package service

import (
	"subjectInformation/model"
)

func AddOneNews(form model.News) (news model.News, err error) {
	res := model.DB.Select("title", "module", "region", "department", "date", "text").Create(&form)
	err = res.Error
	Category{}.AddCategory(form.Category, form.Id, "News")
	return form, err
}

func GetSomeNews(form model.GetSomeNews) (newsList []model.News, err error) {
	res := model.DB.
		Limit(form.Limit).
		Offset((form.Page - 1) * form.Limit).
		Joins("INNER JOIN categories ON categories.foreign_key = news.id").
		Find(&newsList)
	err = res.Error
	return
}

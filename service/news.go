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

type NewsResult struct {
	model.News
	Category1  bool
	Category2  bool
	Category3  bool
	Category4  bool
	Category5  bool
	Category6  bool
	Category7  bool
	Category8  bool
	Category9  bool
	Category10 bool
}

func GetSomeNews(form model.GetSomeNews) (newsList []NewsResult, err error) {
	res := model.DB.
		Model(&model.News{}).
		Select("news.id", "news.text", "categories.Category1").
		Limit(form.Limit).
		Offset((form.Page-1)*form.Limit).
		Joins("INNER JOIN categories ON categories.foreign_key = news.id AND categories.tablee = ?", "News").
		Scan(&newsList)
	err = res.Error
	return
}

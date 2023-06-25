package service

import (
	"strconv"
	"subjectInformation/model"
)

func AddOneNews(form model.News) (news model.News, err error) {
	res := model.DB.Select("title", "module", "region", "department", "date", "text", "category").Create(&form)
	err = res.Error
	Category{}.AddCategory(form.Category, form.Id, "News")
	return form, err
}

func GetSomeNews(form model.GetSomeNews) (newsList []model.News, err error) {
	condition := "INNER JOIN categories ON categories.foreign_key = news.id AND categories.tablee = 'News' AND( "
	for key, value := range form.Category {
		if value == 49 { // ascii 1 = 49
			condition += "categories.category" + strconv.Itoa(key+1) + " ='1'" + " OR "
		}
	}
	condition += "1=0)"
	res := model.DB.
		Model(&model.News{}).
		Limit(form.Limit).
		Offset((form.Page - 1) * form.Limit).
		Joins(condition).
		Scan(&newsList)
	err = res.Error
	return
}

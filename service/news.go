package service

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	"subjectInformation/model"
)

func AddOneNews(form model.News) (news model.News, err error) {
	res := model.DB.Select("title", "module", "region", "department", "date", "text", "category").Create(&form)
	err = res.Error
	Category{}.AddCategory(form.Category, form.Id, "News")
	return form, err
}

func GetSomeNews(form model.GetSomeNews) (newsList []model.NewsPreview, total int, err error) {
	order := ""
	if form.Name == "" || form.Order == "" {
		order = "date desc"
	} else {
		switch form.Name {
		case "time":
			switch form.Order {
			case "increase":
				{
					order = "date asc"
				}
			case "decrease":
				{
					order = "date desc"
				}
			}
		case "popularity":
			switch form.Order {
			case "increase":
				{
					order = "click asc"
				}
			case "decrease":
				{
					order = "click desc"
				}
			}
		}
	}
	condition := "INNER JOIN categories ON categories.foreign_key = news.id AND categories.tablee = 'News' AND( "
	for key, value := range form.Category {
		if value == 49 { // ascii 1 = 49
			condition += "categories.category" + strconv.Itoa(key+1) + " ='1'" + " OR "
		}
	}
	condition += "1=0)"
	res := model.DB.
		Model(&model.News{}).
		Select("news.id", "news.title", "news.category", "news.module", "news.region", "news.department", "news.click", "news.date").
		Where("news.module = ?", form.Module).
		Limit(form.Limit).
		Offset((form.Page - 1) * form.Limit).
		Joins(condition).
		Order(order).
		Scan(&newsList)
	err = res.Error
	return
}

func GetOneNew(id int) (ApieceOfNews model.NewsDetail, err error) {
	var news model.News
	res := model.DB.
		Table("news").
		Where("id = ?", id).
		First(&news).
		Scan(&ApieceOfNews)
	err = res.Error
	return
}

func ClickNewsOnce(id int) (err error) {
	res := model.DB.
		Table("news").
		Where("id = ?", id).
		Update("click", gorm.Expr("click + ?", 1))
	err = res.Error
	return
}

func EditNews(form model.News) (err error) {
	var news model.News
	res := model.DB.
		Model(&news).
		Where("id = ?", form.Id).
		Updates(form)
	num := res.RowsAffected
	if num == 0 {
		err = errors.New("can't find")
	} else {
		err = res.Error
	}
	return
}

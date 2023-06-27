package service

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	"subjectInformation/model"
)

type NewsService struct {
}

func (NewsService) AddOneNews(form model.News) (news model.News, err error) {
	res := model.DB.Select("title", "module", "region", "department", "date", "text", "category").Create(&form)
	err = res.Error
	Category{}.AddCategory(form.Category, form.Id, "News")
	return form, err
}

func (NewsService) GetSomeNews(form model.GetSomeNews) (newsList []model.NewsPreview, total int, err error) {
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
	module := ""
	if form.Module == 4 {
		module = " 4 = ? "
	} else {
		module = "news.module = ?"
	}
	res := model.DB.
		Model(&model.News{}).
		Select("news.id", "news.title", "news.category", "news.module", "news.region", "news.department", "news.click", "news.date").
		Where(module, form.Module).
		Limit(form.Limit).
		Offset((form.Page - 1) * form.Limit).
		Joins(condition).
		Order(order).
		Scan(&newsList)
	err = res.Error
	return
}

func (NewsService) GetOneNew(id int) (ApieceOfNews model.NewsDetail, err error) {
	var news model.News
	res := model.DB.
		Table("news").
		Where("id = ?", id).
		First(&news).
		Scan(&ApieceOfNews)
	err = res.Error
	return
}

func (NewsService) ClickNewsOnce(id int) (err error) {
	res := model.DB.
		Table("news").
		Where("id = ?", id).
		Update("click", gorm.Expr("click + ?", 1))
	err = res.Error
	return
}

func (NewsService) EditNews(form model.News) (err error) {
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
	if err == nil {
		if form.Category != "" {
			err = Category{}.ChangeCategory(form.Category, form.Id, "News")
		}
	}
	return
}

func (NewsService) DeleteNews(id int) (err error) {
	res := model.DB.Delete(model.News{}, id)
	num := res.RowsAffected
	if num == 0 {
		err = errors.New("can't find")
	} else {
		err = res.Error
	}
	if err == nil {
		err = Category{}.DeleteCategory(id, "News")
	}
	return
}

func (NewsService) SearchNews(form model.NewsSearchRequest) (total int, newsList []model.NewsPreview, err error) {
	join := " INNER JOIN categories ON categories.foreign_key = news.id AND categories.tablee = 'News' AND ( "
	for key, value := range form.Category {
		if value == 49 { // ascii 1 = 49
			join += "categories.category" + strconv.Itoa(key+1) + " ='1'" + " OR "
		}
	}
	join += "1=0) "

	module := ""
	if form.Module != 4 {
		module = "`module` = " + strconv.Itoa(int(form.Module)) + " AND "
	}

	condition := "SELECT * , MATCH(title) AGAINST ('" +
		form.Content + "' IN NATURAL LANGUAGE MODE)" +
		" AS title_score,  MATCH(text) AGAINST ( '" +
		form.Content + "' IN NATURAL LANGUAGE MODE) AS text_score " +
		"FROM news " + join + "WHERE " + module + "MATCH(title, text) AGAINST ('" +
		form.Content + "' IN NATURAL LANGUAGE MODE) "

	order := "ORDER BY "

	if form.Name == "" || form.Order == "" || form.Name == "relativity" {
		order += "(title_score + text_score) DESC"
	} else {
		switch form.Name {
		case "time":
			switch form.Order {
			case "increase":
				{
					order += "date asc"
				}
			case "decrease":
				{
					order += "date desc"
				}
			}
		case "popularity":
			switch form.Order {
			case "increase":
				{
					order += "click asc"
				}
			case "decrease":
				{
					order += "click desc"
				}
			}
		}
	}

	var count []model.NewsPreview
	model.DB.Raw(condition).Scan(&count)
	total = len(count)

	condition += order

	limit := " LIMIT " + strconv.Itoa((form.Page-1)*form.Limit) + " , " + strconv.Itoa(form.Limit)

	condition += limit

	model.DB.Raw(condition).Scan(&newsList)
	return
}

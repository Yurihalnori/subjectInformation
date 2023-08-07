package service

import (
	"fmt"
	"strconv"
	"subjectInformation/model"
)

type ListService struct {
}

func (h ListService) Order(field string, order string) (orderstr string) {
	if field == "" {
		orderstr = "id"
	} else {
		orderstr = field
	}
	if order == "desc" {
		orderstr = orderstr + " desc"
	}
	return
}

func (h ListService) GetList(form model.GetInfoForm, table string) interface{} {
	orderstr := h.Order(form.Field, form.Order)

	condition := "INNER JOIN categories ON categories.foreign_key =" +
		table + "s.id AND categories.tablee = '" + table + "s" + "' AND( "
	for key, value := range form.Category {
		if value == 49 { // ascii 1 = 49
			condition += "categories.category" + strconv.Itoa(key+1) + " ='1'" + " OR "
		}
	}
	condition += "1=0)"

	switch table {
	case "unique_database":
		{
			var list []model.UniqueDatabase
			result := model.DB.Joins(condition).
				Limit(form.Limit).Order(orderstr).
				Offset(form.Page*form.Limit - 10).Find(&list)
			var data = model.UniqueList{
				Total:          int(result.RowsAffected),
				List:           list,
				CategoryNumber: h.GetCategory("unique_databases"),
			}
			return data
		}
	case "article":
		{
			var list []model.Article
			result := model.DB.Limit(form.Limit).Order(orderstr).Offset(form.Page*form.Limit - 10).Find(&list)
			var data = model.ArticleList{
				Total:          int(result.RowsAffected),
				List:           list,
				CategoryNumber: h.GetCategory("articles"),
			}
			return data
		}
	case "book":
		{
			var list []model.Book
			result := model.DB.Limit(form.Limit).Order(orderstr).Offset(form.Page*form.Limit - 10).Find(&list)
			var data = model.BookList{
				Total:          int(result.RowsAffected),
				List:           list,
				CategoryNumber: h.GetCategory("books"),
			}
			return data
		}
	case "dissertation":
		{
			var list []model.Dissertation
			result := model.DB.Limit(form.Limit).Order(orderstr).Offset(form.Page*form.Limit - 10).Find(&list)
			var data = model.DissertationList{
				Total:          int(result.RowsAffected),
				List:           list,
				CategoryNumber: h.GetCategory("dissertations"),
			}
			return data
		}
	case "institute":
		{
			var list []model.Institute
			result := model.DB.Limit(form.Limit).Order(orderstr).Offset(form.Page*form.Limit - 10).Find(&list)
			var data = model.InstituteList{
				Total:          int(result.RowsAffected),
				List:           list,
				CategoryNumber: h.GetCategory("institutes"),
			}
			return data
		}
	case "project":
		{
			var list []model.Project
			result := model.DB.Limit(form.Limit).Order(orderstr).Offset(form.Page*form.Limit - 10).Find(&list)
			var data = model.ProjectList{
				Total:          int(result.RowsAffected),
				List:           list,
				CategoryNumber: h.GetCategory("projects"),
			}
			return data
		}
	case "teamwork":
		{
			var list []model.Teamwork
			result := model.DB.Limit(form.Limit).Order(orderstr).Offset(form.Page*form.Limit - 10).Find(&list)
			var data = model.TeamworkList{
				Total:          int(result.RowsAffected),
				List:           list,
				CategoryNumber: h.GetCategory("teamworks"),
			}
			return data
		}
	case "tutor":
		{
			var list []model.Tutor
			result := model.DB.Limit(form.Limit).Order(orderstr).Offset(form.Page*form.Limit - 10).Find(&list)
			var data = model.TutorList{
				Total:          int(result.RowsAffected),
				List:           list,
				CategoryNumber: h.GetCategory("tutors"),
			}
			return data
		}
	}
	return nil
}

func (h ListService) GetCategory(table string) [10]int64 {
	var category [10]int64
	model.DB.Where("category1 = 1 AND tablee = ?", table).Table("categories").Count(&category[0])
	model.DB.Where("category2 = 1 AND tablee = ?", table).Table("categories").Count(&category[1])
	model.DB.Where("category3 = 1 AND tablee = ?", table).Table("categories").Count(&category[2])
	model.DB.Where("category4 = 1 AND tablee = ?", table).Table("categories").Count(&category[3])
	model.DB.Where("category5 = 1 AND tablee = ?", table).Table("categories").Count(&category[4])
	model.DB.Where("category6 = 1 AND tablee = ?", table).Table("categories").Count(&category[5])
	model.DB.Where("category7 = 1 AND tablee = ?", table).Table("categories").Count(&category[6])
	model.DB.Where("category8 = 1 AND tablee = ?", table).Table("categories").Count(&category[7])
	model.DB.Where("category9 = 1 AND tablee = ?", table).Table("categories").Count(&category[8])
	model.DB.Where("category10 = 1 AND tablee = ?", table).Table("categories").Count(&category[9])
	fmt.Println(category)
	return category
}

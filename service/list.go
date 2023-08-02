package service

import (
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
	switch table {
	case "unique":
		{
			var list []model.UniqueDatabase
			result := model.DB.Limit(form.Limit).Order(orderstr).Offset(form.Page*form.Limit - 10).Find(&list)
			var data = model.UniqueList{
				Total: int(result.RowsAffected),
				List:  list,
			}
			return data
		}
	case "article":
		{
			var list []model.Article
			result := model.DB.Limit(form.Limit).Order(orderstr).Offset(form.Page*form.Limit - 10).Find(&list)
			var data = model.ArticleList{
				Total: int(result.RowsAffected),
				List:  list,
			}
			return data
		}
	case "book":
		{
			var list []model.Book
			result := model.DB.Limit(form.Limit).Order(orderstr).Offset(form.Page*form.Limit - 10).Find(&list)
			var data = model.BookList{
				Total: int(result.RowsAffected),
				List:  list,
			}
			return data
		}
	case "dissertation":
		{
			var list []model.Dissertation
			result := model.DB.Limit(form.Limit).Order(orderstr).Offset(form.Page*form.Limit - 10).Find(&list)
			var data = model.DissertationList{
				Total: int(result.RowsAffected),
				List:  list,
			}
			return data
		}
	case "institute":
		{
			var list []model.Institute
			result := model.DB.Limit(form.Limit).Order(orderstr).Offset(form.Page*form.Limit - 10).Find(&list)
			var data = model.InstituteList{
				Total: int(result.RowsAffected),
				List:  list,
			}
			return data
		}
	case "project":
		{
			var list []model.Project
			result := model.DB.Limit(form.Limit).Order(orderstr).Offset(form.Page*form.Limit - 10).Find(&list)
			var data = model.ProjectList{
				Total: int(result.RowsAffected),
				List:  list,
			}
			return data
		}
	case "teamwork":
		{
			var list []model.Teamwork
			result := model.DB.Limit(form.Limit).Order(orderstr).Offset(form.Page*form.Limit - 10).Find(&list)
			var data = model.TeamworkList{
				Total: int(result.RowsAffected),
				List:  list,
			}
			return data
		}
	case "tutor":
		{
			var list []model.Tutor
			result := model.DB.Limit(form.Limit).Order(orderstr).Offset(form.Page*form.Limit - 10).Find(&list)
			var data = model.TutorList{
				Total: int(result.RowsAffected),
				List:  list,
			}
			return data
		}
	}
	return nil
}

package service

import "subjectInformation/model"

func AddOneNews(form model.News) (news model.News, err error) {
	res := model.DB.Select("title", "module", "region", "department", "date", "content").Create(&form)
	err = res.Error
	return
}
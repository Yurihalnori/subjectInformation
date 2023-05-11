package service

import (
	"subjectInformation/model"
)

type DatabaseService struct {
}

type message struct {
	Check string `form:"check" json:"check"`
	Id    int    `form:"id" json:"id"`
	Title string `form:"title" json:"title"`
}

func (h DatabaseService) Add(form model.Form) interface{} {
	//最后返回信息的切片
	s := []message{}
	for i := range form.List {
		module := form.List[i].Module
		switch module {
		case "0":
			s = append(s, Add0(form, i)) //科研项目
		case "1":
			s = append(s, Add1(form, i)) //学位点
		case "2":
			s = append(s, Add2(form, i)) //学科图书
		case "3":
			s = append(s, Add3(form, i)) //学位论文
		case "4":
			s = append(s, Add4(form, i)) //期刊论文
		}
	}
	return s
}

func Add0(form model.Form, i int) message {
	data := model.Project{
		Name:           form.List[i].Name,
		Classification: form.List[i].Classification,
		Sponsor:        form.List[i].Sponsor,
		ApprovalNumber: form.List[i].ApprovalNumber,
		CreateDate:     form.List[i].CreateDate,
		Superintendent: form.List[i].Superintendent,
		Organization:   form.List[i].Organization,
		Region:         form.List[i].Region,
		Click:          form.List[i].Click,
		Download:       form.List[i].Download,
	}
	model.DB.Create(&data)
	this := message{
		Check: "ok",
		Id:    data.Id,
		Title: data.Name,
	}
	return this
}

func Add1(form model.Form, i int) message {
	data := model.Institute{
		Name:           form.List[i].Name,
		Classification: form.List[i].Classification,
		Nation:         form.List[i].Nation,
		Province:       form.List[i].Province,
		City:           form.List[i].City,
		University:     form.List[i].University,
		College:        form.List[i].College,
		Click:          form.List[i].Click,
	}
	model.DB.Create(&data)
	this := message{
		Check: "ok",
		Id:    data.Id,
		Title: data.Name,
	}
	return this
}

func Add2(form model.Form, i int) message {
	data := model.Book{
		Name:      form.List[i].Name,
		Nation:    form.List[i].Nation,
		Language:  form.List[i].Language,
		Author:    form.List[i].Author,
		Publisher: form.List[i].Publisher,
		Time:      form.List[i].CreateDate,
		Digest:    form.List[i].Digest,
		Text:      form.List[i].Text,
	}
	model.DB.Create(&data)
	this := message{
		Check: "ok",
		Id:    data.Id,
		Title: data.Name,
	}
	return this
}

func Add3(form model.Form, i int) message {
	data := model.Dissertation{
		Title:      form.List[i].Title,
		Author:     form.List[i].Author,
		Tutor:      form.List[i].Tutor,
		Province:   form.List[i].Province,
		City:       form.List[i].City,
		University: form.List[i].University,
		College:    form.List[i].College,
		Year:       form.List[i].Year,
		Technique:  form.List[i].Technique,
		KeyWord:    form.List[i].KeyWord,
		Digest:     form.List[i].Digest,
		Data:       form.List[i].Data,
		Click:      form.List[i].Click,
		Download:   form.List[i].Download,
	}
	model.DB.Create(&data)
	this := message{
		Check: "ok",
		Id:    data.Id,
		Title: data.Title,
	}
	return this
}

func Add4(form model.Form, i int) message {
	data := model.Article{
		Nation:       form.List[i].Nation,
		Periodical:   form.List[i].Periodical,
		Title:        form.List[i].Title,
		Author:       form.List[i].Author,
		Organization: form.List[i].Organization,
		CreateDate:   form.List[i].CreateDate,
		Technique:    form.List[i].Technique,
		KeyWord:      form.List[i].KeyWord,
		Digest:       form.List[i].Digest,
		Data:         form.List[i].Data,
		Text:         form.List[i].Text,
		Click:        form.List[i].Click,
		Download:     form.List[i].Download,
	}
	model.DB.Create(&data)
	this := message{
		Check: "ok",
		Id:    data.Id,
		Title: data.Title,
	}
	return this
}

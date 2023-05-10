package service

import (
	"fmt"
	"subjectInformation/model"
	"time"
)

type DatabaseService struct {
}

func (h DatabaseService) Search(msg string) string {
	return fmt.Sprintf("hello %v", msg)
}

func (h DatabaseService) Change(date time.Time) string {
	return fmt.Sprintf("tomorrow is %v", date.Format("2006-01-02"))
}

func (h DatabaseService) Add(form model.Form) string {
	for i := range form.List {
		module := form.List[i].Module
		switch module {
		case "0":
			Add0(form, i)
		case "1":
			type typ model.Institute
			break
		case "2":
			type typ model.Book
			break
		case "3":
			type typ model.Article
			break
		}
	}
	return fmt.Sprintf("%v", form)
}

func Add0(form model.Form, i int) {
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
	fmt.Print(data)
	// model.DB.Create(&data).Table("")
}

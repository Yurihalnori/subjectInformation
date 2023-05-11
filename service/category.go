package service

import "subjectInformation/model"

type Category struct {
}

func (h Category) AddCategory(category string, foreignKey int, table string) {
	for i := 0; i < len(category); i++ {
		if category[i] == '1' {
			data := model.Category{
				Category:   i,
				ForeignKey: foreignKey,
				Table:      table,
			}
			model.DB.Create(&data)
		}
	}
}

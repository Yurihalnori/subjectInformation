package service

import "subjectInformation/model"

type Category struct {
}

func (h Category) AddCategory(category string, foreignKey int, table string) {
	data := model.Category{
		ForeignKey: foreignKey,
		Table:      table,
		Category1:  false,
		Category2:  false,
		Category3:  false,
		Category4:  false,
		Category5:  false,
		Category6:  false,
		Category7:  false,
		Category8:  false,
		Category9:  false,
		Category10: false,
	}

	if category[0] == '1' {
		data.Category1 = true
	}
	if category[1] == '1' {
		data.Category2 = true
	}
	if category[2] == '1' {
		data.Category3 = true
	}
	if category[3] == '1' {
		data.Category4 = true
	}
	if category[4] == '1' {
		data.Category5 = true
	}
	if category[5] == '1' {
		data.Category6 = true
	}
	if category[6] == '1' {
		data.Category7 = true
	}
	if category[7] == '1' {
		data.Category8 = true
	}
	if category[8] == '1' {
		data.Category9 = true
	}
	if category[9] == '1' {
		data.Category10 = true
	}
	model.DB.Create(&data)
}

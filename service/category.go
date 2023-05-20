package service

import (
	"subjectInformation/model"
)

type Category struct {
}

func DisposeCategory(category string, foreignKey int, table string) model.Category {
	data := model.Category{
		ForeignKey: foreignKey,
		Tablee:     table,
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
	return data
}

func (h Category) AddCategory(category string, foreignKey int, table string) {
	data := DisposeCategory(category, foreignKey, table)
	model.DB.Create(&data)
}

func (h Category) ChangeCategory(category string, foreignKey int, table string) {
	toData := DisposeCategory(category, foreignKey, table)
	var data model.Category
	result := model.DB.Where("foreign_key = ? AND tablee = ?", foreignKey, table).First(&data)
	if result.Error == nil {
		data.Category1 = toData.Category1
		data.Category2 = toData.Category2
		data.Category3 = toData.Category3
		data.Category4 = toData.Category4
		data.Category5 = toData.Category5
		data.Category6 = toData.Category6
		data.Category7 = toData.Category7
		data.Category8 = toData.Category8
		data.Category9 = toData.Category9
		data.Category10 = toData.Category10
		model.DB.Save(&data)
	}
}

package service

import (
	"errors"
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

func (h Category) ChangeCategory(category string, foreignKey int, table string) error {
	_Data := DisposeCategory(category, foreignKey, table)
	var data model.Category
	result := model.DB.Where("foreign_key = ? AND tablee = ?", foreignKey, table).First(&data)
	if result.Error == nil {
		data.Category1 = _Data.Category1
		data.Category2 = _Data.Category2
		data.Category3 = _Data.Category3
		data.Category4 = _Data.Category4
		data.Category5 = _Data.Category5
		data.Category6 = _Data.Category6
		data.Category7 = _Data.Category7
		data.Category8 = _Data.Category8
		data.Category9 = _Data.Category9
		data.Category10 = _Data.Category10
		model.DB.Save(&data)
		return nil
	}
	return errors.New("未找到对应category(理论上并不会发生,除非直接改了数据库)")
}

func (h Category) DeleteCategory(foreignKey int, table string) error {
	var data model.Category
	result := model.DB.Where("foreign_key = ? AND tablee = ?", foreignKey, table).First(&data)
	if result.Error == nil {
		model.DB.Delete(&model.Category{}, data.Id)
		return nil
	}
	return errors.New("未找到对应category(理论上并不会发生,除非直接改了数据库)")
}

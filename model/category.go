package model

import "time"

type Category struct { //学科分类
	Id         string    `json:"id"`         //主键
	Category   string    `json:"category"`   //学科分类 0-9
	ForeignKey string    `json:"foreignKey"` //外键
	Form       string    `json:"form"`       //来自哪张表
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

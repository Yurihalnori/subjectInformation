package model

import "time"

type Category struct { //学科分类
	Id         int       `json:"id"`         //主键
	Category   int       `json:"category"`   //学科分类 0-9
	ForeignKey int       `json:"foreignKey"` //外键
	Table      string    `json:"table"`      //来自哪张表
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

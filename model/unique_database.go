package model

import "time"

type UniqueDatabase struct {
	Id           int       `json:"id"`                               //序号
	Name         string    `json:"name" validate:"required"`         //名称
	Trimmer      string    `json:"trimmer" validate:"required"`      //整理者
	KeyWord      string    `json:"keyWord" validate:"required"`      //关键字
	Introduction string    `json:"introduction" validate:"required"` //简介
	Data         string    `json:"data" validate:"required"`         //数据
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Category     string    `json:"category" gorm:"-" validate:"required"` // 学科分类
}

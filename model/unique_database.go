package model

import "time"

type UniqueDatabase struct {
	Id           int       `json:"id"`                              //序号
	Name         string    `json:"name" binding:"required"`         //名称
	Trimmer      string    `json:"trimmer" binding:"required"`      //整理者
	KeyWord      string    `json:"keyWord" binding:"required"`      //关键字
	Introduction string    `json:"introduction" binding:"required"` //简介
	Data         string    `json:"data" binding:"required"`         //数据
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Category     string    `json:"category" gorm:"-" binding:"category"` // 学科分类
}

type UniqueDatabaseOmitempty struct {
	Id           int       `json:"id" binding:"omitempty"`           //序号
	Name         string    `json:"name" binding:"omitempty"`         //名称
	Trimmer      string    `json:"trimmer" binding:"omitempty"`      //整理者
	KeyWord      string    `json:"keyWord" binding:"omitempty"`      //关键字
	Introduction string    `json:"introduction" binding:"omitempty"` //简介
	Data         string    `json:"data" binding:"omitempty"`         //数据
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Category     string    `json:"category" gorm:"-" binding:"omitempty,category00"` // 学科分类
}

type UniqueResponseMsg struct {
	Id int `json:"id"`
}

type UniqueForm struct {
	Total int              `json:"total" binding:"numeric"`
	List  []UniqueDatabase `json:"list" binding:"required,dive"`
}

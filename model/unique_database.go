package model

import (
	"gorm.io/gorm"
	"time"
)

type UniqueDatabase struct {
	Id           int       `json:"id"`                              //序号
	Name         string    `json:"name" binding:"required"`         //名称
	Trimmer      string    `json:"trimmer" binding:"required"`      //整理者
	KeyWord      string    `json:"keyWord" binding:"required"`      //关键字
	Introduction string    `json:"introduction" binding:"required"` //简介
	Attachment   string    `json:"attachment" binding:"required"`   //附件
	Click        int       `json:"click"`                           //点击数
	Download     int       `json:"download"`                        //下载数
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Category     string    `json:"category" gorm:"-" binding:"category"` // 学科分类
	FurtherData  string    `gorm:"type:json"`
	Blank        string    `json:"blank"`
}

type UniqueDatabaseOmitempty struct {
	Id           int       `json:"id"`                               //序号
	Name         string    `json:"name" binding:"omitempty"`         //名称
	Trimmer      string    `json:"trimmer" binding:"omitempty"`      //整理者
	KeyWord      string    `json:"keyWord" binding:"omitempty"`      //关键字
	Introduction string    `json:"introduction" binding:"omitempty"` //简介
	Attachment   string    `json:"attachment" binding:"omitempty"`   //附件
	Click        int       `json:"click"`                            //点击数
	Download     int       `json:"download"`                         //下载数
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Category     string    `json:"category" gorm:"-" binding:"omitempty,category"` // 学科分类
	Blank        string    `json:"blank"`
}

type UniqueResponseMsg struct {
	Id int `json:"id"`
}

type UniqueForm struct {
	Total int              `json:"total" binding:"numeric"`
	List  []UniqueDatabase `json:"list" binding:"required,dive"`
}

type GetInfoForm struct {
	Page     int    `json:"page" binding:"numeric"`
	Limit    int    `json:"limit" binding:"numeric"`
	Field    string `json:"field" binding:"omitempty"`
	Order    string `json:"order" binding:"omitempty,oneof=desc asc"`
	Category string `json:"category" binding:"category"`
}

type UniqueList struct {
	Total          int              `json:"total"`
	List           []UniqueDatabase `json:"list"`
	CategoryNumber [10]int64        `json:"categoryNumber"`
}

func (uniqueDB *UniqueDatabase) BeforeSave(tx *gorm.DB) error {
	uniqueDB.FurtherData = "{}"
	return nil
}

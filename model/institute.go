package model

import "time"

type Institute struct { //学位点
	Id             int                    `json:"id"`                                            //序号
	Name           string                 `json:"name" binding:"omitempty"`                      //学位点名称
	University     string                 `json:"university" binding:"required"`                 //学校
	College        string                 `json:"college" binding:"required"`                    //学院
	Nation         string                 `json:"nation" binding:"required"`                     //国家
	Province       string                 `json:"province" binding:"required"`                   //省(通指一级地方单位)
	City           string                 `json:"city" binding:"required"`                       //市(通指二级地方单位)
	Classification string                 `json:"classification" binding:"required,oneof=0 1 2"` //学位点类型 0.学术硕士 1.专业硕士 2.博士
	Click          int                    `json:"click"`                                         //点击数
	CreatedAt      time.Time              `json:"createdAt"`
	UpdatedAt      time.Time              `json:"updatedAt"`
	Totors         []Tutor                `json:"tutors" binding:"omitempty"`
	Category       string                 `json:"category" gorm:"-" binding:"category"` // 学科分类
	FurtherData    map[string]interface{} `gorm:"type:json"`
	Blank          string                 `json:"blank"`
}

type InstituteOmitempty struct { //学位点
	Id             int       `json:"id" binding:"omitempty"`                         //序号
	Name           string    `json:"name" binding:"omitempty"`                       //学位点名称
	University     string    `json:"university" binding:"omitempty"`                 //学校
	College        string    `json:"college" binding:"omitempty"`                    //学院
	Nation         string    `json:"nation" binding:"omitempty"`                     //国家
	Province       string    `json:"province" binding:"omitempty"`                   //省(通指一级地方单位)
	City           string    `json:"city" binding:"omitempty"`                       //市(通指二级地方单位)
	Classification string    `json:"classification" binding:"omitempty,oneof=0 1 2"` //学位点类型 0.学术硕士 1.专业硕士 2.博士
	Click          int       `json:"click" binding:"omitempty"`                      //点击数
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Totors         []Tutor   `json:"tutors" binding:"omitempty"`
	Category       string    `json:"category" gorm:"-" binding:"omitempty,category"` // 学科分类
	Blank          string    `json:"blank"`
}

type InstituteResponseMsg struct {
	Id int `json:"id"`
}

type InstituteForm struct {
	Total int         `json:"total" binding:"numeric"`
	List  []Institute `json:"list" binding:"required,dive"`
}

type InstituteList struct {
	Total          int         `json:"total"`
	List           []Institute `json:"list"`
	CategoryNumber [10]int64   `json:"categoryNumber"`
}

package model

import "time"

type Tutor struct { //导师
	Id             int       `json:"id"`                                //序号
	Name           string    `json:"name" binding:"required"`           //导师名称
	University     string    `json:"university" binding:"required"`     //学校
	College        string    `json:"college" binding:"required"`        //学院
	Nation         string    `json:"nation" binding:"required"`         //国家
	Province       string    `json:"province" binding:"required"`       //省(通指一级地方单位)
	City           string    `json:"city" binding:"required"`           //市(通指二级地方单位)
	Direction      string    `json:"direction" binding:"required"`      //研究方向
	Title          string    `json:"title" binding:"required"`          //职称
	Contact        string    `json:"contact" binding:"required"`        //联系方式
	Classification string    `json:"classification" binding:"required"` //类型(博导 硕导)
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	// InstituteID    uint      `json:"InstituteID"`
	Category string `json:"category" gorm:"-" binding:"category"` // 学科分类
	Blank    string `json:"blank"`
}

type TutorOmitempty struct { //导师
	Id             int       `json:"id" binding:"omitempty"`             //序号
	Name           string    `json:"name" binding:"omitempty"`           //导师名称
	University     string    `json:"university" binding:"omitempty"`     //学校
	College        string    `json:"college" binding:"omitempty"`        //学院
	Nation         string    `json:"nation" binding:"omitempty"`         //国家
	Province       string    `json:"province" binding:"omitempty"`       //省(通指一级地方单位)
	City           string    `json:"city" binding:"omitempty"`           //市(通指二级地方单位)
	Direction      string    `json:"direction" binding:"omitempty"`      //研究方向
	Title          string    `json:"title" binding:"omitempty"`          //职称
	Contact        string    `json:"contact" binding:"omitempty"`        //联系方式
	Classification string    `json:"classification" binding:"omitempty"` //类型(博导 硕导)
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	// InstituteID    uint      `json:"InstituteID"`
	Category string `json:"category" gorm:"-" binding:"omitempty,category"` // 学科分类
	Blank    string `json:"blank"`
}

type TutorResponseMsg struct {
	Id int `json:"id"`
}

type TutorForm struct {
	Total int     `json:"total" binding:"numeric"`
	List  []Tutor `json:"list" binding:"required,dive"`
}

type TutorList struct {
	Total          int       `json:"total"`
	List           []Tutor   `json:"list"`
	CategoryNumber [10]int64 `json:"categoryNumber"`
}

package model

import "time"

type Tutor struct { //导师
	Id             int       `json:"id"`             //序号
	Name           string    `json:"name"`           //导师名称
	University     string    `json:"university"`     //学校
	College        string    `json:"college"`        //学院
	Nation         string    `json:"nation"`         //国家
	Province       string    `json:"province"`       //省(通指一级地方单位)
	City           string    `json:"city"`           //市(通指二级地方单位)
	Direction      string    `json:"direction"`      //研究方向
	Title          string    `json:"title"`          //职称
	Contact        string    `json:"contact"`        //联系方式
	Classification string    `json:"classification"` //类型(博导 硕导)
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	InstituteId    int       `json:"instituteId"` //外键
	Institute      Institute `binding:"omitempty"`
	Category       string    `json:"category" gorm:"-:migration"` // 学科分类
}

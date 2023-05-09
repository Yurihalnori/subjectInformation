package model

import "time"

type Institute struct {
	Id             int       `json:"id"`             //序号
	Name           string    `json:"name"`           //学位点名称
	University     string    `json:"university"`     //学校
	College        string    `json:"college"`        //学院
	Nation         string    `json:"nation"`         //国家
	Province       string    `json:"province"`       //省(通指一级地方单位)
	City           string    `json:"city"`           //市(通指二级地方单位)
	Classification string    `json:"classification"` //学位点类型 0.学术硕士 1.专业硕士 2.博士
	Click          int       `json:"click"`          //点击数
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

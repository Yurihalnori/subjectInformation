package model

import "time"

type Dissertation struct { //学位论文
	Id         int       `json:"id"`         //序号
	Title      string    `json:"title"`      //题目
	Author     string    `json:"author"`     //作者
	Tutor      string    `json:"tutor"`      //导师
	Province   string    `json:"province"`   //省(通指一级地方单位)
	City       string    `json:"city"`       //市(通指二级地方单位)
	University string    `json:"university"` //学校
	College    string    `json:"college"`    //学院
	Year       string    `json:"year"`       //年份
	Technique  string    `json:"technique"`  //研究方法
	KeyWord    string    `json:"keyWord"`    //关键字
	Digest     string    `json:"digest"`     //摘要
	Data       string    `json:"data"`       //数据
	Text       string    `json:"text"`       //全文
	Click      int       `json:"click"`      //点击数
	Download   int       `json:"download"`   //下载数
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

package model

import "time"

type Article struct { //期刊论文
	Id           int       `json:"id"`           //序号
	Nation       string    `json:"nation"`       //国家
	Title        string    `json:"title"`        //题名
	Periodical   string    `json:"periodical"`   //期刊名称
	Author       string    `json:"author"`       //作者
	Organization string    `json:"organization"` //作者单位
	CreateDate   string    `json:"CreateDate"`   //发表时间
	Technique    string    `json:"technique"`    //研究方法
	KeyWord      string    `json:"keyWord"`      //关键字
	Digest       string    `json:"digest"`       //摘要
	Data         string    `json:"data"`         //数据
	Text         string    `json:"text"`         //全文
	Click        int       `json:"click"`        //点击数
	Download     int       `json:"download"`     //下载数
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Category     string    `json:"category" gorm:"-:migration"` // 学科分类
}

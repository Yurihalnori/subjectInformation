package model

import "time"

type Book struct { //学科图书
	Id        int       `json:"id"`        //序号
	Name      string    `json:"name"`      //图书名称
	Nation    string    `json:"nation"`    //国家
	Language  string    `json:"language"`  //语言
	Author    string    `json:"author"`    //作者
	Publisher string    `json:"publisher"` //出版商
	Time      string    `json:"time"`      //出版年月
	Digest    string    `json:"digest"`    //摘要
	Text      string    `json:"text"`      //全文
	Click     int       `json:"click"`     //点击数
	Download  int       `json:"download"`  //下载数
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Category  string    `json:"category" gorm:"-:migration"` // 学科分类
}

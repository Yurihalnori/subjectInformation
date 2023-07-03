package model

import "time"

type Book struct { //学科图书
	Id        int       `json:"id"`                           //序号
	Name      string    `json:"name" binding:"required"`      //图书名称
	Nation    string    `json:"nation" binding:"required"`    //国家
	Language  string    `json:"language" binding:"required"`  //语言
	Author    string    `json:"author" binding:"required"`    //作者
	Publisher string    `json:"publisher" binding:"required"` //出版商
	Time      string    `json:"time" binding:"required"`      //出版年月
	Digest    string    `json:"digest" binding:"required"`    //摘要
	Text      string    `json:"text" binding:"required"`      //全文
	Click     int       `json:"click"`                        //点击数
	Download  int       `json:"download"`                     //下载数
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Category  string    `json:"category" gorm:"-" binding:"required"` // 学科分类
}

type BookOmitempty struct { //学科图书
	Id        int       `json:"id"`                            //序号
	Name      string    `json:"name" binding:"omitempty"`      //图书名称
	Nation    string    `json:"nation" binding:"omitempty"`    //国家
	Language  string    `json:"language" binding:"omitempty"`  //语言
	Author    string    `json:"author" binding:"omitempty"`    //作者
	Publisher string    `json:"publisher" binding:"omitempty"` //出版商
	Time      string    `json:"time" binding:"omitempty"`      //出版年月s
	Digest    string    `json:"digest" binding:"omitempty"`    //摘要
	Text      string    `json:"text" binding:"omitempty"`      //全文
	Click     int       `json:"click"`                         //点击数
	Download  int       `json:"download"`                      //下载数
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Category  string    `json:"category" gorm:"-" binding:"omitempty"` // 学科分类
}

type BookResponseMsg struct {
	Id int `json:"id"`
}

type BookForm struct {
	Total int    `json:"total" binding:"numeric"`
	List  []Book `json:"list" binding:"required,dive"`
}

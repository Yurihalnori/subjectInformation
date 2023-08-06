package model

import "time"

type Book struct { //学科图书
	Id          int                    `json:"id"`                                            //序号
	Title       string                 `json:"name" binding:"required"`                       //图书名称
	Nation      string                 `json:"nation" binding:"required"`                     //国家
	Language    string                 `json:"language" binding:"required"`                   //语言
	Author      string                 `json:"author" binding:"required"`                     //作者
	Publisher   string                 `json:"publisher" binding:"required"`                  //出版商
	Time        time.Time              `json:"time"  gorm:"type:datetime" binding:"required"` //出版年月
	Digest      string                 `json:"digest" binding:"required"`                     //摘要
	Text        string                 `json:"text"`                                          //全文
	Click       int                    `json:"click"`                                         //点击数
	Download    int                    `json:"download"`                                      //下载数
	CreatedAt   time.Time              `json:"createdAt"`
	UpdatedAt   time.Time              `json:"updatedAt"`
	Category    string                 `json:"category" gorm:"-" binding:"category"` // 学科分类
	FurtherData map[string]interface{} `gorm:"type:json"`
}

type BookOmitempty struct { //学科图书
	Id        int       `json:"id"`                                            //序号
	Name      string    `json:"name" binding:"omitempty"`                      //图书名称
	Nation    string    `json:"nation" binding:"omitempty"`                    //国家
	Language  string    `json:"language" binding:"omitempty"`                  //语言
	Author    string    `json:"author" binding:"omitempty"`                    //作者
	Publisher string    `json:"publisher" binding:"omitempty"`                 //出版商
	Time      time.Time `json:"time" gorm:"type:datetime" binding:"omitempty"` //出版年月s
	Digest    string    `json:"digest" binding:"omitempty"`                    //摘要
	Text      string    `json:"text"`                                          //全文
	Click     int       `json:"click"`                                         //点击数
	Download  int       `json:"download"`                                      //下载数
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Category  string    `json:"category" gorm:"-" binding:"omitempty,category"` // 学科分类
}

type BookResponseMsg struct {
	Id int `json:"id"`
}

type BookForm struct {
	Total int    `json:"total" binding:"numeric"`
	List  []Book `json:"list" binding:"required,dive"`
}

type BookList struct {
	Total          int       `json:"total"`
	List           []Book    `json:"list"`
	CategoryNumber [10]int64 `json:"categoryNumber"`
}

type BookRes struct {
	Id        int       `json:"id"`                                            //序号
	Title     string    `json:"name" binding:"required"`                       //图书名称
	Nation    string    `json:"nation" binding:"required"`                     //国家
	Language  string    `json:"language" binding:"required"`                   //语言
	Author    string    `json:"author" binding:"required"`                     //作者
	Publisher string    `json:"publisher" binding:"required"`                  //出版商
	Time      time.Time `json:"time"  gorm:"type:datetime" binding:"required"` //出版年月
	Digest    string    `json:"digest" binding:"required"`                     //摘要
	Text      string    `json:"text" binding:"required"`                       //全文
	Click     int       `json:"click"`                                         //点击数
	Download  int       `json:"download"`                                      //下载数
}

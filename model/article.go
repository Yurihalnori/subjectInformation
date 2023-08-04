package model

import "time"

type Article struct { //期刊论文
	Id           int       `json:"id"`                                             //序号
	Nation       string    `json:"nation" binding:"required"`                      //国家
	Title        string    `json:"title" binding:"required"`                       //题名
	Periodical   string    `json:"periodical" binding:"required"`                  //期刊名称
	Author       string    `json:"author" binding:"required"`                      //作者
	Organization string    `json:"organization" binding:"required"`                //作者单位
	CreateDate   time.Time `json:"CreateDate" gorm:"type:date" binding:"required"` //发表时间
	Technique    string    `json:"technique" binding:"required"`                   //研究方法
	KeyWord      string    `json:"keyWord" binding:"required"`                     //关键字
	Digest       string    `json:"digest" binding:"required"`                      //摘要
	Data         string    `json:"data" binding:"required"`                        //数据
	Text         string    `json:"text" binding:"required"`                        //全文
	Click        int       `json:"click"`                                          //点击数
	Download     int       `json:"download"`                                       //下载数
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Category     string    `json:"category" gorm:"-" binding:"required"` // 学科分类
}

type ArticleOmitempty struct { //期刊论文
	Id           int       `json:"id"`                                              //序号
	Nation       string    `json:"nation" binding:"omitempty"`                      //国家
	Title        string    `json:"title" binding:"omitempty"`                       //题名
	Periodical   string    `json:"periodical" binding:"omitempty"`                  //期刊名称
	Author       string    `json:"author" binding:"omitempty"`                      //作者
	Organization string    `json:"organization" binding:"omitempty"`                //作者单位
	CreateDate   time.Time `json:"CreateDate" gorm:"type:date" binding:"omitempty"` //发表时间
	Technique    string    `json:"technique" binding:"omitempty"`                   //研究方法
	KeyWord      string    `json:"keyWord" binding:"omitempty"`                     //关键字
	Digest       string    `json:"digest" binding:"omitempty"`                      //摘要
	Data         string    `json:"data" binding:"omitempty"`                        //数据
	Text         string    `json:"text" binding:"omitempty"`                        //全文
	Click        int       `json:"click"`                                           //点击数
	Download     int       `json:"download"`                                        //下载数
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Category     string    `json:"category" gorm:"-" binding:"omitempty"` // 学科分类
}

type ArticleResponseMsg struct {
	Id int `json:"id"`
}

type ArticleForm struct {
	Total int       `json:"total" binding:"numeric"`
	List  []Article `json:"list" binding:"required,dive"`
}

type ArticleList struct {
	Total int       `json:"total"`
	List  []Article `json:"list"`
	// CategoryNumber []int            `json:"categoryNumber"`
}

type ArticleRes struct { //期刊论文
	Id           int       `json:"id"`                                             //序号
	Nation       string    `json:"nation" binding:"required"`                      //国家
	Title        string    `json:"title" binding:"required"`                       //题名
	Periodical   string    `json:"periodical" binding:"required"`                  //期刊名称
	Author       string    `json:"author" binding:"required"`                      //作者
	Organization string    `json:"organization" binding:"required"`                //作者单位
	CreateDate   time.Time `json:"CreateDate" gorm:"type:date" binding:"required"` //发表时间
	Technique    string    `json:"technique" binding:"required"`                   //研究方法
	KeyWord      string    `json:"keyWord" binding:"required"`                     //关键字
	Digest       string    `json:"digest" binding:"required"`                      //摘要
	Data         string    `json:"data" binding:"required"`                        //数据
	Text         string    `json:"text" binding:"required"`                        //全文
	Click        int       `json:"click"`                                          //点击数
	Download     int       `json:"download"`                                       //下载数
}

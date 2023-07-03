package model

import "time"

type Dissertation struct { //学位论文
	Id         int       `json:"id"`                            //序号
	Title      string    `json:"title" binding:"required"`      //题目
	Author     string    `json:"author" binding:"required"`     //作者
	Tutor      string    `json:"tutor" binding:"required"`      //导师
	Province   string    `json:"province" binding:"required"`   //省(通指一级地方单位)
	City       string    `json:"city" binding:"required"`       //市(通指二级地方单位)
	University string    `json:"university" binding:"required"` //学校
	College    string    `json:"college" binding:"required"`    //学院
	Year       string    `json:"year" binding:"required"`       //年份
	Technique  string    `json:"technique" binding:"required"`  //研究方法
	KeyWord    string    `json:"keyWord" binding:"required"`    //关键字
	Digest     string    `json:"digest" binding:"required"`     //摘要
	Data       string    `json:"data" binding:"required"`       //数据
	Text       string    `json:"text" binding:"required"`       //全文
	Click      int       `json:"click"`                         //点击数
	Download   int       `json:"download"`                      //下载数
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Category   string    `json:"category" gorm:"-" binding:"required"` // 学科分类
}

type DissertationOmitempty struct { //学位论文
	Id         int       `json:"id"`                             //序号
	Title      string    `json:"title" binding:"omitempty"`      //题目
	Author     string    `json:"author" binding:"omitempty"`     //作者
	Tutor      string    `json:"tutor" binding:"omitempty"`      //导师
	Province   string    `json:"province" binding:"omitempty"`   //省(通指一级地方单位)
	City       string    `json:"city" binding:"omitempty"`       //市(通指二级地方单位)
	University string    `json:"university" binding:"omitempty"` //学校
	College    string    `json:"college" binding:"omitempty"`    //学院
	Year       string    `json:"year" binding:"omitempty"`       //年份
	Technique  string    `json:"technique" binding:"omitempty"`  //研究方法
	KeyWord    string    `json:"keyWord" binding:"omitempty"`    //关键字
	Digest     string    `json:"digest" binding:"omitempty"`     //摘要
	Data       string    `json:"data" binding:"omitempty"`       //数据
	Text       string    `json:"text" binding:"omitempty"`       //全文
	Click      int       `json:"click"`                          //点击数
	Download   int       `json:"download"`                       //下载数
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Category   string    `json:"category" gorm:"-" binding:"omitempty"` // 学科分类
}

type DissertationResponseMsg struct {
	Id int `json:"id"`
}

type DissertationForm struct {
	Total int            `json:"total" binding:"numeric"`
	List  []Dissertation `json:"list" binding:"required,dive"`
}

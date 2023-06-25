package model

import "time"

type News struct { //学界咨询
	Id         int       `json:"id"`                                 //序号
	Title      string    `gorm:"type:varchar(64)" json:"title"`      //题目
	Module     uint      `json:"module"`                             //模块  0:行业资讯，1：学术会议，2：学科竞赛，3：招聘信息
	Department string    `gorm:"type:varchar(64)" json:"department"` //主体单位
	Text       string    `json:"content"`                            //全文
	Click      uint      `json:"click" gorm:"default:0"`             //点击数
	Date       string    `gorm:"type:varchar(64)" json:"date"`       // 发布时间
	Region     uint      `json:"region"`                             //domestic|foreign
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Category   string    `json:"category"` // 学科分类
}

type GetSomeNews struct {
	Module   int    `form:"module" json:"module"`
	Page     int    `form:"page" json:"page"`
	Limit    int    `form:"limit" json:"limit"`
	Category string `form:"category" json:"category"`
	Name     string `form:"name" json:"name"`
	Order    string `form:"order" json:"order"`
}

type NewsPreview struct {
	Id         int    `json:"id"`                                 //序号
	Title      string `gorm:"type:varchar(64)" json:"title"`      //题目
	Module     uint   `json:"module"`                             //模块  0:行业资讯，1：学术会议，2：学科竞赛，3：招聘信息
	Department string `gorm:"type:varchar(64)" json:"department"` //主体单位
	Click      uint   `json:"click"`                              //点击数
	Date       string `gorm:"type:varchar(64)" json:"date"`       // 发布时间
	Region     uint   `json:"region"`                             //domestic|foreign
	Category   string `json:"category"`                           // 学科分类
}

type NewsDetail struct {
	Id         int    `json:"id"`                                 //序号
	Title      string `gorm:"type:varchar(64)" json:"title"`      //题目
	Module     uint   `json:"module"`                             //模块  0:行业资讯，1：学术会议，2：学科竞赛，3：招聘信息
	Department string `gorm:"type:varchar(64)" json:"department"` //主体单位
	Click      uint   `json:"click"`                              //点击数
	Text       string `json:"content"`                            //全文
	Date       string `gorm:"type:varchar(64)" json:"date"`       // 发布时间
	Region     uint   `json:"region"`                             //domestic|foreign
	Category   string `json:"category"`                           // 学科分类
}

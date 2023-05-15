package model

import "time"

type News struct { //学界咨询
	Id         int       `json:"id"`                                   //序号
	Title      string    `gorm:"type:varchar(64)" json:"title"`        //题目
	Module     uint      `json:"module"`                               //模块  0:行业资讯，1：学术会议，2：学科竞赛，3：招聘信息
	Department string    `gorm:"type:varchar(64)" json:"organization"` //主体单位
	Text       string    `json:"text"`                                 //全文
	Click      uint      `json:"click"`                                //点击数
	Date       string    `gorm:"type:varchar(64)" json:"time"`         // 发布时间
	Region     uint      `json:"region"`                               //domestic|foreign
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Category   string    `json:"category" gorm:"-:migration"` // 学科分类
}

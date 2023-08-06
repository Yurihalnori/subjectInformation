package model

import "time"

type News struct { //学界咨询
	Id          int                    `json:"id"`                                 //序号
	Title       string                 `gorm:"type:varchar(64)" json:"title"`      //题目
	Module      uint                   `json:"module"`                             //模块  0:行业资讯，1：学术会议，2：学科竞赛，3：招聘信息
	Department  string                 `gorm:"type:varchar(64)" json:"department"` //主体单位
	Text        string                 `gorm:"type:longtext" json:"content"`       //全文
	Click       uint                   `json:"click" gorm:"default:0"`             //点击数
	Date        time.Time              `gorm:"type:datetime" json:"date"`          // 发布时间
	Region      uint                   `json:"region"`                             //domestic|foreign
	CreatedAt   time.Time              `json:"createdAt"`
	UpdatedAt   time.Time              `json:"updatedAt"`
	Category    string                 `json:"category" gorm:"type:varchar(64)" binding:"category"` // 学科分类
	FurtherData map[string]interface{} `gorm:"type:json"`
}

type NewsEditRequest struct { //学界咨询
	Id         int       `json:"id"`                                 //序号
	Title      string    `gorm:"type:varchar(64)" json:"title"`      //题目
	Module     uint      `json:"module"`                             //模块  0:行业资讯，1：学术会议，2：学科竞赛，3：招聘信息
	Department string    `gorm:"type:varchar(64)" json:"department"` //主体单位
	Text       string    `gorm:"type:longtext" json:"content"`       //全文
	Click      uint      `json:"click" gorm:"default:0"`             //点击数
	Date       time.Time `gorm:"type:datetime" json:"date"`          // 发布时间
	Region     uint      `json:"region"`                             //domestic|foreign
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Category   string    `json:"category" gorm:"type:varchar(64)"`
	Blank      string    `json:"blank"`
}

type GetSomeNews struct {
	Module   int    `form:"module" json:"module" binding:"required,oneof=1 2 3 4 5"`
	Page     int    `form:"page" json:"page" binding:"numeric"`
	Limit    int    `form:"limit" json:"limit" binding:"numeric"`
	Category string `form:"category" json:"category" binding:"category" `
	Name     string `form:"name" json:"name"`
	Order    string `form:"order" json:"order"`
}

type NewsForm struct {
	Total    int    `json:"total"`
	NewsList []News `json:"list"`
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
	Id         int       `json:"id"`                                 //序号
	Title      string    `gorm:"type:varchar(64)" json:"title"`      //题目
	Module     uint      `json:"module"`                             //模块  0:行业资讯，1：学术会议，2：学科竞赛，3：招聘信息
	Department string    `gorm:"type:varchar(64)" json:"department"` //主体单位
	Click      uint      `json:"click"`                              //点击数
	Text       string    `json:"content"`                            //全文
	Date       time.Time `gorm:"type:datetime" json:"date"`          // 发布时间
	Region     uint      `json:"region"`                             //domestic|foreign
	Category   string    `json:"category"`                           // 学科分类
}

type NewsSearchRequest struct {
	Content  string `json:"content" binding:"required"`
	Module   uint   `json:"module" binding:"required,oneof=1 2 3 4 5"`
	Name     string `json:"name"`
	Order    string `json:"order"`
	Page     int    `form:"page" json:"page" binding:"numeric"`
	Limit    int    `form:"limit" json:"limit" binding:"numeric"`
	Category string `form:"category" json:"category" binding:"category"`
}

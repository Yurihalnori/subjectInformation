package model

import "time"

type Project struct { //科研项目
	Id             int                    `json:"id"`                                                         //序号
	Title          string                 `json:"title" binding:"required"`                                   //项目名称
	Classification string                 `json:"classification" binding:"required"`                          //项目类别
	Sponsor        string                 `json:"sponsor" binding:"required"`                                 //资助主体
	ApprovalNumber string                 `json:"approvalNumber" binding:"required"`                          //项目批准号
	CreateDate     time.Time              `gorm:"type:datetime" json:"createDate" binding:"omitempty,timing"` //立项时间
	Superintendent string                 `json:"superintendent" binding:"required"`                          //项目负责人
	Organization   string                 `json:"organization" binding:"required"`                            //工作单位
	Region         string                 `json:"region" binding:"required"`                                  //地区
	Click          int                    `json:"click"`                                                      //点击数
	Download       int                    `json:"download"`                                                   //下载数
	Data           int                    `json:"data"`
	CreatedAt      time.Time              `json:"createdAt"`
	UpdatedAt      time.Time              `json:"updatedAt"`
	Category       string                 `json:"category" gorm:"-" binding:"category"` // 学科分类
	FurtherData    map[string]interface{} `gorm:"type:json"`
	Blank          string                 `json:"blank"`
}

type ProjectOmitempty struct {
	Id             int       `json:"id" binding:"omitempty"`
	Title          string    `json:"title" binding:"omitempty"`
	Classification string    `json:"classification" binding:"omitempty"`
	Sponsor        string    `json:"sponsor" binding:"omitempty"`
	ApprovalNumber string    `json:"approvalNumber" binding:"omitempty"`
	CreateDate     time.Time `gorm:"type:datetime" json:"createDate" binding:"omitempty,timing"`
	Superintendent string    `json:"superintendent" binding:"omitempty"`
	Organization   string    `json:"organization" binding:"omitempty"`
	Region         string    `json:"region" binding:"omitempty"`
	Click          int       `json:"click" binding:"omitempty"`
	Download       int       `json:"download" binding:"omitempty"`
	Data           int       `json:"data"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Category       string    `json:"category" gorm:"-" binding:"omitempty,category"`
	Blank          string    `json:"blank"`
}

type ProjectResponseMsg struct {
	Id int `json:"id"`
}

type ProjectForm struct {
	Total int       `json:"total" binding:"numeric"`
	List  []Project `json:"list" binding:"required,dive"`
}

type ProjectList struct {
	Total          int       `json:"total"`
	List           []Project `json:"list"`
	CategoryNumber [10]int64 `json:"categoryNumber"`
}

type ProjectRes struct { //科研项目
	Id             int       `json:"id"`                                                         //序号
	Title          string    `json:"title" binding:"required"`                                   //项目名称
	Classification string    `json:"classification" binding:"required"`                          //项目类别
	Sponsor        string    `json:"sponsor" binding:"required"`                                 //资助主体
	ApprovalNumber string    `json:"approvalNumber" binding:"required"`                          //项目批准号
	CreateDate     time.Time `gorm:"type:datetime" json:"createDate" binding:"omitempty,timing"` //立项时间
	Superintendent string    `json:"superintendent" binding:"required"`                          //项目负责人
	Organization   string    `json:"organization" binding:"required"`                            //工作单位
	Region         string    `json:"region" binding:"required"`                                  //地区
	Click          int       `json:"click"`                                                      //点击数
	Download       int       `json:"download"`                                                   //下载数
	Blank          string    `json:"blank"`
}

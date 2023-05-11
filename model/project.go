package model

import "time"

type Project struct { //科研项目
	Id             int       `json:"id"`             //序号
	Name           string    `json:"name"`           //项目名称
	Classification string    `json:"classification"` //项目类别
	Sponsor        string    `json:"sponsor"`        //资助主体
	ApprovalNumber string    `json:"approvalNumber"` //项目批准号
	CreateDate     string    `json:"createDate"`     //立项时间
	Superintendent string    `json:"superintendent"` //项目负责人
	Organization   string    `json:"organization"`   //工作单位
	Region         string    `json:"region"`         //地区
	Click          int       `json:"click"`          //点击数
	Download       int       `json:"download"`       //下载数
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

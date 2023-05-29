package model

import (
	"time"
)

type Teamwork struct { //团队项目地图
	Id          int       `json:"id"`                                         //序号
	Name        string    `json:"name" binding:"required"`                    //项目名称
	Grade       string    `json:"grade" binding:"oneof=0 1 2 3 4 5,required"` //项目层级 0.国际合作 1.国家级 2.省部级 3.市级 4.区县级 5.乡镇级
	Direction   string    `json:"direction" binding:"required"`               //项目类别 one of=‘crosswise lengthway’ 横向/纵向
	Sponsor     string    `json:"sponsor" binding:"required"`                 //资助主体
	Number      string    `json:"number" binding:"required"`                  //项目批准号
	Time        string    `json:"time" binding:"required"`                    //立项时间
	Principal   string    `json:"principal" binding:"required"`               //项目负责人
	Member      string    `json:"member" binding:"required"`                  //项目成员
	Region      Region    `json:"region" binding:"required,dive"`             //所属地区
	Picture     string    `json:"picture" binding:"required"`                 //调研图片 附带链接
	Achievement string    `json:"achievement" binding:"required"`             //项目成果 附带链接
	Click       int       `json:"click"`                                      //点击数
	Download    int       `json:"download"`                                   //下载数
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Category    string    `json:"category" gorm:"-:migration"` // 学科分类
}

type Region struct {
	Province string `json:"province"` //省
	City     string `json:"city"`     //市
	County   string `json:"county"`   //区
}

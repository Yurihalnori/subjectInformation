package model

import (
	"time"
)

type Teamwork struct { //团队项目地图
	Id          int                    `json:"id"`                                                     //序号
	Name        string                 `json:"name" binding:"required"`                                //项目名称
	Grade       string                 `json:"grade" binding:"oneof=0 1 2 3 4 5,required"`             //项目层级 0.国际合作 1.国家级 2.省部级 3.市级 4.区县级 5.乡镇级
	Direction   string                 `json:"direction" binding:"oneof=horizontal vertical,required"` //项目类别 one of=‘horizontal vertical’ 横向/纵向
	Sponsor     string                 `json:"sponsor" binding:"required"`                             //资助主体
	Number      string                 `json:"number" binding:"required"`                              //项目批准号
	Time        time.Time              `json:"time" binding:"required"`                                //立项时间
	Principal   string                 `json:"principal" binding:"required"`                           //项目负责人
	Member      string                 `json:"member" binding:"required"`                              //项目成员
	Province    string                 `json:"province" binding:"required"`                            //省
	City        string                 `json:"city" binding:"required"`                                //市
	County      string                 `json:"county" binding:"required"`                              //区
	Picture     string                 `json:"picture"`                                                //调研图片 附带链接
	Text        string                 `json:"text"`                                                   //项目成果 附带链接
	Click       int                    `json:"click"`                                                  //点击数
	Download    int                    `json:"download"`                                               //下载数
	CreatedAt   time.Time              `json:"createdAt"`
	UpdatedAt   time.Time              `json:"updatedAt"`
	FurtherData map[string]interface{} `gorm:"type:json"`
	Category    string                 `json:"category" gorm:"-" binding:"category"` // 学科分类
}

type TeamworkOmitempty struct {
	Id        int       `json:"id"`
	Name      string    `json:"name" binding:"omitempty"`
	Grade     string    `json:"grade" binding:"omitempty,oneof=0 1 2 3 4 5"`
	Direction string    `json:"direction" binding:"omitempty,oneof=horizontal vertical"`
	Sponsor   string    `json:"sponsor" binding:"omitempty"`
	Number    string    `json:"number" binding:"omitempty"`
	Time      string    `json:"time" binding:"omitempty"`
	Principal string    `json:"principal" binding:"omitempty"`
	Member    string    `json:"member" binding:"omitempty"`
	Province  string    `json:"province" binding:"omitempty"`
	City      string    `json:"city" binding:"omitempty"`
	County    string    `json:"county" binding:"omitempty"`
	Picture   string    `json:"picture" binding:"omitempty"`
	Text      string    `json:"text"`
	Click     int       `json:"click"`
	Download  int       `json:"download"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Category  string    `json:"category" gorm:"-" binding:"omitempty,category"`
}

type TeamworkResponseMsg struct {
	Id int `json:"id"`
}

type TeamworkForm struct {
	Total int        `json:"total" binding:"numeric"`
	List  []Teamwork `json:"list" binding:"required,dive"`
}

type TeamworkList struct {
	Total          int        `json:"total"`
	List           []Teamwork `json:"list"`
	CategoryNumber [10]int64  `json:"categoryNumber"`
}

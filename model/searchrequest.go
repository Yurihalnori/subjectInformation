package model

import "time"

type SearchCommonDBRequest struct {
	Title    string `json:"title"`
	Category string `json:"category" binding:"category"`
	Module   int    `json:"module" binding:"required,oneof=1 2 3 4 5"`
	Page     int    `json:"page" binding:"numeric"`
	Limit    int    `json:"limit" binding:"numeric"`
	Name     string `json:"name"`
	Order    string `json:"order"`
}

type SearchCommonDBPreview struct {
	Title     string `json:"title"`
	Id        int    `json:"id"`
	TableName string `json:"tableName"`
	Author    string `json:"author"`
	Time      string `json:"time"`
}

type SearchUniqueDBRequest struct {
	Title    string `json:"title" binding:"required"`
	Category string `json:"category" binding:"category"`
	Page     int    `json:"page" binding:"numeric"`
	Limit    int    `json:"limit" binding:"numeric"`
	Name     string `json:"name"`
	Order    string `json:"order"`
}

type SearchUniqueDBPreview struct {
	Id       int    `json:"id"`                          //序号
	Name     string `json:"name" binding:"omitempty"`    //名称
	Trimmer  string `json:"trimmer" binding:"omitempty"` //整理者
	Key_word string `json:"keyWord" binding:"omitempty"` //关键字
	Click    int    `json:"click"`                       //点击数
	Download int    `json:"download"`                    //下载数
}

type SearchTeamworkRequest struct {
	Title    string `json:"title" binding:"required"`
	Category string `json:"category" binding:"category"`
	Page     int    `json:"page" binding:"numeric"`
	Limit    int    `json:"limit" binding:"numeric"`
	Name     string `json:"name"`
	Order    string `json:"order"`
	Grade    int    `json:"grade"`
}

type SearchTeamworkPreview struct { //团队项目地图
	Id        int       `json:"id"`                                                     //序号
	Name      string    `json:"name" binding:"required"`                                //项目名称
	Grade     string    `json:"grade" binding:"oneof=0 1 2 3 4 5 6,required"`           //项目层级 0.国际合作 1.国家级 2.省部级 3.市级 4.区县级 5.乡镇级
	Direction string    `json:"direction" binding:"oneof=horizontal vertical,required"` //项目类别 one of=‘horizontal vertical’ 横向/纵向
	Sponsor   string    `json:"sponsor" binding:"required"`                             //资助主体
	Time      time.Time `json:"time" binding:"required"`                                //立项时间
	Principal string    `json:"principal" binding:"required"`                           //项目负责人
	Province  string    `json:"province" binding:"required"`                            //省
	City      string    `json:"city" binding:"required"`                                //市
	County    string    `json:"county" binding:"required"`                              //区
	Click     int       `json:"click"`                                                  //点击数
	Download  int       `json:"download"`                                               //下载数
}

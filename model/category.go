package model

import "time"

type Category struct { //学科分类
	Id         int       `json:"id"`         //主键
	ForeignKey int       `json:"foreignKey"` //外键
	Tablee     string    `json:"table"`
	Category1  bool      //新闻学
	Category2  bool      //传播学
	Category3  bool      //网络与新媒体
	Category4  bool      //广告学
	Category5  bool      //国际新闻与传播
	Category6  bool      //编辑出版学
	Category7  bool      //数字出版
	Category8  bool      //时尚传播
	Category9  bool      //广播电视学
	Category10 bool      //会展
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

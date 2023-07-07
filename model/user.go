package model

import "time"

type User struct { //用户   手机号，个人资料等后续讨论
	Id        int       `json:"id"` //序号
	Username  string    `gorm:"type:varchar(32);unique" json:"username" binding:"required,min=4,max=32"`
	Password  string    `gorm:"type:varchar(32)" json:"password" binding:"required,min=6,max=32"`
	Usertype  int       `gorm:"default:0" json:"usertype" binding:"oneof=0 1"` // 管理员
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserChangeRequest struct {
	Id       int `json:"userId" binding:"numeric,required"`
	UserType int `json:"role" binding:"required,oneof=0 1"`
}

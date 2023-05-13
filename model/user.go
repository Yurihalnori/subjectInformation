package model

import "time"

type User struct { //用户   手机号，个人资料等后续讨论
	Id        int       `json:"id"` //序号
	Username  string    `gorm:"type:varchar(32);unique" json:"username"`
	Password  string    `gorm:"type:varchar(32)" json:"password"`
	Usertype  int       // 管理员
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

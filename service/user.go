package service

import (
	"subjectInformation/model"
)

func CreateAUser(user *model.User) (err error) {
	err = model.DB.Create(&user).Error
	return err
}

func CheckUsername(username string) bool {
	var user model.User
	res := model.DB.Where("username = ?", username).First(&user)
	if res.RowsAffected == 1 {
		return false
	}
	return true
}
func CheckPassword(username string) (user *model.User, err error) {
	model.DB.Where("username = ?", username).First(&user)
	return
}

func CheckAdmin(number string) (user *model.User) {
	model.DB.Where("id = ?", number).First(&user)
	return
}

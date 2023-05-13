package service

import "subjectInformation/model"

func CreateAUser(user *model.User) (err error) {
	err = model.DB.Create(&user).Error
	return err
}

func CheckPassword(username string) (user *model.User) {
	model.DB.Where("username = ?", username).First(&user)
	return
}

func CheckAdmin(number string) (user *model.User) {
	model.DB.Where("id = ?", number).First(&user)
	return
}

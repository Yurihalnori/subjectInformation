package service

import (
	"subjectInformation/model"
)

type UserService struct {
}

func (UserService) CreateAUser(user *model.User) (err error) {
	err = model.DB.Create(&user).Error
	return err
}

func (UserService) CheckUsername(username string) bool {
	var user model.User
	res := model.DB.Where("username = ?", username).First(&user)
	if res.RowsAffected == 1 {
		return false
	}
	return true
}

func (UserService) CheckPassword(username string) (user *model.User, err error) {
	model.DB.Where("username = ?", username).First(&user)
	return
}

func (UserService) CheckInfo(userid string) (user *model.User, err error) {
	model.DB.Where("id = ?", userid).First(&user)
	return
}

func (UserService) ChangeInfo(form model.UserChangeRequest) (err error) {
	res := model.DB.Save(form)
	err = res.Error
	return
}

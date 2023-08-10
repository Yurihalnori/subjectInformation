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
	return res.RowsAffected != 1
}

func (UserService) CheckPassword(username string) (user *model.User, err error) {
	err = model.DB.Where("username = ?", username).First(&user).Error
	return
}

func (UserService) CheckInfo(userid string) (user *model.User, err error) {
	err = model.DB.Where("id = ?", userid).First(&user).Error
	return
}

func (UserService) ChangeInfo(form model.UserChangeRequest) (err error) {
	var user model.User
	res := model.DB.Model(&user).Where("id=?", form.Id).Update("usertype", form.UserType)
	err = res.Error
	return
}

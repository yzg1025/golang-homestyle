package service

import (
	"gin/global"
	"gin/models"
	"gin/utils"
)

func Login(phone,password string) (err error,user *models.Login) {
	var U models.Login
	password = utils.MD5([]byte(password))
	err = global.HS_DB.Where("phone = ? AND password = ?",phone,password).First(&U).Error
	return err,&U
}
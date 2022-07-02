package service

import (
	"errors"
	"fmt"
	"gin/global"
	"gin/models"
	"gin/utils"

	"gorm.io/gorm"
)

func Login(phone, password string) (err error, user *models.Login) {
	var U models.Login
	password = utils.MD5([]byte(password))
	err = global.HS_DB.Where("phone = ? AND password = ?", phone, password).First(&U).Error
	fmt.Print("3232323", U)
	return err, &U
}

func SaveLog(logs models.Logs) (err error) {
	err = global.HS_DB.Create(&logs).Error
	return err
}

func GetLog() (err error, list interface{}) {
	var L []models.Logs
	err = global.HS_DB.Find(&L).Error
	return err, &L
}

func ChangePassword(login *models.Login, newpassword string) (err error, userInter *models.Login) {
	var u models.Login

	login.Password = utils.MD5([]byte(login.Password))
	fmt.Println("111", login.Password, login.Phone, newpassword)
	err = global.HS_DB.Where("phone = ? AND password = ?", login.Phone, login.Password).First(&u).Update("password", utils.MD5([]byte(newpassword))).Error
	return err, login
}

func QueryCountry() (err error, list interface{}) {
	var S []models.AreaCode
	err = global.HS_DB.Find(&S).Error
	return err, &S
}

func IsExitCountry(codes string) (isExit bool, code interface{}) {
	var S models.AreaCode
	global.HS_DB.Where("code = ?", codes).First(&S)
	if S.ID > 0 {
		return true, &S
	}
	return false, nil
}

func CreateCountry(code models.AreaCode) (err error) {
	err = global.HS_DB.Create(&code).Error
	return err
}

func Register(u models.Login) (err error, userInter models.Login) {
	var user models.Login
	if !errors.Is(global.HS_DB.Where("phone = ?", u.Phone).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	u.Password = utils.MD5([]byte(u.Password))
	err = global.HS_DB.Create(&u).Error
	return err, u
}

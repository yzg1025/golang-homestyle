package service

import (
	"errors"
	"gin/global"
	"gin/models"
	"gin/utils"

	"gorm.io/gorm"
)

type UserService interface {
	Login(phone, password string) (err error, user *models.Login)
	SaveLog(logs []models.Logs) (err error)
	GetLog() (err error, list interface{})
	ChangePassword(login *models.Login, newpassword string) (err error, userInter *models.Login)
	QueryCountry() (err error, list interface{})
	IsExitCountry(codes string) (isExit bool, code interface{})
	Register(u models.Login) (err error, userInter models.Login)
	CreateCountry(code models.AreaCode) (err error)
	UpdateHeader(url, uid string) error
}

type UserServiceCenter struct {
}

func (us *UserServiceCenter) Login(phone, password string) (err error, user *models.Login) {
	var U models.Login
	password = utils.MD5([]byte(password))
	err = global.HS_DB.Where("phone = ? AND password = ?", phone, password).Omit("password").First(&U).Error
	return err, &U
}

func (us *UserServiceCenter) SaveLog(logs []models.Logs) (err error) {
	err = global.HS_DB.Create(&logs).Error
	return err
}

func (us *UserServiceCenter) GetLog() (err error, list interface{}) {
	var L []models.Logs
	err = global.HS_DB.Find(&L).Error
	return err, &L
}

func (us *UserServiceCenter) ChangePassword(login *models.Login, newpassword string) (err error, userInter *models.Login) {
	var u models.Login

	login.Password = utils.MD5([]byte(login.Password))
	err = global.HS_DB.Where("phone = ? AND password = ?", login.Phone, login.Password).First(&u).Update("password", utils.MD5([]byte(newpassword))).Error
	return err, login
}

func (us *UserServiceCenter) QueryCountry() (err error, list interface{}) {
	var S []models.AreaCode
	err = global.HS_DB.Find(&S).Error
	return err, &S
}

func (us *UserServiceCenter) IsExitCountry(codes string) (isExit bool, code interface{}) {
	var S models.AreaCode
	global.HS_DB.Where("code = ?", codes).First(&S)
	if S.ID > 0 {
		return true, &S
	}
	return false, nil
}

func (us *UserServiceCenter) CreateCountry(code models.AreaCode) (err error) {
	err = global.HS_DB.Create(&code).Error
	return err
}

func (us *UserServiceCenter) Register(u models.Login) (err error, userInter models.Login) {
	var user models.Login
	if !errors.Is(global.HS_DB.Where("phone = ?", u.Phone).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	u.Password = utils.MD5([]byte(u.Password))
	err = global.HS_DB.Create(&u).Error
	return err, u
}

func (us UserServiceCenter) UpdateHeader(url, uid string) error {
	return global.HS_DB.Where("uid = ?", uid).First(&models.Login{}).Update("header", url).Error
}

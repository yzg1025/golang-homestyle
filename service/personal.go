package service

import (
	"gin/global"
	"gin/models"
	"gin/utils"
)

func CreateCatePer(name string) (err error) {
	var PC  models.PersonalCateGory
	PC.ID = "H" + utils.CreateCaptcha()
    PC.CategoryName = name
	err = global.HS_DB.Create(&PC).Error
	return err
}

// GetAllCatePer 获取所有个人分类
func GetAllCatePer() (err error,list interface{}) {
	var AllCate []models.PersonalCateGory
	err = global.HS_DB.Find(&AllCate).Error
	return err,&AllCate
}
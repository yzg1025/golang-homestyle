package service

import (
	"errors"
	"fmt"
	"gin/global"
	"gin/models"
	"gorm.io/gorm"
)

func GetBanners() (err error,list interface{}) {
	var B []models.Banner
	err = global.HS_DB.Find(&B).Error
	return err,&B
}

func CreateBanner(b models.Banner) (err error) {
	if !errors.Is(global.HS_DB.Where("url = ? AND redirect_url = ?", b.Url, b.RedirectUrl).First(&models.Banner{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.HS_DB.Create(&b).Error
}

func DeleteB(b models.Banner) (err error) {
	err = global.HS_DB.Delete(b).Error
	return err
}

func UpdateB(b models.Banner) (err error) {
	var oldA models.Banner
	errold := global.HS_DB.Where("id = ?", b.ID).First(&oldA).Error
	if errold != nil {
		return errors.New("查找不到该条记录")
	}
	fmt.Println(oldA.Url == b.Url || oldA.RedirectUrl == b.RedirectUrl)
	if oldA.Url == b.Url || oldA.RedirectUrl == b.RedirectUrl {
		if !errors.Is(global.HS_DB.Where("url = ? AND redirect_url = ?", b.Url, b.RedirectUrl).First(&models.Banner{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同的banner")
		}
	}
	err = global.HS_DB.Save(b).Error
	return err
}
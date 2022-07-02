package service

import (
	"errors"
	"fmt"
	"gin/global"
	"gin/models"
	"gorm.io/gorm"
)

type BannerServiceInterface interface {
	GetBanners() (err error, list interface{})
	CreateBanner(b models.Banner) (err error)
	DeleteB(b models.Banner) (err error)
	UpdateB(b models.Banner) (err error)
}
type DbBanner struct{}

//var BannerService = new(DbBanner)

func (d *DbBanner) GetBanners() (err error, list interface{}) {
	var B []models.Banner
	err = global.HS_DB.Find(&B).Error
	return err, &B
}

func (d *DbBanner) CreateBanner(b models.Banner) (err error) {
	if !errors.Is(global.HS_DB.Where("url = ? AND redirect_url = ?", b.Url, b.RedirectUrl).First(&models.Banner{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同的banner")
	}
	return global.HS_DB.Create(&b).Error
}

func (d *DbBanner) DeleteB(b models.Banner) (err error) {
	err = global.HS_DB.Delete(b).Error
	return err
}

func (d *DbBanner) UpdateB(b models.Banner) (err error) {
	var oldA models.Banner
	errorOld := global.HS_DB.Where("id = ?", b.ID).First(&oldA).Error
	if errorOld != nil {
		return errors.New("查找不到该条记录")
	}
	fmt.Println(oldA.Url == b.Url || oldA.RedirectUrl == b.RedirectUrl)
	err = global.HS_DB.Save(b).Error
	return err
}

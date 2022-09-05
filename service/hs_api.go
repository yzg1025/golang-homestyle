package service

import (
	"errors"
	"fmt"
	"gin/global"
	"gin/models"
)

type ApiService interface {
	CreateApi(api models.SimpleAPI) (err error)
	QueryApiList(page, pageSize int, apiName, apiPath string) ([]models.SimpleAPI, error)
}

type Api struct{}

func CreateApi(api models.SimpleAPI) (err error) {
	var ap models.SimpleAPI
	global.HS_DB.Where("apiPath = ?", api.ApiPath).First(&ap)
	if api.ApiPath == ap.ApiPath {
		return errors.New("此api已存在")
	}
	err = global.HS_DB.Create(&api).Error
	return
}

func QueryApiList(page, pageSize int, apiName, apiPath string) (interface{}, error) {
	var total int64
	limit := pageSize
	offset := pageSize * (page - 1)
	var apiList []models.SimpleAPI
	tx := global.HS_DB.Model(&apiList)
	if apiName != "" {
		tx.Where("apiName = ?", apiName)
	}
	if apiPath != "" {
		tx.Where("apiPath = ?", apiPath)
	}
	tx.Count(&total)
	err := tx.Limit(limit).Offset(offset).Preload("SimpleApiRequestField").Preload("SimpleApiResponseField").Find(&apiList).Error
	result := map[string]interface{}{
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
		"data":     apiList,
	}
	return result, err
}

func RemoveApi(id string) error {
	var api models.SimpleAPI
	err := global.HS_DB.Delete(&api, "api_id = ?", id).Error
	return err
}

func QueryApiById(id string) (api models.SimpleAPI, err error) {
	err = global.HS_DB.Where("api_id = ?", id).Preload("SimpleApiRequestField").Preload("SimpleApiResponseField").First(&api).Error
	return api, err
}

func InsertData(model interface{}, tableName string) (err error) {
	name := fmt.Sprintf("auto_%s", tableName)
	return global.HS_DB.Table(name).Create(model).Error
}

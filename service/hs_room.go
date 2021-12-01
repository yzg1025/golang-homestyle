package service

import (
	"errors"
	"gin/global"
	"gin/models"

	"gorm.io/gorm"
)

// 查询 roomcount
func QueryRoomCount(r *models.RoomCountParams) (data interface{}, err error) {
	var rc []models.RoomCount
	sqlstr := "longitude >= ? AND longitude < ? and latitude >= ? AND latitude < ? AND zoom =? "
	err = global.HS_DB.Where(sqlstr, r.MinLng, r.MaxLng, r.MinLat, r.MaxLat, r.Zoom).Find(&rc).Error

	return rc, err
}

// 添加 roomcount
func AddRoomCount(r models.RoomCount) (err error) {
	err = global.HS_DB.Where("longitude = ? AND latitude = ? AND code = ? ", r.Longitude, r.Latitude, r.Code).First(&r).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("存在相同的数据")
	}
	return global.HS_DB.Create(&r).Error

}

// 查询 roomlist
func QueryRoomList(r models.RoomCountParams) (data interface{}, total int64, err error) {
	var rc []models.RoomList
	limit := r.Pagination.PageSize
	offset := r.Pagination.PageSize * (r.Pagination.PageNo - 1)

	db := global.HS_DB.Model(&rc)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if r.MinLng != "" && r.MaxLng != "" {
		sqlstr := "longitude >= ? AND longitude < ? "
		db.Where(sqlstr, r.MinLng, r.MaxLng)
	}

	// if r.MinLat != "" && r.MaxLat != "" {
	// 	db.Where("latitude >= ? AND latitude < ?", r.MinLat, r.MaxLat)
	// }

	if r.Pagination.PageNo > 0 && r.Pagination.PageSize > 0 {
		db.Limit(limit).Offset(offset)
	}

	err = db.Preload("Location").Preload("Tags").Find(&rc).Error

	return rc, total, err
}

// 逮主键删除
func RemoveList(r models.RoomList) (err error) {
	errtag := global.HS_DB.Where("id = ?", r.ID).Delete(&models.Tag{}).Error
	errlocal := global.HS_DB.Where("id = ?", r.ID).Delete(&models.Location{}).Error

	if errtag == nil && errlocal == nil {
		return global.HS_DB.Delete(&r).Error
	}

	return nil
}

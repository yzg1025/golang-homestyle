package service

import (
	"gin/global"
	"gin/models"
	"strconv"
	"strings"
)

type Albums struct{}

func GetRankAlbum(am models.RankAlbumParams) (albumRankPageList interface{}, err error) {
	var ra []models.RankResult
	var a []models.Albums

	limit := am.Pagination.PageSize
	offset := am.Pagination.PageSize * (am.Pagination.PageNo - 1)

	strSplit := strings.Split(am.RankIds, ",")
	intIds := make([]int, 0, len(strSplit))

	for _, v := range strSplit {
		id, _ := strconv.Atoi(v)
		intIds = append(intIds, int(id))
	}

	err = global.HS_DB.Where("rankId in (?)", intIds).Limit(limit).Offset(offset).Find(&ra).Error

	apl := make([]models.AlbumRankPageList, len(strSplit))
	if err == nil {
		for i, v := range ra {
			strRaIds := strings.Split(v.Ids, ",")
			err1 := global.HS_DB.Where("id in (?)", strRaIds).Find(&a).Error
			apl[i].RankResult = ra[i]
			apl[i].Albums = a
			if err1 != nil {
				return nil, err1
			}
		}
	}

	data := map[string]interface{}{
		"albumRankPageList": apl,
		"rankIds":           intIds,
		"pageNum":           limit,
		"pageSize":          offset,
	}

	return data, nil
}

func GetGuessLike() (data interface{}, err error) {
	var ral []models.RecommendAnchorList
	var ril []models.RecommendInfoList

	err1 := global.HS_DB.Model(&ral).Find(&ral).Error

	if err1 != nil {
		return nil, err1
	}

	err = global.HS_DB.Model(&ril).Find(&ril).Error

	data = map[string]interface{}{
		"count":               len(ral) + len(ril),
		"recommendAnchorList": ral,
		"recommendInfoList":   ril,
	}

	return data, err
}

// 查询分类
func GetCategory() (data interface{}, err error) {
	var cate []models.CategoryAll
	err = global.HS_DB.Preload("Categories.Subcategories").Find(&cate).Error
	return cate, err
}

// 获取 radio
func GetRadio(page models.Pagination) (data interface{}, err error) {
	var radio []models.Radio
	var total int64

	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	err = global.HS_DB.Model(&radio).Count(&total).Error
	if err != nil {
		return
	}

	err = global.HS_DB.Limit(limit).Offset(offset).Find(&radio).Error

	data = map[string]interface{}{
		"total": total,
		"radio": radio,
	}

	return data, err
}

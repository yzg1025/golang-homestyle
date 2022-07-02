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

// GetRadio 获取 radio
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

// SimpleAlbum 获取单个专辑
func SimpleAlbum(id int) (list interface{}, err error) {
	var simple models.Albumsimple
	err = global.HS_DB.Where("albumId = ? ", id).First(&simple).Error
	list = map[string]interface{}{
		"albumId":         id,
		"albumsimpleInfo": simple,
	}
	return list, err
}

func GetTracksList(params models.AlbumIdParams) (data interface{}, err error) {
	var list []models.TracksAlbum
	var list1 []models.TracksAlbum
	var total int64
	limit := params.Pagination.PageSize
	offset := params.Pagination.PageSize * (params.Pagination.PageNo - 1)

	global.HS_DB.Where("albumId = ? ", params.AlbumId).Find(&list1).Count(&total)

	err = global.HS_DB.Where("albumId = ? ", params.AlbumId).Limit(limit).Offset(offset).Find(&list).Order("index desc").Error

	data = map[string]interface{}{
		"albumId":    params.AlbumId,
		"pageNo":     params.Pagination.PageNo,
		"pageSize":   params.Pagination.PageSize,
		"tracks":     list,
		"totalCount": total,
	}

	return data, err
}

func GetTabList() (data interface{}, err error) {
	var tabs []models.TabList

	err = global.HS_DB.Find(&tabs).Error

	return tabs, err
}

func GetTabViewList() (data interface{}, err error) {
	var tabViews []models.TabViewList

	err = global.HS_DB.Find(&tabViews).Error

	return tabViews, err
}

func HotWordCategory() (data interface{}, err error) {
	var hotcate []models.HotWordBillboardCategory
	err = global.HS_DB.Find(&hotcate).Error

	return hotcate, err
}

func GetHotWordBillboard(id string) (data interface{}, err error) {
	var hotrod []models.HotWordBillboard
	err = global.HS_DB.Where("categoryInt = ?", id).Find(&hotrod).Error

	return hotrod, err
}

func GetBasicAnchorInfo(uid string) (data interface{}, err error) {
	var info models.AnchorBasicInfo
	err = global.HS_DB.Where("uid = ?", uid).First(&info).Error
	return info, err
}

func FollowAnchorByUid(uid string, follow string) error {
	var info models.AnchorBasicInfo
	var nofollow bool = false
	if follow == "0" {
		nofollow = false
	} else {
		nofollow = true
	}
	return global.HS_DB.Where("uid = ?", uid).First(&info).Update("isFollow", nofollow).Error
}

func AlbumComments(params models.AlbumIdParams) (data interface{}, err error) {
	var commont []models.AlbumComments
	var score models.ScoreDiagram
	var total int64
	limit := params.Pagination.PageSize
	offset := params.Pagination.PageSize * (params.Pagination.PageNo - 1)

	global.HS_DB.Where("albumId = ? ", params.AlbumId).First(&score)

	global.HS_DB.Where("albumId = ? ", params.AlbumId).Find(&commont).Count(&total)

	err = global.HS_DB.Where("albumId = ? ", params.AlbumId).Limit(limit).Offset(offset).Find(&commont).Order("likes desc").Error

	list := map[string]interface{}{
		"lists": commont,
	}
	data = map[string]interface{}{
		"pageNo":     params.Pagination.PageNo,
		"pageSize":   params.Pagination.PageSize,
		"commont":    list,
		"score":      score,
		"totalCount": total,
	}
	return data, err
}

func GetAlbumAnchor(uid string) (data interface{}, err error) {
	ids, _ := strconv.Atoi(uid)
	var albums []models.Anchoralbums
	err = global.HS_DB.Where("anchorUid = ?", ids).Find(&albums).Error
	return albums, err
}

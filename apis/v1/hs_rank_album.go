package v1

import (
	"fmt"
	"gin/models"
	"gin/service"
	"gin/utils"

	"github.com/gin-gonic/gin"
)

type Albums struct{}

func GetRankAlbum(c *gin.Context) {
	var am models.RankAlbumParams
	_ = c.ShouldBind(&am)
	if err := utils.Verify(am, utils.RankAlbumVer); err != nil {
		utils.FailWithDetailed(err.Error(), "", c)
		return
	}

	data, err := service.GetRankAlbum(am)
	if err != nil {
		utils.FailMag(err.Error(), c)
		return
	}
	utils.SuccessData(data, c)
}

func GetGuessLike(c *gin.Context) {
	data, err := service.GetGuessLike()
	if err != nil {
		utils.FailMag(err.Error(), c)
		return
	}

	utils.SendDetail(data, "请求成功", c)
}

func GetCategory(c *gin.Context) {
	data, err := service.GetCategory()
	if err != nil {
		utils.FailMag(err.Error(), c)
		return
	}

	utils.SuccessData(data, c)
}

func GetRadio(c *gin.Context) {
	var page models.Pagination
	_ = c.ShouldBind(&page)
	if err := utils.Verify(page, utils.RankAlbumVer); err != nil {
		utils.FailWithDetailed(err.Error(), "", c)
		return
	}
	data, err := service.GetRadio(page)

	if err != nil {
		utils.FailMag(err.Error(), c)
		return
	}

	utils.SendDetail(data, "请求成功", c)
}

// 获取单一专辑
func SimpleAlbum(c *gin.Context) {
	var albumId models.AlbumIdParams
	_ = c.ShouldBind(&albumId)
	if err := utils.Verify(albumId, utils.AlbumIdVar); err != nil {
		utils.FailWithDetailed(err.Error(), "", c)
		return
	}
	data, err := service.SimpleAlbum(albumId.AlbumId)
	if err != nil {
		utils.FailMag("查询失败", c)
		return
	}

	utils.SuccessData(data, c)
}

// 获取专辑播放列表
func GetTracksList(c *gin.Context) {
	var album models.AlbumIdParams
	_ = c.ShouldBind(&album)
	if err := utils.Verify(album, utils.AlbumIdVar); err != nil {
		utils.FailWithDetailed(err.Error(), "", c)
		return
	}

	data, err := service.GetTracksList(album)
	if err != nil {
		utils.FailMag("查询失败", c)
		return
	}
	utils.SuccessData(data, c)
}

// 获取tab tag
func GetTabList(c *gin.Context) {
	data, err := service.GetTabList()
	if err != nil {
		utils.FailMag(err.Error(), c)
		return
	}
	utils.SendDetail(data, "请求成功", c)
}

// TabView
func GetTabViewList(c *gin.Context) {
	data, err := service.GetTabViewList()
	if err != nil {
		utils.FailMag(err.Error(), c)
		return
	}
	utils.SendDetail(data, "请求成功", c)
}

// HotWordCategory

func HotWordCategory(c *gin.Context) {
	data, err := service.HotWordCategory()
	if err != nil {
		utils.FailMag(err.Error(), c)
		return
	}
	utils.SendDetail(data, "请求成功", c)
}

// GetHotWordBillboard
func GetHotWordBillboard(c *gin.Context) {
	var id models.HotWordBillboardParams
	_ = c.ShouldBind(&id)
	if err := utils.Verify(id, utils.HotWordBillboardVar); err != nil {
		utils.FailWithDetailed(err.Error(), "", c)
		return
	}
	data, err := service.GetHotWordBillboard(id.CategoryId)
	if err != nil {
		utils.FailMag(err.Error(), c)
		return
	}
	utils.SendDetail(data, "请求成功", c)
}

// 获取主播信息
func GetBasicAnchorInfo(c *gin.Context) {
	var uid models.AnchorBasicInfoParams
	_ = c.ShouldBind(&uid)
	if err := utils.Verify(uid, utils.BasicInfoVar); err != nil {
		utils.FailWithDetailed(err.Error(), "", c)
		return
	}
	data, err := service.GetBasicAnchorInfo(uid.Uid)
	if err != nil {
		utils.FailMag(err.Error(), c)
		return
	}
	utils.SendDetail(data, "请求成功", c)
}

// 关注主播
func FollowAnchor(c *gin.Context) {
	var uid models.AnchorFollowParams
	_ = c.ShouldBind(&uid)
	if err := utils.Verify(uid, utils.AnchorFollowVar); err != nil {
		utils.FailWithDetailed(err.Error(), "", c)
		return
	}
	err := service.FollowAnchorByUid(uid.Uid, uid.Follow)
	if err != nil {
		utils.FailMag(err.Error(), c)
		return
	}
	utils.SuccessMsg("关注成功", c)
}

// 专辑评论
func AlbumComments(c *gin.Context) {
	var album models.AlbumIdParams
	_ = c.ShouldBind(&album)
	if err := utils.Verify(album, utils.AlbumIdVar); err != nil {
		utils.FailWithDetailed(err.Error(), "", c)
		return
	}
	data, err := service.AlbumComments(album)
	if err != nil {
		utils.FailMag(err.Error(), c)
		return
	}
	utils.SendDetail(data, "请求成功", c)
}

func GetAlbumAnchor(c *gin.Context) {
	var uid models.AnchorBasicInfoParams
	fmt.Print("uid", uid)
	_ = c.ShouldBind(&uid)
	if err := utils.Verify(uid, utils.BasicInfoVar); err != nil {
		utils.FailWithDetailed(err.Error(), "", c)
		return
	}
	data, err := service.GetAlbumAnchor(uid.Uid)
	if err != nil {
		utils.FailMag(err.Error(), c)
		return
	}
	utils.SendDetail(data, "请求成功", c)
}

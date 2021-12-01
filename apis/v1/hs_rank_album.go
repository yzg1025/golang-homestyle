package v1

import (
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

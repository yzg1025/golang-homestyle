package v1

import (
	"gin/global"
	"gin/models"
	"gin/service"
	"gin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 查询
func RoomCount(c *gin.Context) {
	var R models.RoomCount
	var Rq models.RoomCountParams
	_ = c.ShouldBindJSON(&R)
	_ = c.ShouldBind(&Rq)

	if err := utils.Verify(Rq, utils.RoomCountVer); err != nil {

		utils.FailWithDetailed(err.Error(), "", c)
		return
	}

	data, err := service.QueryRoomCount(&Rq)
	if err != nil {
		utils.FailMag(err.Error(), c)
		return
	}

	utils.SuccessData(data, c)
}

// 添加
func AddRoomCount(c *gin.Context) {
	var rc models.RoomCount
	_ = c.ShouldBindJSON(&rc)
	if err := utils.Verify(rc, utils.RoomCountAddVer); err != nil {
		utils.FailWithDetailed(err.Error(), "", c)
		return
	}
	err := service.AddRoomCount(rc)
	if err != nil {
		utils.FailWithDetailed(err.Error(), "", c)
		return
	}
	utils.SuccessMsg("添加成功", c)
}

func QueryList(c *gin.Context) {
	var rc models.RoomCountParams
	_ = c.ShouldBind(&rc)
	if err := utils.Verify(rc, utils.RoomList); err != nil {
		utils.FailWithDetailed(err.Error(), "", c)
		return
	}
	data, total, err := service.QueryRoomList(rc)
	if err != nil {
		utils.FailMag("查询失败", c)
		return
	}

	list := map[string]interface{}{
		"room":     data,
		"total":    total,
		"pageNo":   rc.Pagination.PageNo,
		"pageSize": rc.Pagination.PageSize,
	}

	utils.SendDetail(list, "查询成功", c)
}

func RemoveList(c *gin.Context) {
	var r models.RoomList
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.IDVer); err != nil {
		utils.FailWithDetailed(err.Error(), "", c)
		return
	}
	err := service.RemoveList(r)
	if err != nil {
		global.HS_LOG.Error("删除失败!", zap.Any("err", err))
		utils.FailMag(err.Error(), c)
		return
	}
	utils.SuccessMsg("删除成功", c)
}

package v1

import (
	"gin/global"
	"gin/models"
	"gin/service"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func BannerC(c *gin.Context)  {
	var B models.Banner
	_ = c.ShouldBindJSON(&B)
	if err := utils.Verify(B,utils.BannerVer);err != nil{
		global.HS_LOG.Error("banner 参数不为空",zap.Any("err",err))
		utils.FailMag(err.Error(),c)
		return
	}
	if err := service.CreateBanner(B); err != nil {
		global.HS_LOG.Error("创建失败!", zap.Any("err", err))
		utils.FailMag("创建失败", c)
		return
	}
	utils.SuccessMsg("创建成功",c)
}

func Banners(c *gin.Context)  {
	err,list := service.GetBanners()
	if err != nil {
		global.HS_LOG.Error("获取失败",zap.Any("err",err))
		utils.FailMag("获取失败",c)
	}
	utils.SuccessData(list,c)
}

func BannerD(c *gin.Context)  {
	var B models.Banner
	_ = c.ShouldBindJSON(&B)
	if err := utils.Verify(B,utils.IDVer);err != nil{
		global.HS_LOG.Error("banner id 不为空",zap.Any("err",err))
		utils.FailMag(err.Error(),c)
		return
	}
	if err := service.DeleteB(B); err != nil {
		global.HS_LOG.Error("删除失败!", zap.Any("err", err))
		utils.FailMag("删除失败", c)
		return
	}
	utils.SuccessMsg("删除成功", c)
}

func BannerE(c *gin.Context)  {
	var B models.Banner
	_ = c.ShouldBindJSON(&B)
	if err := utils.Verify(B,utils.BannerVer);err != nil{
		global.HS_LOG.Error("banner 参数不为空",zap.Any("err",err))
		utils.FailMag(err.Error(),c)
		return
	}
	if err := service.UpdateB(B); err != nil {
		global.HS_LOG.Error("修改失败!", zap.Any("err", err))
		utils.FailMag("修改失败", c)
		return
	}
	utils.SuccessMsg("修改成功", c)
}

package v1

import (
	"gin/global"
	"gin/models"
	"gin/service"
	"gin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BannerSer struct {
	serv service.DbBanner
}

var Banner = new(BannerSer)

type BannerInter interface {
	BannerC(c *gin.Context)
	Banners(c *gin.Context)
	BannerD(c *gin.Context)
	BannerE(c *gin.Context)
}

func NewBannerInter() *BannerSer {
	return &BannerSer{}
}

func (b *BannerSer) BannerC(c *gin.Context) {
	var B models.Banner
	_ = c.ShouldBindJSON(&B)
	if err := utils.Verify(B, utils.BannerVer); err != nil {
		global.HS_LOG.Error("banner 参数不为空", zap.Any("err", err), zap.Any("method", "BannerC"))
		utils.FailMag(err.Error(), c)
		return
	}
	if err := b.serv.CreateBanner(B); err != nil {
		global.HS_LOG.Error("创建失败!", zap.Any("err", err))
		utils.FailMag("创建失败", c)
		return
	}
	utils.SuccessMsg("创建成功", c)
}

func (b *BannerSer) Banners(c *gin.Context) {
	err, list := b.serv.GetBanners()
	if err != nil {
		global.HS_LOG.Error("获取失败", zap.Any("err", err))
		utils.FailMag("获取失败", c)
	}
	utils.SuccessData(list, c)
}

func (b *BannerSer) BannerD(c *gin.Context) {
	var B models.Banner
	_ = c.ShouldBindJSON(&B)
	if err := utils.Verify(B, utils.IDVer); err != nil {
		global.HS_LOG.Error("banner id 不为空", zap.Any("err", err))
		utils.FailMag(err.Error(), c)
		return
	}
	if err := b.serv.DeleteB(B); err != nil {
		global.HS_LOG.Error("删除失败!", zap.Any("err", err))
		utils.FailMag("删除失败", c)
		return
	}
	utils.SuccessMsg("删除成功", c)
}

func (b *BannerSer) BannerE(c *gin.Context) {
	var B models.Banner
	_ = c.ShouldBindJSON(&B)
	if err := utils.Verify(B, utils.BannerVer); err != nil {
		global.HS_LOG.Error("banner 参数不为空", zap.Any("err", err))
		utils.FailMag(err.Error(), c)
		return
	}
	if err := b.serv.UpdateB(B); err != nil {
		global.HS_LOG.Error("修改失败!", zap.Any("err", err))
		utils.FailMag("修改失败", c)
		return
	}
	utils.SuccessMsg("修改成功", c)
}

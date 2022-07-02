package router

import (
	v1 "gin/apis/v1"
	"github.com/gin-gonic/gin"
)

func BaseRouterBanner(Router *gin.RouterGroup) (R gin.IRoutes) {

	BR := Router.Group("hs_base")
	//var banner = v1.Banner
	banner := v1.NewBannerInter()
	{
		BR.POST("banners_create", banner.BannerC)
		BR.GET("banners_list", banner.Banners)
		BR.POST("banners_delete", banner.BannerD)
		BR.POST("banners_edit", banner.BannerE)
	}
	return BR
}

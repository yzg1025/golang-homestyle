package router

import (
	v1 "gin/apis/v1"
	"github.com/gin-gonic/gin"
)

func BaseRouterBanner(Router *gin.RouterGroup) (R gin.IRoutes) {
	BR := Router.Group("hs_base")
	{
		BR.POST("banners_create", v1.BannerC)
		BR.GET("banners_list", v1.Banners)
		BR.POST("banners_delete", v1.BannerD)
		BR.POST("banners_edit", v1.BannerE)
	}
	return BR
}

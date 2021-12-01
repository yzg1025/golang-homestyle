package router

import (
	v1 "gin/apis/v1"

	"github.com/gin-gonic/gin"
)

func BaseRouterRankAlbum(Router *gin.RouterGroup) gin.IRoutes {
	BR := Router.Group("hs_base")
	{
		BR.GET("getRankAlbum", v1.GetRankAlbum)
		BR.GET("getGuessLike", v1.GetGuessLike)
		BR.GET("getCategory", v1.GetCategory)
		BR.GET("getRadio", v1.GetRadio)
	}
	return BR
}

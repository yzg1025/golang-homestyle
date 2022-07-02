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
		BR.POST("simpleAlbum", v1.SimpleAlbum)
		BR.POST("getTracksList", v1.GetTracksList)
		BR.POST("tablist", v1.GetTabList)
		BR.POST("tabViewList", v1.GetTabViewList)
		BR.GET("hotWordBillboardCategory", v1.HotWordCategory)
		BR.GET("hotWordBillboard", v1.GetHotWordBillboard)
		BR.POST("basicAnchorInfo", v1.GetBasicAnchorInfo)
		BR.POST("followAnchor", v1.FollowAnchor)
		BR.POST("album/comment", v1.AlbumComments)
		BR.GET("albumAnchor", v1.GetAlbumAnchor)
	}
	return BR
}

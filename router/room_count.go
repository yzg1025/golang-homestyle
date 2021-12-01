package router

import (
	v1 "gin/apis/v1"

	"github.com/gin-gonic/gin"
)

func BaseRouterRoomCount(Router *gin.RouterGroup) gin.IRoutes {
	BR := Router.Group("room")
	{
		BR.GET("count", v1.RoomCount)
		BR.POST("count_add", v1.AddRoomCount)
		BR.GET("list", v1.QueryList)
		BR.POST("remove_list", v1.RemoveList)
	}
	return BR
}

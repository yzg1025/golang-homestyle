package router

import (
	v1 "gin/apis/v1"
	"github.com/gin-gonic/gin"
)

func BaseRouterOther(Router *gin.RouterGroup) (R gin.IRoutes) {
	BR := Router.Group("hs_base")
	var log = v1.NewLogs()
	{
		BR.POST("save_log", log.SaveLogs)
		BR.GET("get_log", log.GetLogs)
	}
	return BR
}

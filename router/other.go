package router

import (
	v1 "gin/apis/v1"
	"github.com/gin-gonic/gin"
)

func BaseRouterOther(Router *gin.RouterGroup) (R gin.IRoutes)  {
	BR := Router.Group("hs_base")
	{
		BR.POST("save_log",v1.SaveLogs)
		BR.GET("get_log",v1.GetLogs)
	}
	return BR
}

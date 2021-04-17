package router

import (
	v1 "gin/apis/v1"
	"github.com/gin-gonic/gin"
)

func BaseRouter(Router *gin.RouterGroup) (R gin.IRoutes)  {
	BR := Router.Group("hs_base")
	{
		BR.POST("login",v1.Login)
		BR.POST("register",v1.Register)
		BR.POST("save_log",v1.SaveLogs)
		BR.POST("send_code",v1.SendMsgCode)
		BR.GET("select_country",v1.SelectCode)
		BR.POST("create_country",v1.CreateCode)
	}
	return BR
}

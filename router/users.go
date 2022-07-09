package router

import (
	v1 "gin/apis/v1"
	"github.com/gin-gonic/gin"
)

func BaseRouterUser(Router *gin.RouterGroup) (R gin.IRoutes) {
	BR := Router.Group("hs_base")
	user := v1.NewUser()
	{
		BR.POST("login", user.Login)
		BR.POST("register", user.Register)
		BR.POST("send_code", user.SendMsgCode)

		BR.GET("select_country", user.SelectCode)
		BR.POST("create_country", user.CreateCode)
	}
	return BR
}

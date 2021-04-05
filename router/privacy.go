package router

import (
	v1 "gin/apis/v1"
	"github.com/gin-gonic/gin"
)

func PrivacyRouter(Router *gin.RouterGroup) (R gin.IRoutes)  {
	PR := Router.Group("hs_base")
	{
		PR.GET("test",v1.Test)
	}
	return PR
}

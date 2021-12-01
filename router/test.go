package router

import (
	v1 "gin/apis/v1"

	"github.com/gin-gonic/gin"
)

func BaseRouterTest(Router *gin.RouterGroup) (R gin.IRoutes) {
	BR := Router.Group("hs_base")
	{
		BR.GET("test", v1.Test)
	}
	return BR
}

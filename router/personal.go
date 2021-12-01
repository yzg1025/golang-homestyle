package router

import (
	v1 "gin/apis/v1"

	"github.com/gin-gonic/gin"
)

func BaseRouterPersonal(Router *gin.RouterGroup) (R gin.IRoutes) {
	BR := Router.Group("hs_base")
	{
		BR.GET("personal_category_get", v1.GetPersonalCategory)
		BR.POST("personal_category_create", v1.CreatePersonalCategory)
	}
	return BR
}

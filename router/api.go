package router

import (
	v1 "gin/apis/v1"
	"github.com/gin-gonic/gin"
)

func BaseRouterAPI(Router *gin.RouterGroup) (R gin.IRoutes) {
	BR := Router.Group("hs_base")
	var api = v1.NewAPIInter()
	{
		BR.POST("create_api", api.CreateApi)
		BR.POST("query_api", api.QueryApi)
		BR.POST("remove_api", api.RemoveApi)
		BR.GET("query_api_id", api.QueryApiById)
		BR.POST("insertData", api.InsertData)
	}
	return BR
}

package initialization

import (
	"gin/generate_code/auto_router"
	"gin/middleware"
	"gin/router"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.Cors())
	BaseGroup := Router.Group("api")
	{
		router.BaseRouterUser(BaseGroup)
		router.BaseRouterBanner(BaseGroup)
		router.BaseRouterOther(BaseGroup)
		router.BaseRouterRankAlbum(BaseGroup)
		// auto_code #
		auto_router.BaseRouterSearchList(BaseGroup)

	}
	PHSGroup := Router.Group("api")
	PHSGroup.Use(middleware.JWTAuth())
	{
		router.PrivacyRouter(PHSGroup)
		router.BaseRouterAPI(PHSGroup)
		// auto_code &
	}
	Router.NoRoute(func(c *gin.Context) { utils.FailMag("你访问的路由不存在", c) })
	return Router
}

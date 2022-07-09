package initialization

import (
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
		router.BaseRouterPersonal(BaseGroup)
		router.BaseRouterRoomCount(BaseGroup)
		router.BaseRouterRankAlbum(BaseGroup)
	}
	PHSGroup := Router.Group("api")
	PHSGroup.Use(middleware.JWTAuth())
	{
		router.PrivacyRouter(PHSGroup)
	}
	Router.NoRoute(func(c *gin.Context) { utils.FailMag("你访问的路由不存在", c) })
	return Router
}

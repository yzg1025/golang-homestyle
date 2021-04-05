package initialization

import (
	"fmt"
	"gin/middleware"
	"gin/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.Cors())

	BaseGroup := Router.Group("api")
	{
       router.BaseRouter(BaseGroup)
	}
	fmt.Println("路由注册成功")
	return Router
}
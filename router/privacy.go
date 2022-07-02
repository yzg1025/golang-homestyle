package router

import (
	v1 "gin/apis/v1"
	"github.com/gin-gonic/gin"
)

func PrivacyRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	PR := Router.Group("hs_privacy")
	var privacy = v1.UploadFilePart
	{
		PR.POST("upload", privacy.UploadFile)
		PR.POST("change_password", v1.ChangePassword)
	}
	return PR
}

package router

import (
	v1 "gin/apis/v1"
	"github.com/gin-gonic/gin"
)

func PrivacyRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	PR := Router.Group("hs_privacy")
	var privacy = v1.UploadFilePart
	var user = v1.NewUser()
	{
		PR.POST("upload", privacy.UploadFile)
		PR.POST("uploadFiles", privacy.UploadBigFile)
		PR.POST("change_password", user.ChangePassword)
		PR.POST("updateHeader", user.UpdateUserHeader)
	}
	return PR
}

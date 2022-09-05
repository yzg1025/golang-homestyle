package auto_router

import (
	"gin/generate_code/auto_api/{{ .ApiPackageName}}"
	"github.com/gin-gonic/gin"
)

func {{ if .Api.AccessAuth }}PrivacyRouter{{.Api.ApiName}}{{ else }}BaseRouter{{.Api.ApiName}}{{ end }}(Router *gin.RouterGroup) (R gin.IRoutes) {
	BR := Router.Group("hs_base")
	{
		BR.{{ .Api.ApiMethod}}("{{.Api.ApiMethod}}", {{ .ApiPackageName}}.Auto{{ .Api.ApiMethod}})
	}
	return BR
}

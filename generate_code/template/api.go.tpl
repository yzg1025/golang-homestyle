package {{ .ApiPackageName}}

import (
  	"gin/generate_code/auto_model"
  	"gin/generate_code/auto_service/{{.ApiPackageName}}"
  	"gin/global"
  	"gin/utils"
  	"github.com/gin-gonic/gin"
  	"go.uber.org/zap"
)

// Auto{{.Api.ApiMethod}} 请求 方法{{ .Api.ApiMethod}}
func Auto{{.Api.ApiMethod}}(c *gin.Context) {
	var autoModelReq auto_model.{{.Api.ApiName}}Req
	_ = c.ShouldBindJSON(&autoModelReq)
	{{- range .autoModelReq}}
        {{ if .FieldRequire eq false}}
            global.HS_LOG.Error("{{ .FieldName}} 参数不为空", zap.Any("err", err), zap.Any("method", "Auto{{.Api.ApiMethod}}"))
       	    utils.FailMag(err.Error(), c)
            return
        {{ end }}
	{{- end}}
	var autoModelRes auto_model.{{.Api.ApiName}}
	if err := {{.ApiPackageName}}.Auto{{ .Api.ApiMethod}}Ser(autoModelRes); err != nil {
    	global.HS_LOG.Error(err.Error(), zap.Any("err", err))
    	utils.FailMag("创建失败", c)
    	return
    }
	utils.SuccessData("3232", c)
}

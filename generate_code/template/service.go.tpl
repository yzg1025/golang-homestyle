package {{ .ApiPackageName}}

import (
	"errors"
    "gin/generate_code/auto_model"
    "gin/global"
    "gorm.io/gorm"
)

func Auto{{ .Api.ApiMethod}}Ser(autoModel auto_model.{{.Api.ApiName}}) (err error) {
	if !errors.Is(global.HS_DB.Where("url = ? ", autoModel.Url).First(&auto_model.{{.Api.ApiName}}{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同的banner")
	}
	return global.HS_DB.Create(&autoModel).Error
}
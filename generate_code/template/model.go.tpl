package auto_model

type {{.Api.ApiName}}Req struct {
	{{- range .Api.SimpleApiRequestField}}
    {{.FieldName}}  {{.FieldType}} `gorm:"column:{{.FieldName}}" json:"{{.FieldName}}"`
    {{- end}}
}

type {{.Api.ApiName}} struct {
	{{- range .Api.SimpleApiResponseField}}
    {{.ColumnName}} {{.FieldType}} `gorm:"column:{{ .ColumnNameLower }}{{- if .DefaultValue -}};default:{{.DefaultValue}}{{- end }}{{- if .ColumnDesc -}};comment:{{ .ColumnDesc}}{{- end }}" json:"{{ .ColumnNameLower}}"`
    {{- end}}
}

func ({{.Api.ApiName}}) TableName() string {
  return "auto_{{.Api.ApiName}}"
}
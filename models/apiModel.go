package models

import (
	"time"
)

type SimpleAPI struct {
	ApiId                  int64                    `gorm:"primarykey" json:"apiId"`
	AccessAuth             bool                     `gorm:"column:accessAuth;default:false" json:"accessAuth"`
	ApiDesc                string                   `gorm:"column:apiDesc" json:"apiDesc"`
	ReqMethod              string                   `gorm:"column:reqMethod" json:"reqMethod"`
	ApiName                string                   `gorm:"column:apiName" json:"apiName"`
	ApiMethod              string                   `gorm:"column:apiMethod" json:"apiMethod"`
	ApiPath                string                   `gorm:"column:apiPath" json:"apiPath"`
	CreateTime             time.Time                `json:"createTime"`
	SimpleApiRequestField  []SimpleApiRequestField  `gorm:"foreignKey:FieldId;references:ApiId" json:"simpleApiRequestField"`
	SimpleApiResponseField []SimpleApiResponseField `gorm:"foreignKey:ColumnId;references:ApiId" json:"simpleApiResponseField"`
}

type SimpleApiRequestField struct {
	FieldId      int64  `gorm:"column:requestFieldId" json:"requestFieldId"`
	FieldType    string `gorm:"column:fieldType" json:"fieldType"`
	InType       string `json:"inType"`
	FieldName    string `gorm:"column:fieldName" json:"fieldName"`
	FieldDesc    string `gorm:"column:FieldDesc" json:"fieldDesc"`
	FieldRequire bool   `gorm:"column:fieldRequire;default:false" json:"fieldRequire"`
}

type SimpleApiResponseField struct {
	ColumnId        int64  `gorm:"column:column_id" json:"column_id"`
	ColumnName      string `gorm:"column:columnName" json:"columnName"`
	ColumnNameLower string `gorm:"column:columnNameLower" json:"columnNameLower"`
	FieldType       string `gorm:"column:fieldType" json:"fieldType"`
	DefaultValue    string `gorm:"column:defaultValue" json:"defaultValue"`
	ColumnDesc      string `gorm:"column: columnDesc" json:"columnDesc"`
}

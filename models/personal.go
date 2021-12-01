package models

type PersonalCateGory struct {
	ID           string `gorm:"primary_key;auto_increment" json:"id"`
	CategoryName string `json:"category_name"`
}

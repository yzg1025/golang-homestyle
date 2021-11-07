package models

type AreaCode struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CountryCode string `json:"country_code"`
	Cname string `json:"cname"`
	Code string `json:"code"`
}

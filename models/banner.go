package models

type Banner struct {
	ID          int    `gorm:"primarykey" json:"id"`
	Url         string `json:"url"`
	RedirectUrl string `json:"redirect_url"`
	Order       int    `json:"order"`
	Title       string `json:"title"`
}

package models

type Login struct {
	ID int `gorm:"primarykey" json:"id"`
	NickName string `json:"user_name"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	LoginMethod string `json:"login_method"`
	Code string `json:"code"`
	RegisterTime string `json:"register_time"`
	Header string `json:"header"`
}

type ChangePassword struct {
	Phone    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}
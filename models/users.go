package models

import "time"

type Login struct {
	Uid int `gorm:"primarykey" json:"uid"`
	NickName string `json:"user_name"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	LoginMethod string `json:"login_method"`
	RegisterTime time.Time `json:"register_time"`
	Header string `json:"header" gorm:"default:http://qqbi9utzj.hd-bkt.clouddn.com/Fhh6BiidOhTl-a86HdU2fozVI8BI;comment:用户头像"`
}

type Register struct {
	Phone string `json:"phone"`
	NickName string `json:"nickname"`
	Password string `json:"password"`
}

type ChangePassword struct {
	Phone    string `json:"phone"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}
package models

import "gorm.io/gorm"

type Logs struct {
	gorm.Model
	Content string `gorm:"type:longtext" json:"content"`
}
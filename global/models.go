package global

import (
	"gorm.io/gorm"
	"time"
)

type Models struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type GetById struct {
	Id int `json:"id" form:"id"`
}

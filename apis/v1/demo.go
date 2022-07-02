package v1

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Demo interface {
	GetName(name string) string
	GetAge(age string) string
	MyDesc() string
}

var _ Demo = (*InterDb)(nil)

type InterDb struct {
	db   *gorm.DB
	Age  string
	Name string
}

func NewInterDb(name, age string) *InterDb {
	return &InterDb{
		Name: name,
		Age:  age,
	}
}

func (d *InterDb) GetName(name string) string {
	nameOne := fmt.Sprintf("我是%s,我来自河南安阳", name)
	return nameOne
}

func (d *InterDb) GetAge(age string) string {
	ageOne := fmt.Sprintf("我今年%s岁了,难得有这个时间,我现在正在玩golang的interface", age)
	return ageOne
}

func (d *InterDb) MyDesc() string {
	desc := d.GetName(d.Name) + d.GetAge(d.Age)
	return desc
}

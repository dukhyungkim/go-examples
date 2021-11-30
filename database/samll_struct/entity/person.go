package entity

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name    string `gorm:"size:24"`
	Age     int
	Address string
	Phone   string
}

func (Person) TableName() string {
	return "person"
}

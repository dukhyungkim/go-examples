package entity

import (
	"gorm.io/datatypes"
	"time"
)

type Person struct {
	ID          uint
	Age         uint8
	Address     string `gorm:"size:64"`
	Phone       string `gorm:"size:16"`
	Information datatypes.JSON
	CreatedAt   time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:milli"`
}

func (Person) TableName() string {
	return "persons"
}

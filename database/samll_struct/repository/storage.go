package repository

import (
	"go-examples/database/samll_struct/entity"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func (s *Storage) Migration() error {
	return s.db.AutoMigrate(&entity.Person{})
}

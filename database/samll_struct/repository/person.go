package repository

import (
	"go-examples/database/samll_struct/entity"
	"gorm.io/gorm"
)

type PersonRepository interface {
	GetPerson(id int64) (*entity.Person, error)
	SavePerson(person *entity.Person) error
	UpdatePerson(person *entity.Person) (*entity.Person, error)
	DeletePerson(id int64) error
}

type personRepository struct {
	db *gorm.DB
}

func NewPersonRepository(storage *Storage) PersonRepository {
	return &personRepository{db: storage.db}
}

func (p *personRepository) GetPerson(id int64) (*entity.Person, error) {
	panic("implement me")
}

func (p *personRepository) SavePerson(person *entity.Person) error {
	return p.db.Create(person).Error
}

func (p *personRepository) UpdatePerson(person *entity.Person) (*entity.Person, error) {
	panic("implement me")
}

func (p *personRepository) DeletePerson(id int64) error {
	panic("implement me")
}

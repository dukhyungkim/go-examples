package repository

import "go-examples/database/samll_struct/entity"

type PersonRepository interface {
	GetPerson(id int64) (*entity.Person, error)
	CreatePerson(person *entity.Person) (int64, error)
	UpdatePerson(person *entity.Person) (*entity.Person, error)
	DeletePerson(id int64) error
}

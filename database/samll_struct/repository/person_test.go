package repository

import (
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go-examples/database/samll_struct/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"time"
)

type AnyTime struct {
}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestNewPersonRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	storage := &Storage{db: gdb}

	type args struct {
		storage *Storage
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"create new person repository",
			args{storage: storage},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPersonRepository(tt.args.storage)
			assert.NotNil(t, got)
		})
	}

}

func Test_personRepository_CreatePerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	assert.NoError(t, err)

	p := NewPersonRepository(&Storage{db: gdb})

	person := entity.Person{
		Model:   gorm.Model{},
		Name:    "user",
		Age:     20,
		Address: "address",
		Phone:   "010-1234-5678",
	}

	type args struct {
		person *entity.Person
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr error
	}{
		{
			"add new user",
			args{person: &person},
			1,
			nil,
		},
	}

	const sqlInsert = `
					INSERT INTO "person" ("name","age","address","phone") 
						VALUES ($1,$2,$3,$4)`
	const newId = 1
	mock.ExpectBegin() // start transaction
	//mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).
	mock.ExpectExec("INSERT INTO `person`").
		WithArgs(AnyTime{}, AnyTime{}, nil, person.Name, person.Age, person.Address, person.Phone).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit() // commit transaction

	// we make sure that all expectations were met
	//if err := mock.ExpectationsWereMet(); err != nil {
	//	t.Errorf("there were unfulfilled expectations: %s", err)
	//}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := p.SavePerson(tt.args.person)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
			assert.NoError(t, err)
		})
	}
}

func Test_personRepository_DeletePerson(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &personRepository{
				db: tt.fields.db,
			}
			if err := p.DeletePerson(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeletePerson() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_personRepository_GetPerson(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Person
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &personRepository{
				db: tt.fields.db,
			}
			got, err := p.GetPerson(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPerson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPerson() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_personRepository_UpdatePerson(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		person *entity.Person
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Person
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &personRepository{
				db: tt.fields.db,
			}
			got, err := p.UpdatePerson(tt.args.person)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdatePerson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdatePerson() got = %v, want %v", got, tt.want)
			}
		})
	}
}

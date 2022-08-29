package entity

import (
	"dvdrental/graph/model"
	"dvdrental/repository"
	"fmt"
)

type Customer struct {
	model.Customer
}

func NewCustomer(c *repository.Customer) *Customer {
	var active *int
	if c.Active.Valid {
		intActive := int(c.Active.Int32)
		active = &intActive
	}

	return &Customer{
		model.Customer{
			ID:         fmt.Sprint(c.CustomerID),
			StoreID:    fmt.Sprint(c.StoreID),
			FirstName:  c.FirstName,
			LastName:   c.LastName,
			Email:      c.Email.String,
			AddressID:  fmt.Sprint(c.AddressID),
			CreateDate: c.CreateDate,
			LastUpdate: &c.LastUpdate.Time,
			Active:     active,
		},
	}
}

func (c *Customer) ToModel() *model.Customer {
	return &c.Customer
}

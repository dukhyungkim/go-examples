package entity

import (
	"fmt"
	"graphql/graph/model"
	"strconv"
)

type Link struct {
	ID      int64
	Title   string
	Address string
	//User    *User `gorm:"foreignKey:ID"`
}

func NewLink(link *model.Link) (*Link, error) {
	//user, err := NewUser(link.User)
	//if err != nil {
	//	return nil, err
	//}

	var id int64
	var err error
	if link.ID != "" {
		id, err = strconv.ParseInt(link.ID, 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return &Link{
		ID:      id,
		Title:   link.Title,
		Address: link.Address,
		//User:    user,
	}, nil
}

func (Link) TableName() string {
	return "links"
}

func (l *Link) ToModel() *model.Link {
	return &model.Link{
		ID:      fmt.Sprint(l.ID),
		Title:   l.Title,
		Address: l.Address,
		User:    nil,
	}
}

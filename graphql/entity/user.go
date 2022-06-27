package entity

import (
	"graphql/graph/model"
	"strconv"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func NewUser(user *model.User) (*User, error) {
	if user == nil {
		return nil, nil
	}

	id, err := strconv.ParseInt(user.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       id,
		Username: user.Name,
	}, nil
}

func (User) TableName() string {
	return "users"
}

package entity

import (
	"dvdrental/graph/model"
	"dvdrental/repository"
	"fmt"
)

type Actor struct {
	model.Actor
}

func NewActor(a *repository.Actor) *Actor {
	return &Actor{
		model.Actor{
			ID:          fmt.Sprint(a.ActorID),
			FirstName:   a.FirstName,
			LastName:    a.LastName,
			LastUpdates: a.LastUpdate,
		},
	}
}

func (a *Actor) ToModel() *model.Actor {
	return &a.Actor
}

//go:build wireinject

package main

import "github.com/google/wire"

func InitializeEvent(s string) Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}

package main

type Message string

func NewMessage(s string) Message {
	return Message(s)
}

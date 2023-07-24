package main

func main() {
	message := NewMessage("Hi there!")
	greeter := NewGreeter(message)
	event := NewEvent(greeter)

	event.Start()
}

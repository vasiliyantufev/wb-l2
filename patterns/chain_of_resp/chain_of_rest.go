package main

import "fmt"

func main() {
	handlers := &ConcreteHandlerA{
		next: &ConcreteHandlerB{
			next: &ConcreteHandlerC{},
		},
	}

	result := handlers.SendRequest(3)

	fmt.Println(result)
}

// Handler предоставляет интерфейс обработчика.
type Handler interface {
	SendRequest(message int) string
}

// ConcreteHandlerA реализует обработчик "A".
type ConcreteHandlerA struct {
	next Handler
}

// SendRequest реализация.
func (h *ConcreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler 1"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// ConcreteHandlerB реализует обработчик "B".
type ConcreteHandlerB struct {
	next Handler
}

// SendRequest реализация.
func (h *ConcreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Im handler 2"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// ConcreteHandlerC реализует обработчик "C".
type ConcreteHandlerC struct {
	next Handler
}

// SendRequest реализация.
func (h *ConcreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Im handler 3"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

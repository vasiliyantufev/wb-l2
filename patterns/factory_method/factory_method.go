package main

import (
	"fmt"
	"log"
)

// action помогает клиентам узнать доступные действия.
type action string

const (
	A action = "A"
	B action = "B"
	C action = "C"
)

func main() {
	factory := NewCreator()

	products := []string{
		factory.CreateProduct(A).Use(),
		factory.CreateProduct(B).Use(),
		factory.CreateProduct(C).Use(),
	}

	fmt.Println(products)
}

// Creator предоставляет заводской интерфейс.
type Creator interface {
	CreateProduct(action action) Product // Фабричный метод
}

// Product обеспечивает интерфейс продукта.
// Все продукты, возвращаемые фабрикой, должны иметь единый интерфейс.
type Product interface {
	Use() string //Каждый продукт должен быть полезным
}

// ConcreteCreator реализует Creator интерфейс.
type ConcreteCreator struct{}

// NewCreator это ConcreteCreator конструктор.
func NewCreator() Creator {
	return &ConcreteCreator{}
}

// CreateProduct фабричный метод.
func (p *ConcreteCreator) CreateProduct(action action) Product {
	var product Product

	switch action {
	case A:
		product = &ConcreteProductA{string(action)}
	case B:
		product = &ConcreteProductB{string(action)}
	case C:
		product = &ConcreteProductC{string(action)}
	default:
		log.Fatalln("Unknown Action")
	}

	return product
}

// ConcreteProductA реализует продукт "A".
type ConcreteProductA struct {
	action string
}

// Use возвращает действие продукта.
func (p *ConcreteProductA) Use() string {
	return p.action
}

// ConcreteProductB реализует продукт "B".
type ConcreteProductB struct {
	action string
}

// Use возвращает действие продукта.
func (p *ConcreteProductB) Use() string {
	return p.action
}

// ConcreteProductC реализует продукт "C".
type ConcreteProductC struct {
	action string
}

// Use возвращает действие продукта.
func (p *ConcreteProductC) Use() string {
	return p.action
}

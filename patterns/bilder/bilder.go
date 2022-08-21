package main

import "fmt"

func main() {

	product := new(Product)

	director := Director{&ConcreteBuilder{product}}
	director.Construct()

	result := product.Show()

	fmt.Println(result)
}

// Builder предоставляет интерфейс.
type Builder interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
}

// Director реализует менеджер
type Director struct {
	builder Builder
}

// Construct говорит что делать и в каком порядке.
func (d *Director) Construct() {
	d.builder.MakeHeader("Header")
	d.builder.MakeBody("Body")
	d.builder.MakeFooter("Footer")
}

// ConcreteBuilder реализует интерфейс Builder.
type ConcreteBuilder struct {
	product *Product
}

// MakeHeader реализует заголовок документа
func (b *ConcreteBuilder) MakeHeader(str string) {
	b.product.Content += "<header>" + str + "</header>"
}

// MakeBody реализует тело документа
func (b *ConcreteBuilder) MakeBody(str string) {
	b.product.Content += "<article>" + str + "</article>"
}

// MakeFooter реализует подвал документа
func (b *ConcreteBuilder) MakeFooter(str string) {
	b.product.Content += "<footer>" + str + "</footer>"
}

// Product
type Product struct {
	Content string
}

// Show возвращает product.
func (p *Product) Show() string {
	return p.Content
}

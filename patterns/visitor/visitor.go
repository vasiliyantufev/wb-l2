package main

import "fmt"

func main() {
	city := new(City)
	city.Add(&SushiBar{})
	city.Add(&Pizzeria{})
	city.Add(&BurgerBar{})
	result := city.Accept(&People{})
	fmt.Println(result)
}

// Visitor обеспечивает интерфейс посетителя.
type Visitor interface {
	VisitSushiBar(p *SushiBar) string
	VisitPizzeria(p *Pizzeria) string
	VisitBurgerBar(p *BurgerBar) string
}
// Place предоставляет интерфейс для места, которое посетитель должен посетить.
type Place interface {
	Accept(v Visitor) string
}
// People реализует интерфейс посетителя.
type People struct {
}
// VisitSushiBar осуществляет посещение SushiBar.
func (v *People) VisitSushiBar(p *SushiBar) string {
	return p.BuySushi()
}
// VisitPizzeria осуществляет посещение Pizzeria.
func (v *People) VisitPizzeria(p *Pizzeria) string {
	return p.BuyPizza()
}
// VisitBurgerBar осуществляет посещение BurgerBar.
func (v *People) VisitBurgerBar(p *BurgerBar) string {
	return p.BuyBurger()
}
// City реализует коллекцию мест для посещения.
type City struct {
	places []Place
}
// Add добавляет место в коллекцию.
func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}
// Accept осуществляет посещение всех мест в городе.
func (c *City) Accept(v Visitor) string {
	var result string
	for _, p := range c.places {
		result += p.Accept(v)
	}
	return result
}
type SushiBar struct {
}
func (s *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(s)
}
func (s *SushiBar) BuySushi() string {
	return "Buy sushi..."
}
type Pizzeria struct {
}
func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(p)
}
func (p *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

type BurgerBar struct {
}
func (b *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(b)
}
func (b *BurgerBar) BuyBurger() string {
	return "Buy burger..."
}

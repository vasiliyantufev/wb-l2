// Package facade is an example of the Facade Pattern.
package main

import (
	"fmt"
	"strings"
)

/*
Фасад
*/

type Man struct {
	house *House
	tree  *Tree
	child *Child
}

// Конструктор
func NewMan() *Man {
	return &Man{
		house: &House{},
		tree:  &Tree{},
		child: &Child{},
	}
}

// Возвращает man
func (m *Man) todo(h, t, c string) string {
	result := []string{
		m.house.Build(h),
		m.tree.Grow(t),
		m.child.Born(c),
	}
	return strings.Join(result, "\n")
}

// Реализует подсистему house
type House struct {
}

func (h *House) Build(s string) string {
	return s
}

// Реализуем подсистему tree
type Tree struct {
}

func (t *Tree) Grow(s string) string {
	return s
}

// Реализуем подсистему children
type Child struct {
}

func (c *Child) Born(s string) string {
	return s
}

func main() {
	man := NewMan().todo("построил", "посадил", "воспитал")
	fmt.Println(man)
}

package main

import "fmt"

func main() {

	invoker := &Invoker{}
	receiver := &Receiver{}

	invoker.StoreCommand(&ToggleOnCommand{receiver: receiver})
	invoker.StoreCommand(&ToggleOffCommand{receiver: receiver})

	result := invoker.Execute()

	fmt.Println(result)
}

// Command обеспечивает командный интерфейс.
type Command interface {
	Execute() string
}

// ToggleOnCommand реализует командный интерфейс
type ToggleOnCommand struct {
	receiver *Receiver
}

// Выполнить команду.
func (c *ToggleOnCommand) Execute() string {
	return c.receiver.ToggleOn()
}

// ToggleOffCommand реализует командный интерфейс.
type ToggleOffCommand struct {
	receiver *Receiver
}

// Выполнить команду.
func (c *ToggleOffCommand) Execute() string {
	return c.receiver.ToggleOff()
}

// Receiver реализация.
type Receiver struct {
}

// ToggleOn реализация.
func (r *Receiver) ToggleOn() string {
	return "Toggle On"
}

// ToggleOff реализация.
func (r *Receiver) ToggleOff() string {
	return "Toggle Off"
}

// Invoker реализация.
type Invoker struct {
	commands []Command
}

// StoreCommand добавляет команду.
func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

// UnStoreCommand удаляет команду.
func (i *Invoker) UnStoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

// Выполнить все команды.
func (i *Invoker) Execute() string {
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n"
	}
	return result
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*


Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit)
*/

func main() {

	//h := exec.Command("pwd")
	//h.Run()

	//var c = exec.Command(command)

	//print(c)

	scan := bufio.NewScanner(os.Stdin)
	//fmt.Print("Command: ")
	for {
		fmt.Print("Enter command: ")
		if scan.Scan() {
			command := scan.Text()
			commands := strings.Split(command, " | ")
			//fmt.Print(commands)
			exec(commands)
		}
	}
}

func exec(commands []string) {

	for _, command := range commands {
		cmd := strings.Split(command, " ")
		//print(cmd[0])
		switch cmd[0] {
		case "cd":
			//fmt.Print(cmd[1])
			os.Chdir(cmd[1])
		case `\exit`:
			os.Exit(1)
		case "echo":
			for i := 1; i < len(cmd); i++ {
				fmt.Fprintf(os.Stdout, cmd[i]+" ")
			}
			fmt.Println()
			//default:

			//c.Stderr = os.Stderr
			//c.Stdout = os.Stdout

		}
	}
	//for _, command := range commands {
	//}

}

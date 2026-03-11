package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n') // читаем все до нажатия enter
		parts := strings.Fields(input)
		if len(parts) == 0 { //если ничего не ввели на экран, переходим к следующей итерации цикла for
			continue
		}

		command := parts[0]
		args := parts[1:]
		switch {
		case command == "dir":
			if len(args) == 0 {
				directory, err := os.ReadDir(".") // если в args нет пути, значит читаем текущую директорию
				if err != nil {
					fmt.Println("Cannot read directory")
					continue
				}
				for _, entry := range directory {
					if entry.IsDir() { // true = directory, false = file
						fmt.Println("[Dir] ", entry.Name())
					} else {
						fmt.Println(entry.Name())
					}
				}
			} else {
				entries, err := os.ReadDir(args[0])
				if err != nil {
					fmt.Println("Path error")
					continue
				}
				for _, entry := range entries {
					if entry.IsDir() {
						fmt.Println("[Dir] ", entry.Name())
					} else {
						fmt.Println(entry.Name())
					}

				}
			}
		case command == "cd":
			if len(args) == 0 {
				home, err := os.UserHomeDir()
				if err != nil {
					fmt.Println("Cannot find home directory:", err)
					continue
				}
				err = os.Chdir(home) // Переход в домашнюю директорию
				if err != nil {
					fmt.Println("Cannot change to home directory:", err)
					continue
				}
			} else {
				err := os.Chdir(args[0])
				if err != nil {
					fmt.Println("Path error")
				}
			}
		case command == "mkdir":
			err := os.Mkdir(args[0], 0755)
			if err != nil {
				fmt.Println("Cannot create directory:", err)
			}
		case command == "rmdir":
			err := os.Remove(args[0])
			if err != nil {
				fmt.Println("Cannot remove:", err)
			}
		case command == "exit":
			return
		default:
			fmt.Println("Incorrect command")
		}
	}
}

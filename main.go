package main

import (
	"bufio"
	"fmt"
	"os"
	"project/methodwithjson"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("")
		fmt.Println("Добро пожаловать в Task-менеджер")

		fmt.Println("Выберите действие: ")
		fmt.Println("1 - добавить задачу")
		fmt.Println("2 - удалить задачу")
		fmt.Println("0 - EXIT")

		fmt.Print("Ваш выбор: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			add := ""
			fmt.Println("")
			fmt.Print("Введите название задачи: ")
			fmt.Scanln(&add)
			methodwithjson.Add(add)
		case "2":
			del := ""
			fmt.Println("")
			fmt.Print("Введите название задачи которую хотите удалить: ")
			fmt.Scanln(&del)
			methodwithjson.Delete(del)
		case "0":
			fmt.Println("")
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

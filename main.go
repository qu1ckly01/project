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
		fmt.Println("3 - просмотреть задачи")
		fmt.Println("4 - посмотреть выполненые задачи")
		fmt.Println("5 - посмотреть невыполненые задачи")
		fmt.Println("6 - обновить статус задачи задачи")
		fmt.Println("0 - EXIT")

		fmt.Print("Ваш выбор: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("")
			fmt.Print("Введите название задачи: ")
			add, _ := reader.ReadString('\n')
			add = strings.TrimSpace(add)
			methodwithjson.Add(add)
		case "2":
			fmt.Println("")
			fmt.Print("Введите название задачи которую хотите удалить: ")
			del, _ := reader.ReadString('\n')
			del = strings.TrimSpace(del)
			methodwithjson.Delete(del)
		case "3":
			fmt.Println("")
			fmt.Println("Задачи в списке: ")
			methodwithjson.Out()
		case "4":
			fmt.Println("")
			fmt.Println("Выполненые задачи: ")
			methodwithjson.TrueTaskOut()
		case "5":
			fmt.Println("")
			fmt.Println("Невыполненые задачи: ")
			methodwithjson.FalseTaskOut()
		case "6":
			fmt.Println("")
			fmt.Print("Введите название задачи: ")
			updnametask, _ := reader.ReadString('\n')
			updnametask = strings.TrimSpace(updnametask)
			methodwithjson.Update(updnametask)
		case "0":
			fmt.Println("")
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

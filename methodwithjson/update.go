package methodwithjson

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"project/workjson"
)

func Update(name string) {
	file := "task.json"

	data, err := os.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("ошибка чтения файла: %v", err)
	}

	var taskes []workjson.Task
	if len(data) > 0 {
		if err := json.Unmarshal(data, &taskes); err != nil {
			log.Fatalf("ошибка чтения JSON: %v", err)
		}
	}

	upd := false
	for i := range taskes {
		if taskes[i].Name == name {
			if taskes[i].Status {
				fmt.Println("Задача уже выполнена")
				return
			}
			taskes[i].Status = true
			upd = true
			break
		}
	}

	if !upd {
		fmt.Println("Задача не найдена")
		return
	}

	updated, err := json.MarshalIndent(taskes, "", "  ")
	if err != nil {
		log.Fatalf("ошибка конвератции JSON: %v", err)
	}
	if err := os.WriteFile(file, updated, 0644); err != nil {
		log.Fatalf("ошибка записи JSON: %v", err)
	}

	fmt.Println("Статус задачи обновлен")
}

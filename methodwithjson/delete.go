package methodwithjson

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"project/workjson"
)

func Delete(name string) {
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

	deltask := false
	for i, u := range taskes {
		if u.Name == name {
			taskes = append(taskes[:i], taskes[i+1:]...)
			deltask = true
			break
		}
	}

	if !deltask {
		fmt.Println("Такой задачи нет в списке или она уже удалена")
		return
	}

	updated, _ := json.MarshalIndent(taskes, "", "  ")
	os.WriteFile(file, updated, 0644)

	fmt.Printf("Задача %s удалена \n", name)
}

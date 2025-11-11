package methodwithjson

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"project/workjson"
	"time"
)

func Add(name string) {

	second := time.Now().Unix()
	rand.Seed(second)
	randId := rand.Intn(100) + 1

	newTask := workjson.Task{Name: name, Id: randId, Status: false}

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

	for i := range taskes {
		if taskes[i].Name == name {
			fmt.Println("Такая задача уже существует")
			return
		}
	}

	taskes = append(taskes, newTask)

	updated, err := json.MarshalIndent(taskes, "", "  ")
	if err != nil {
		log.Fatalf("ошибка конвертации в JSON: %v", err)
	}
	if err := os.WriteFile(file, updated, 0644); err != nil {
		log.Fatalf("ошибка записи JSON: %v", err)
	}

	fmt.Println("Добавлена новая задача")
}

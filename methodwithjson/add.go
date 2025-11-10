package methodwithjson

import (
	"encoding/json"
	"fmt"
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
		panic(err)
	}

	var taskes []workjson.Task
	if len(data) > 0 {
		if err := json.Unmarshal(data, &taskes); err != nil {
			panic(fmt.Errorf("ошибка чтения JSON: %v", err))
		}
	}

	taskes = append(taskes, newTask)

	updated, err := json.MarshalIndent(taskes, "", "  ")
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(file, updated, 0644); err != nil {
		panic(err)
	}

	fmt.Println("Добавлена новая задача")
}

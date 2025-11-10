package workjson

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Name   string `json:"name"`
	Id     int    `json:"id"`
	Status bool   `json:"status"`
}

func InJSON(t Task) {

	taskJSON, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Ошибка конвентации в JSON", err)
		return
	}

	err = os.WriteFile("task.json", taskJSON, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(taskJSON))
}

package methodwithjson

import (
	"encoding/json"
	"fmt"
	"os"
	"project/workjson"
)

func Delete(name string) {
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

	for i, u := range taskes {
		if u.Name == name {
			taskes = append(taskes[:i], taskes[i+1:]...)
			break
		}
	}

	updated, _ := json.MarshalIndent(taskes, "", "  ")
	os.WriteFile(file, updated, 0644)

	fmt.Printf("Задача %s удалена \n", name)
}

package methodwithjson

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"project/workjson"
)

func Out() {
	file := "task.json"

	data, err := os.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	var taskes []workjson.Task
	if len(data) > 0 {
		if err := json.Unmarshal(data, &taskes); err != nil {
			log.Fatalf("ошибка чтения JSON: %v", err)
		}
	}

	for i, t := range taskes {
		fmt.Printf("%d)\n", i+1)
		fmt.Printf("  Name: %s\n", t.Name)
		fmt.Printf("  ID: %v\n", t.Id)
		fmt.Printf("  Status: %v\n", t.Status)
		fmt.Println("---")
	}
}

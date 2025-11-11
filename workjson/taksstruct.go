package workjson

type Task struct {
	Name   string `json:"name"`
	Id     int    `json:"id"`
	Status bool   `json:"status"`
}

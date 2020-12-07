package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Rank int    `json:"rank"`
	Year int    `json:"year"`
}

func main() {
	jsonData := `{"name":"Fahim", "rank":3, "year":2018}`

	var student Student
	err := json.Unmarshal([]byte(jsonData), &student)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", student)
}

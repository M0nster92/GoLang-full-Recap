package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Info Info   `json:"info"`
}

type Info struct {
	Rank  int    `json:"rank"`
	Year  int    `json:"year"`
	Grade string `json:"grade"`
}

func main() {
	info := Info{Rank: 3, Year: 2020, Grade: "B"}
	student := Student{Name: "Fahim Al Sami", Info: info}

	m, err := json.Marshal(student)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(m))
}

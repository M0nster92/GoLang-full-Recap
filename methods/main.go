package main

import "fmt"

type Student struct {
	Name string
	Id   int
}

func (s *Student) ChangeID(newId int) {
	s.Id = newId
}

func main() {
	s := Student{Name: "Fahim Al Sami", Id: 15}
	fmt.Println(s)
	s.ChangeID(5)
	fmt.Println(s)
}

package main

import (
	"fmt"
)

type MyClass struct {
	Name string
}

type MyAge struct {
	Age int32
}

type My struct {
	myclass MyClass
	myage   MyAge
}

func test(class interface{}) {
	classStr := class.(My)
	fmt.Println(classStr.myclass.Name)
}

func main() {
	var my My
	my.myclass.Name = "John"
	my.myage.Age = 23
	test(my)
}

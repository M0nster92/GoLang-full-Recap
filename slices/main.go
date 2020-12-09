package main

import "fmt"

func main() {
	//slices are similar to array but more dynamic

	nameList := [5]string{"Jack", "Robert", "Ryan", "Neo", "Kai"}

	sliceList := nameList[1:5]

	fmt.Println(sliceList)
}

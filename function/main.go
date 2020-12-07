package main

import "fmt"

func simpleFunc(fname string, lname string) string {
	fullName := fname + " " + lname
	return fullName
}

func main() {
	fullName := simpleFunc("Fahim", "Sami")
	fmt.Println(fullName)
}

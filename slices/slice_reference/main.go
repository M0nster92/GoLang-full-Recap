package main

import "fmt"

func main() {
	names := [4]string{
		"Theo",
		"Leo",
		"Mia",
		"Steven",
	}

	fmt.Println(names)

	a := names[0:2]

	fmt.Println(a)

	a[0] = "Frank"

	fmt.Println(a)
}

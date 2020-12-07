package main

import "fmt"

func linearSearch(data []int, key int) (value int, b bool) {
	for _, item := range data {
		if item == key {
			value = item
			return value, true
		}
	}

	return 0, false
}

func main() {
	data := []int{34, 56, 32, 65, 342, 786}
	val, b := linearSearch(data, 342)
	if b {
		fmt.Println("Value found ", val)
	} else {
		fmt.Println("Value Not Found")
	}
}

package main

import "fmt"

func factorial(val int) uint64 {
	var result uint64 = 1
	if val < 0 {
		return result
	} else {
		for i := 1; i <= val; i++ {
			result *= uint64(i)
		}
	}

	return result
}

func main() {
	fact := factorial(5)

	fmt.Println(fact)
}

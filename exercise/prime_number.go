package main

/*func main() {
	test := prime(29)
	if test {
		fmt.Println("Current Number is a prime Number")
	} else {
		fmt.Println("Current Number is not a prime Number")
	}
} */

func prime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

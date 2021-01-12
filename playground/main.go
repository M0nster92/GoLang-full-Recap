package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	add := now.Add(time.Minute * 10)
	fmt.Println(add)
	sub := now.Add(time.Minute * (-10)).Unix()
	fmt.Println(sub)
}

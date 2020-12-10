package main

import (
	"fmt"
	"os/exec"
)

func command() {
	out, err := exec.Command("pwd").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println(string(out))
}

func main() {
	command()
}

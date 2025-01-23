package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Enter a string:")
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	containsA := false
	isValid := len(input) >= 3 && input[0] == 'i' && input[len(input)-1] == 'n'
	for i := 0; i < len(input); i++ {
		if input[i] == 'a' {
			containsA = true
			break
		}
	}
	if isValid && containsA {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

package main

import (
	"fmt"
)

func main() {
	var s string
	_, err := fmt.Scanf("%s\n", &s)
	if err != nil {
		panic(err)
	}

	chars := []byte(s)
	for i, char := range chars {
		if char == '5' {
			chars[i] = '*'
		}
	}
	s = string(chars)
	fmt.Println(s)
}

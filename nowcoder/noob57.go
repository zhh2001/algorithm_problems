package main

import (
	"fmt"
)

func Caesar(s string, n int) string {
	chars := []byte(s)
	for i, char := range chars {
		chars[i] = char + byte(n)
		for chars[i] > 'z' {
			chars[i] -= 26
		}
	}
	return string(chars)
}

func main() {
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}

	var s string
	_, err = fmt.Scanf("%s\n", &s)
	if err != nil {
		panic(err)
	}

	fmt.Println(Caesar(s, n))
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}

	digits := strconv.Itoa(n)
	chars := make([]byte, 0, n+n/3)
	for i, char := range digits {
		chars = append(chars, byte(char))
		if (len(digits)-i-1)%3 == 0 && i != len(digits)-1 {
			chars = append(chars, ',')
		}
	}
	str := string(chars)
	fmt.Println(str)
}

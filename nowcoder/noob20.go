package main

import (
	"fmt"
)

func Sum(n int) int {
	return n * (n + 1) / 2
}

func main() {
	n := 0
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}
	fmt.Println(Sum(n))
}

package main

import (
	"fmt"
)

func Collatz(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return n*3 + 1
	}
}

func main() {
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}
	fmt.Println(Collatz(n))
}

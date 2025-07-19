package main

import (
	"fmt"
)

func Fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	a, b := 1, 1
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	var n int
	var err error
	_, err = fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}
	fmt.Println(Fibonacci(n))
}

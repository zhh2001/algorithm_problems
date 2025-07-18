package main

import (
	"fmt"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
)

func main() {
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}
	n *= MB
	n /= 4
	fmt.Println(n)
}

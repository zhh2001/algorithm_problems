package main

import (
	"fmt"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}
	sum := 0
	sum += n % 10
	n /= 10
	sum += n % 10
	n /= 10
	sum += n % 10
	n /= 10
	sum += n % 10
	fmt.Println(sum)
}

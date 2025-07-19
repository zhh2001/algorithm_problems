package main

import (
	"fmt"
)

func main() {
	var n int
	var err error
	_, err = fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}

	var sum, total int
	for i := 1; i <= n; i++ {
		sum += i
		total += sum
	}
	fmt.Println(total)
}

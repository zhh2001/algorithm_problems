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

	var sum float64
	for i := 1; i <= n; i++ {
		sum += 1 / float64(i)
	}
	fmt.Println(sum)
}

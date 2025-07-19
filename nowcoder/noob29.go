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

	var a, b int
	for i := 0; i < n; i++ {
		_, err = fmt.Scanf("%d %d\n", &a, &b)
		if err != nil {
			panic(err)
		}
		fmt.Println(a + b)
	}
}

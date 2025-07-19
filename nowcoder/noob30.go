package main

import (
	"fmt"
)

func main() {
	var a, b int
	var err error
	for {
		_, err = fmt.Scanf("%d %d\n", &a, &b)
		if err != nil {
			panic(err)
		}
		if a == 0 && b == 0 {
			break
		}
		fmt.Println(a + b)
	}
}

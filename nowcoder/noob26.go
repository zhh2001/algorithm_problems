package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	_, err := fmt.Scanf("%d %d %d\n", &a, &b, &c)
	if err != nil {
		panic(err)
	}
	avg := (a + b + c) / 3
	if avg < 60 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

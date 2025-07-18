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
	if n%2 == 0 && n > 50 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

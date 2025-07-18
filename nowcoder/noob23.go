package main

import (
	"fmt"
)

func main() {
	var a, b int
	_, err := fmt.Scanf("%d %d\n", &a, &b)
	if err != nil {
		panic(err)
	}
	if a > b {
		fmt.Println(">")
	} else if a < b {
		fmt.Println("<")
	} else {
		fmt.Println("=")
	}
}

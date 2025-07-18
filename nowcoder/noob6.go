package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		} else {
			fmt.Println(a + b)
		}
	}
}

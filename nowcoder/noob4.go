package main

import (
	"fmt"
)

func main() {
	var str string
	for {
		n, _ := fmt.Scanf("%s\n", &str)
		if n == 0 {
			break
		} else {
			fmt.Println(str)
		}
	}
}

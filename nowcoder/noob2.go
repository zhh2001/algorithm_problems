package main

import (
	"fmt"
)

func main() {
	var i int32
	for {
		n, _ := fmt.Scanf("%d", &i)
		if n == 0 {
			break
		} else {
			fmt.Println(i)
		}
	}
}

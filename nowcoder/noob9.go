package main

import (
	"fmt"
)

func main() {
	var a int
	for {
		n, _ := fmt.Scanf("%d\n", &a)
		if n == 0 {
			break
		} else {
			if a <= 0 {
				a = -a
			}
			fmt.Printf("%d\n", a%10)
		}
	}
}

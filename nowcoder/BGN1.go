package main

import (
	"fmt"
)

func main() {
	var i, n int
	for ; n != 1; i++ {
		_, err := fmt.Scan(&n)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(i)
}

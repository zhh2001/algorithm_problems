package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 0, 100)
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	for n != 0 {
		slice = append(slice, n)
		_, err = fmt.Scan(&n)
		if err != nil {
			panic(err)
		}
	}
	for i := len(slice) - 1; i >= 0; i-- {
		fmt.Printf("%d ", slice[i])
	}
}

package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 0, 100)
	res := make([]int, 0, 100)
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	var m int
	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&m)
		if err != nil {
			panic(err)
		}
		count := 0
		for _, v := range slice {
			if v < m {
				count++
			}
		}
		fmt.Printf("%d ", count)
		res = append(res, count)
		slice = append(slice, m)
	}
}

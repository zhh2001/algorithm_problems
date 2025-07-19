package main

import (
	"fmt"
)

func main() {
	var n, m int
	_, err := fmt.Scan(&n, &m)
	if err != nil {
		panic(err)
	}

	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var t int
			_, err = fmt.Scan(&t)
			if err != nil {
				panic(err)
			}
			sum += t
		}
	}

	fmt.Println(sum)
}

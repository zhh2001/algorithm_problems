package main

import (
	"fmt"
)

func Josephus(n, k, m int) int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i
	}
	for len(slice) > 1 {
		k += m - 1
		k %= len(slice)
		copy(slice[k:], slice[k+1:])
		slice = slice[:len(slice)-1]
	}
	return slice[0]
}

func main() {
	var n, k, m int
	_, err := fmt.Scan(&n, &k, &m)
	if err != nil {
		panic(err)
	}

	fmt.Println(Josephus(n, k, m))
}

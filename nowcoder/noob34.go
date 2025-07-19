package main

import (
	"fmt"
)

func main() {
	var n int
	var err error
	_, err = fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&slice[i])
	}

	var a, b int
	a = Max[int](slice[0], slice[1:]...)
	b = Min[int](slice[0], slice[1:]...)
	fmt.Println(a - b)
}

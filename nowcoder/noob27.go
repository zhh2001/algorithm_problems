package main

import (
	"fmt"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func Max[T Ordered](x T, y ...T) T {
	for _, n := range y {
		if x < n {
			x = n
		}
	}
	return x
}

func Min[T Ordered](x T, y ...T) T {
	for _, n := range y {
		if x > n {
			x = n
		}
	}
	return x
}

func main() {
	var a, b, c int
	_, err := fmt.Scanf("%d %d %d\n", &a, &b, &c)
	if err != nil {
		panic(err)
	}
	m := Max[int](a, b, c)
	n := Min[int](a, b, c)
	fmt.Println("The maximum number is :", m)
	fmt.Println("The minimum number is :", n)
}

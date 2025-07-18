package main

import (
	"fmt"
)

func main() {
	var K, C, F float64
	_, err := fmt.Scanf("%f\n", &K)
	if err != nil {
		panic(err)
	}
	C = K - 273.15
	F = C*1.8 + 32
	fmt.Println(F)
}

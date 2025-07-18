package main

import (
	"fmt"
	"math"
)

func main() {
	var n uint
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%d\n", uint(math.Sqrt(float64(n))))
	}
}

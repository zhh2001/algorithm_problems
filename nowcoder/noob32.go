package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var err error
	_, err = fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum += int(math.Pow(-1, float64(i-1))) * i
	}
	fmt.Println(sum)
}

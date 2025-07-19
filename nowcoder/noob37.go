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

	n = int(math.Abs(float64(n)))
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	fmt.Println(sum)
}

package main

import (
	"fmt"
	"os"
)

func Score(a, b, c int) int {
	x := float64(a) * 0.2
	y := float64(b) * 0.3
	z := float64(c) * 0.5
	return int(x + y + z)
}

func main() {
	var a, b, c int
	n, err := fmt.Scanf("%d %d %d\n", &a, &b, &c)
	if err != nil {
		panic(err)
	}
	if n != 3 {
		os.Exit(1)
	}
	fmt.Println(Score(a, b, c))
}

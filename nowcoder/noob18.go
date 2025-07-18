package main

import (
	"fmt"
	"os"
)

func CalcS(a, b, c int) int {
	return (a*b + b*c + c*a) * 2
}

func CalcV(a, b, c int) int {
	return a * b * c
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
	s := CalcS(a, b, c)
	fmt.Println(s)
	v := CalcV(a, b, c)
	fmt.Println(v)
}

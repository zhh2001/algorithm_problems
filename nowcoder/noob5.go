package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c float64
	var d byte
	var e string

	_, err := fmt.Scanf("%d\n", &a)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Scanf("%d\n", &b)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Scanf("%f\n", &c)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Scanf("%c\n", &d)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Scanf("%s\n", &e)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%.1f\n", c)
	fmt.Printf("%c\n", d)
	fmt.Printf("%s\n", e)
}

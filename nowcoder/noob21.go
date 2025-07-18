package main

import (
	"fmt"
)

func main() {
	var d int
	_, err := fmt.Scanf("%d\n", &d)
	if err != nil {
		panic(err)
	}
	d++
	if d == 8 {
		d = 1
	}
	fmt.Println(d)
}

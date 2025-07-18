package main

import (
	"fmt"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}
	if n%400 == 0 {
		fmt.Println("yes")
	} else if n%4 == 0 && n%100 != 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

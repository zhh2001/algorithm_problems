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
	month := n % 100
	switch month {
	case 3, 4, 5:
		fmt.Println("spring")
	case 6, 7, 8:
		fmt.Println("summer")
	case 9, 10, 11:
		fmt.Println("autumn")
	case 12, 1, 2:
		fmt.Println("winter")
	}
}

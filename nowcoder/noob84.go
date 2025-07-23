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

	var op string
	var x int
	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&op)
		if err != nil {
			panic(err)
		}

		switch op {
		case "push":
			_, err = fmt.Scanln(&x)
			if err != nil {
				panic(err)
			}
			stack = append(stack, x)
		case "pop":
			if len(stack) == 0 {
				fmt.Println("Empty")
			} else {
				stack = stack[:len(stack)-1]
			}
		case "query":
			if len(stack) == 0 {
				fmt.Println("Empty")
			} else {
				fmt.Printf("%d\n", stack[len(stack)-1])
			}
		case "size":
			fmt.Printf("%d\n", len(stack))
		}
	}
}

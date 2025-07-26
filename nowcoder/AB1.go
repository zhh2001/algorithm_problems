package main

import (
	"fmt"
)

func main() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		panic(err)
	}

	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var op string
		_, err = fmt.Scan(&op)
		if err != nil {
			panic(err)
		}

		switch op {
		case "push":
			var x int
			_, err = fmt.Scanln(&x)
			if err != nil {
				panic(err)
			}

			stack = append(stack, 0)
			copy(stack[1:], stack[:len(stack)-1])
			stack[0] = x
		case "pop":
			if len(stack) == 0 {
				fmt.Println("error")
			} else {
				fmt.Println(stack[0])
				stack = stack[1:]
			}
		case "top":
			if len(stack) == 0 {
				fmt.Println("error")
			} else {
				fmt.Println(stack[0])
			}
		}
	}
}

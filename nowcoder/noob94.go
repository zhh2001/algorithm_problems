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

	queen := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var op int
		_, err = fmt.Scan(&op)
		if err != nil {
			panic(err)
		}

		switch op {
		case 1:
			var x int
			_, err = fmt.Scanln(&x)
			if err != nil {
				panic(err)
			}
			queen = append(queen, x)
		case 2:
			if len(queen) == 0 {
				fmt.Println("ERR_CANNOT_POP")
			} else {
				queen = queen[1:]
			}
		case 3:
			if len(queen) == 0 {
				fmt.Println("ERR_CANNOT_QUERY")
			} else {
				fmt.Println(queen[0])
			}
		case 4:
			fmt.Println(len(queen))
		}
	}
}

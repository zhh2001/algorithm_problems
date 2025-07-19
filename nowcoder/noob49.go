package main

import (
	"fmt"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}

	flag := true
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			_, err = fmt.Scan(&arr[i][j])
			if err != nil {
				panic(err)
			}
			if i > j {
				if arr[i][j] != 0 {
					flag = false
					break
				}
			}
		}
		if !flag {
			break
		}
	}

	if flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

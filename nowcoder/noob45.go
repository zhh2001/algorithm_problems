package main

import (
	"fmt"
)

func main() {
	var n, x int
	_, err := fmt.Scan(&n, &x)
	if err != nil {
		panic(err)
	}

	var cnt int
	for i := 1; i <= n; i++ {
		tmp := i
		for tmp != 0 {
			if tmp%10 == x {
				cnt++
			}
			tmp /= 10
		}
	}
	fmt.Println(cnt)
}

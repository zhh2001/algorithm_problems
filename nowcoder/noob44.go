package main

import (
	"fmt"
)

func main() {
	var t int
	_, err := fmt.Scan(&t)
	if err != nil {
		panic(err)
	}

	var n, k int
	for i := 0; i < t; i++ {
		_, err = fmt.Scan(&n, &k)
		if err != nil {
			panic(err)
		}

		var s, cnt int
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			_, err = fmt.Scan(&arr[j])
			if err != nil {
				panic(err)
			}
			if arr[j] >= k {
				s += arr[j]
			}
			if arr[j] == 0 && s >= 1 {
				s--
				cnt++
			}
		}
		fmt.Println(cnt)
	}
}

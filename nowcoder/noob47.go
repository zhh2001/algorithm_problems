package main

import (
	"fmt"
)

func main() {
	var l, m int
	_, err := fmt.Scan(&l, &m)
	if err != nil {
		panic(err)
	}

	arr := make([]bool, l+1)
	for i := 0; i < m; i++ {
		var l, r int
		_, err = fmt.Scan(&l, &r)
		if err != nil {
			panic(err)
		}

		for j := l; j <= r; j++ {
			arr[j] = true
		}
	}

	cnt := 0
	for _, v := range arr {
		if v {
			cnt++
		}
	}

	fmt.Println(l - cnt + 1)
}

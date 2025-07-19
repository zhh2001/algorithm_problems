package main

import (
	"fmt"
)

func main() {
	var n int
	var err error
	_, err = fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}

	var arr [21]int
	arr[1] = 0
	arr[2] = 1
	arr[3] = 1
	for i := 4; i <= n; i++ {
		arr[i] = arr[i-3] + arr[i-2]*2 + arr[i-1]
	}
	fmt.Println(arr[n])
}

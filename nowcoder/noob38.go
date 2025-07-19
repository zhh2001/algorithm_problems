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

	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			continue
		}
		m := i
		for m != 0 {
			if m%10 == 4 {
				break
			}
			m /= 10
		}
		if m != 0 {
			continue
		}
		fmt.Println(i)
	}
}

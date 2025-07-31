package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	_, err := fmt.Scanf("%s\n", &n)
	if err != nil {
		panic(err)
	}

	var a, b, c int
	a, err = strconv.Atoi(n[0:1])
	if err != nil {
		panic(err)
	}
	b, err = strconv.Atoi(n[1:3])
	if err != nil {
		panic(err)
	}
	c = len(n) - 1

	if b < 10 {
		if b >= 5 {
			b = 1
		} else {
			b = 0
		}
	} else {
		if b%10 >= 5 {
			b /= 10
			b++
			if b == 10 {
				b = 0
				a++
				if a == 10 {
					a = 1
					c++
				}
			}
		} else {
			b /= 10
		}
	}
	fmt.Printf("%d.%d*10^%d\n", a, b, c)
}

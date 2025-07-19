package main

import (
	"fmt"
)

func IsPrimeNumber(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var t, n int
	var err error
	_, err = fmt.Scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		_, err = fmt.Scanf("%d\n", &n)
		if err != nil {
			panic(err)
		}
		if IsPrimeNumber(n) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

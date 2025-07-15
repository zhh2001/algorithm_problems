package main

import (
	"fmt"
)

func main() {
	var f float64
	for {
		n, _ := fmt.Scan(&f)
		if n == 0 {
			break
		} else {
			fmt.Printf("%.3f\n", f)
		}
	}
}

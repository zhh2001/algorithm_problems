package main

import (
	"fmt"
)

func main() {
	var c, d float64
	for {
		n, _ := fmt.Scan(&c, &d)
		if n == 0 {
			break
		} else {
			fmt.Printf("%.3f%%\n", d/c*100)
		}
	}
}

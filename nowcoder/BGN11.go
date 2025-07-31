package main

import (
	"fmt"
)

func main() {
	var s string
	_, err := fmt.Scanf("%s\n", &s)
	if err != nil {
		panic(err)
	}

	var rgb [3]int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'R':
			rgb[0]++
		case 'G':
			rgb[1]++
		case 'B':
			rgb[2]++
		default:
			panic("字符串格式错误")
		}
	}
	fmt.Printf("(%d,%d,%d)\n", rgb[0], rgb[1], rgb[2])
}

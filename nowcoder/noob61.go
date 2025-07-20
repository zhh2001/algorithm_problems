package main

import (
	"fmt"
)

func main() {
	var n, m int
	_, err := fmt.Scanf("%d %d\n", &n, &m)
	if err != nil {
		panic(err)
	}

	var s string
	_, err = fmt.Scanf("%s\n", &s)
	if err != nil {
		panic(err)
	}
	chars := []byte(s)

	for i := 0; i < m; i++ {
		var (
			l, r   int
			c1, c2 byte
		)
		_, err = fmt.Scanf("%d %d %c %c\n", &l, &r, &c1, &c2)
		if err != nil {
			panic(err)
		}

		for j := l - 1; j < r; j++ {
			if chars[j] == c1 {
				chars[j] = c2
			}
		}
	}

	s = string(chars)
	fmt.Println(s)
}

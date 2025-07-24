package main

import (
	"fmt"
)

func main() {
	var t int
	_, err := fmt.Scanf("%d\n", &t)
	if err != nil {
		panic(err)
	}

	var s string
	for i := 0; i < t; i++ {
		_, err = fmt.Scanf("%s\n", &s)
		if err != nil {
			panic(err)
		}

		chars := make([]byte, 0)
		for j := 0; j < len(s); j++ {
			if len(chars) == 0 {
				chars = append(chars, s[j])
				continue
			}
			if chars[len(chars)-1] == s[j] {
				if s[j] == 'O' {
					chars = chars[:len(chars)-1]
				} else if s[j] == 'o' {
					chars[len(chars)-1] = 'O'
					for len(chars) > 1 && (chars[len(chars)-1] == 'O' && chars[len(chars)-2] == 'O') {
						chars = chars[:len(chars)-2]
					}
				}
			} else {
				chars = append(chars, s[j])
			}
		}

		fmt.Println(string(chars))
	}
}

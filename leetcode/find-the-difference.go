package main

import (
	"fmt"
)

func findTheDifference(s string, t string) byte {
	chars := map[byte]int{}
	for i := 0; i < len(t); i++ {
		chars[t[i]]++
	}
	for i := 0; i < len(s); i++ {
		chars[s[i]]--
		if chars[s[i]] == 0 {
			delete(chars, s[i])
		}
	}

	var char byte
	for key := range chars {
		char = key
		break
	}
	return char
}

func main() {
	fmt.Println(string(findTheDifference("abcd", "abcde")))
	fmt.Println(string(findTheDifference("", "y")))
}

package main

import (
	"fmt"
)

func main() {
	var str string
	_, err := fmt.Scanf("%s\n", &str)
	if err != nil {
		panic(err)
	}

	flag := true
	stack := make([]byte, 0, len(str))
	for i := range str {
		if str[i] == 'b' {
			if len(stack) > 0 && stack[len(stack)-1] == 'a' {
				stack = stack[:len(stack)-1]
			} else {
				flag = false
				break
			}
		} else if str[i] == 'a' {
			stack = append(stack, str[i])
		} else {
			flag = false
			break
		}
	}
	if len(stack) > 0 {
		flag = false
	}
	if flag {
		fmt.Println("Good")
	} else {
		fmt.Println("Bad")
	}
}

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
		if str[i] == '[' || str[i] == '(' {
			stack = append(stack, str[i])
		} else if str[i] == ']' {
			if len(stack) == 0 {
				flag = false
				break
			} else {
				if stack[len(stack)-1] == '[' {
					stack = stack[:len(stack)-1]
				} else {
					flag = false
					break
				}
			}
		} else if str[i] == ')' {
			if len(stack) == 0 {
				flag = false
				break
			} else {
				if stack[len(stack)-1] == '(' {
					stack = stack[:len(stack)-1]
				} else {
					flag = false
					break
				}
			}
		}
	}
	if len(stack) > 0 {
		flag = false
	}
	fmt.Println(flag)
}

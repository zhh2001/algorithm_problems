package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		minLen := 0
		maxLen := 0
		options := [4]string{}
		for j := 0; j < 4; j++ {
			if scanner.Scan() {
				s := scanner.Text()
				options[j] = s
				if j == 0 {
					maxLen = len(s)
					minLen = len(s)
				} else {
					if maxLen < len(s) {
						maxLen = len(s)
					}
					if minLen > len(s) {
						minLen = len(s)
					}
				}
			}
		}

		minCnt := 0
		maxCnt := 0
		var long, short string
		for _, option := range options {
			if len(option) == minLen {
				minCnt++
				short = option
			}
			if len(option) == maxLen {
				maxCnt++
				long = option
			}
		}
		if minCnt == 1 {
			fmt.Printf("%c\n", short[0])
		} else if maxCnt == 1 && minCnt > 1 {
			fmt.Printf("%c\n", long[0])
		} else {
			fmt.Printf("%c\n", 'C')
		}
	}
}

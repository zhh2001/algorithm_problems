package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	builder := strings.Builder{}
	for scanner.Scan() {
		s := scanner.Text()
		words := strings.Split(s, " ")
		for _, word := range words {
			if len(word) > 0 {
				builder.WriteByte(strings.ToUpper(word)[0])
			}
		}
	}
	fmt.Println(builder.String())
}

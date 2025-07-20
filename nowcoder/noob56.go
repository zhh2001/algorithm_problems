package main

import (
	"fmt"
	"strings"
)

func BobFirstSearch(s string) int {
	return strings.Index(strings.ToLower(s), "bob")
}

func main() {
	var s string
	_, err := fmt.Scanf("%s\n", &s)
	if err != nil {
		panic(err)
	}

	fmt.Println(BobFirstSearch(s))
}

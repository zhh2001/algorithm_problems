package main

import (
	"fmt"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}

	mapping := map[string]int{}
	for i := 0; i < n; i++ {
		var str string
		_, err := fmt.Scanf("%s\n", &str)
		if err != nil {
			panic(err)
		}

		mapping[str]++
	}
	fmt.Println(len(mapping))
}

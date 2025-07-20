package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}

	builder := strings.Builder{}
	for i := 1; len(builder.String()) < n; i++ {
		builder.WriteString(strconv.Itoa(i))
	}

	fmt.Printf("%c\n", builder.String()[n-1])
}

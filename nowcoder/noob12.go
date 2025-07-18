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
	s := strconv.Itoa(n)
	sli := strings.Split(s, "")
	for i, j := 0, len(sli)-1; i < j; i, j = i+1, j-1 {
		sli[i], sli[j] = sli[j], sli[i]
	}
	s = strings.Join(sli, "")
	fmt.Println(s)
}

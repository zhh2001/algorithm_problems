package main

import (
	"fmt"
)

const (
	SECOND = 1
	MINUTE = 60 * SECOND
	HOUR   = 60 * MINUTE
)

func main() {
	var hours, minutes, seconds int
	_, err := fmt.Scanf("%d\n", &seconds)
	if err != nil {
		panic(err)
	}
	hours = seconds / HOUR
	seconds = seconds % HOUR
	minutes = seconds / MINUTE
	seconds = seconds % MINUTE
	fmt.Println(hours, minutes, seconds)
}

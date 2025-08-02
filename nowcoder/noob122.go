package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var n uint64
	scanner.Scan()
	n, err := strconv.ParseUint(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	mapping := make(map[uint64]uint64)
	var sum uint64
	for i := uint64(1); i <= n; i++ {
		var x, y uint64

		// 读取x
		scanner.Scan()
		x, err = strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}

		// 读取y
		scanner.Scan()
		y, err = strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}

		sum += mapping[x] * i
		mapping[x] = y
	}
	fmt.Println(sum)
}

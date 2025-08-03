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

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	wordIndex := make(map[string][]int)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		l, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		for j := 0; j < l; j++ {
			scanner.Scan()
			w := scanner.Text()
			idx, ok := wordIndex[w]
			if ok && idx[len(idx)-1] == i {
				continue
			} else {
				wordIndex[w] = append(wordIndex[w], i)
			}
		}
	}

	scanner.Scan()
	m, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	for i := 0; i < m; i++ {
		scanner.Scan()
		w := scanner.Text()
		if idx, ok := wordIndex[w]; ok {
			for _, val := range idx {
				fmt.Printf("%d ", val)
			}
		}
		fmt.Println()
	}
}

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

	cityCodeMap := make(map[string]string)
	codeCityMap := make(map[string][]string)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()

		scanner.Scan()
		c := scanner.Text()

		cityCodeMap[s] = c
		codeCityMap[s[:2]] = append(codeCityMap[s[:2]], s)
	}

	counter := make(map[string]string)
	for cityA, codeB := range cityCodeMap {
		codeA := cityA[:2]
		if codeA == codeB {
			continue
		}
		cities := codeCityMap[codeB]
		for _, cityB := range cities {
			if cityCodeMap[cityB] == codeA {
				counter[cityA] = cityB
				counter[cityB] = cityA
			}
		}
	}

	fmt.Println(len(counter) / 2)
}

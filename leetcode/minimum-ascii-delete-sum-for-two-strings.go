package main

import (
	"fmt"
)

func minimumDeleteSum(s1 string, s2 string) int {
	total := 0
	for _, c := range s1 {
		total += int(c)
	}
	for _, c := range s2 {
		total += int(c)
	}

	f := make([][]int, len(s1)+1)
	for i := range f {
		f[i] = make([]int, len(s2)+1)
	}
	for i, x := range s1 {
		for j, y := range s2 {
			if x == y {
				f[i+1][j+1] = f[i][j] + int(x)
			} else {
				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
			}
		}
	}
	return total - f[len(s1)][len(s2)]*2
}

func main() {
	res1 := minimumDeleteSum("sea", "eat")
	fmt.Println(res1)
	res2 := minimumDeleteSum("delete", "leet")
	fmt.Println(res2)
}

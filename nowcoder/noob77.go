package main

import (
	"fmt"
	"sort"
)

func main() {
	var q int
	_, err := fmt.Scanf("%d\n", &q)
	if err != nil {
		panic(err)
	}

	seq := make([]int, 0)

	for j := q; j > 0; j-- {
		var n, i, x int
		_, err = fmt.Scan(&n)
		if err != nil {
			panic(err)
		}

		switch n {
		case 1:
			_, err = fmt.Scanln(&x)
			if err != nil {
				panic(err)
			}
			seq = append(seq, x)
		case 2:
			seq = seq[:len(seq)-1]
		case 3:
			_, err = fmt.Scan(&i)
			if err != nil {
				panic(err)
			}
			fmt.Println(seq[i])
		case 4:
			_, err = fmt.Scan(&i, &x)
			if err != nil {
				panic(err)
			}
			seq = append(seq, 0)
			copy(seq[i+2:], seq[i+1:])
			seq[i+1] = x
		case 5:
			sort.Ints(seq)
		case 6:
			sort.Sort(sort.Reverse(sort.IntSlice(seq)))
		case 7:
			fmt.Println(len(seq))
		case 8:
			for i, v := range seq {
				if i == len(seq)-1 {
					fmt.Println(v)
					break
				}
				fmt.Printf("%d ", v)
			}
		}
	}
}

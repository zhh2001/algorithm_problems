package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}

	var m int
	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&m)
		if err != nil {
			panic(err)
		}
		arr := make([]int, m)
		sum := 0
		for j := 0; j < m; j++ {
			_, err = fmt.Scan(&arr[j])
			if err != nil {
				panic(err)
			}
			sum += arr[j]
		}

		var a, b, r int
		a = Max[int](arr[0], arr[1:]...)
		b = Min[int](arr[0], arr[1:]...)
		r = a - b

		var avg, d float64
		avg = float64(sum) / float64(m)
		for _, v := range arr {
			d += math.Pow(float64(v)-avg, 2)
		}
		d /= float64(m)

		fmt.Printf("%d %.3f\n", r, d)
	}
}

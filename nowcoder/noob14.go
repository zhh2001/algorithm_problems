package main

import (
	"fmt"
	"math"
)

func calcEuclideanDistance(x1, y1, x2, y2 float64) float64 {
	t1 := math.Pow(x1-x2, 2)
	t2 := math.Pow(y1-y2, 2)
	t := t1 + t2
	return math.Sqrt(t)
}

func calcManhattanDistance(x1, y1, x2, y2 float64) float64 {
	t1 := math.Abs(x1 - x2)
	t2 := math.Abs(y1 - y2)
	t := t1 + t2
	return t
}

func main() {
	var x1, y1, x2, y2 float64
	_, err := fmt.Scanf("%f %f\n", &x1, &y1)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Scanf("%f %f\n", &x2, &y2)
	if err != nil {
		panic(err)
	}
	e := calcEuclideanDistance(x1, y1, x2, y2)
	m := calcManhattanDistance(x1, y1, x2, y2)
	fmt.Println(math.Abs(m - e))
}

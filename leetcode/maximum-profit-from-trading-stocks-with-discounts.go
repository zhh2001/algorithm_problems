package main

import "fmt"

func maxProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	g := make([][]int, n)
	for _, e := range hierarchy {
		x, y := e[0]-1, e[1]-1
		g[x] = append(g[x], y)
	}

	var dfs func(int) [][2]int
	dfs = func(x int) [][2]int {
		subF := make([][2]int, budget+1)
		for _, y := range g[x] {
			fy := dfs(y)
			for j := budget; j >= 0; j-- {
				for jy, p := range fy[:j+1] {
					for k, resY := range p {
						subF[j][k] = max(subF[j][k], subF[j-jy][k]+resY)
					}
				}
			}
		}
		f := make([][2]int, budget+1)
		for j, p := range subF {
			for k := 0; k < 2; k++ {
				cost := present[x] / (k + 1)
				if j >= cost {
					f[j][k] = max(p[0], subF[j-cost][1]+future[x]-cost)
				} else {
					f[j][k] = p[0]
				}
			}
		}
		return f
	}

	return dfs(0)[budget][0]
}

func main() {
	res1 := maxProfit(2, []int{1, 2}, []int{4, 3}, [][]int{{1, 2}}, 3)
	fmt.Println(res1)
	res2 := maxProfit(2, []int{3, 4}, []int{5, 8}, [][]int{{1, 2}}, 4)
	fmt.Println(res2)
	res3 := maxProfit(3, []int{4, 6, 8}, []int{7, 9, 11}, [][]int{{1, 2}, {1, 3}}, 10)
	fmt.Println(res3)
	res4 := maxProfit(3, []int{5, 2, 3}, []int{8, 5, 6}, [][]int{{1, 2}, {2, 3}}, 7)
	fmt.Println(res4)
}

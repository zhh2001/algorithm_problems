package main

import "fmt"

type Pair struct {
	X int
	Y int
}

var dirs = [4]Pair{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func robotSim(commands []int, obstacles [][]int) int {
	isObstacle := make(map[Pair]bool, len(obstacles))
	for _, obstacle := range obstacles {
		isObstacle[Pair{obstacle[0], obstacle[1]}] = true
	}

	res := 0
	x, y, k := 0, 0, 0
	for _, command := range commands {
		switch command {
		case -1:
			k = (k + 1) % 4
		case -2:
			k = (k + 3) % 4
		default:
			for i := command; i > 0; i-- {
				if isObstacle[Pair{x + dirs[k].X, y + dirs[k].Y}] {
					break
				}
				x += dirs[k].X
				y += dirs[k].Y
			}
			res = max(res, x*x+y*y)
		}
	}
	return res
}

func main() {
	commands := []int{4, -1, 3}
	obstacles := [][]int{}
	res := robotSim(commands, obstacles)
	fmt.Println(res)

	commands = []int{4, -1, 4, -2, 4}
	obstacles = [][]int{{2, 4}}
	res = robotSim(commands, obstacles)
	fmt.Println(res)

	commands = []int{6, -1, -1, 6}
	obstacles = [][]int{{0, 0}}
	res = robotSim(commands, obstacles)
	fmt.Println(res)
}

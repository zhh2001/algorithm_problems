package main

import (
	"fmt"
)

type Noob69 struct {
	Name    string
	Chinese int
	Math    int
	English int
}

func NewNoob69(name string, chinese int, math int, english int) *Noob69 {
	return &Noob69{
		Name:    name,
		Chinese: chinese,
		Math:    math,
		English: english,
	}
}

func (stu *Noob69) TotalScore() int {
	return stu.Chinese + stu.Math + stu.English
}

func (stu *Noob69) String() string {
	return fmt.Sprintf(
		"%s %d %d %d",
		stu.Name,
		stu.Chinese,
		stu.Math,
		stu.English,
	)
}

func main() {
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}

	maxScore := 0
	var curStu *Noob69
	for i := 0; i < n; i++ {
		var s string
		var c1, c2, c3 int
		_, err = fmt.Scanf(
			"%s %d %d %d\n",
			&s, &c1, &c2, &c3,
		)
		if err != nil {
			panic(err)
		}

		stu := NewNoob69(s, c1, c2, c3)
		totalScore := stu.TotalScore()
		if totalScore > maxScore {
			maxScore = totalScore
			curStu = stu
		}
	}

	fmt.Println(curStu)
}

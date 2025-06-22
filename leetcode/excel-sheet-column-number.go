package main

func titleToNumber(columnTitle string) int {
	var number int
	for i, multiple := len(columnTitle)-1, 1; i >= 0; i-- {
		k := columnTitle[i] - 'A' + 1
		number = number + int(k)*multiple
		multiple = multiple * 26
	}
	return number
}

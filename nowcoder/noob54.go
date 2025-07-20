package main

import (
	"fmt"
)

func main() {
	var s1, s2 string
	_, err := fmt.Scanf("%s %s\n", &s1, &s2)
	if err != nil {
		panic(err)
	}

	rules := map[string]map[string]string{
		"elephant": {
			"tiger": "win",
			"mouse": "lose",
		},
		"tiger": {
			"cat":      "win",
			"elephant": "lose",
		},
		"cat": {
			"mouse": "win",
			"tiger": "lose",
		},
		"mouse": {
			"elephant": "win",
			"cat":      "lose",
		},
	}

	res, ok := rules[s1][s2]
	if !ok {
		res = "tie"
	}
	fmt.Println(res)
}

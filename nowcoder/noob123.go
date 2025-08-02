package main

/**
 *
 * @param s string字符串
 * @param c string字符串
 * @return int整型
 */
func isCongruent(s string, c string) int {
	sChar := make(map[byte]int)
	cChar := make(map[byte]int)

	for _, char := range []byte(s) {
		sChar[char]++
	}
	for _, char := range []byte(c) {
		cChar[char]++
	}

	if len(sChar) != len(cChar) {
		return -1
	}
	for key, value := range sChar {
		if cChar[key] != value {
			return -1
		}
	}

	return len(s)
}

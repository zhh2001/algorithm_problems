package main

import "fmt"

/**
 *
 * @param ransomNote string字符串
 * @param magazine string字符串
 * @return bool布尔型
 */
func canConstruct(ransomNote string, magazine string) bool {
	charCount := make(map[byte]int)
	for _, char := range []byte(magazine) {
		charCount[char]++
	}
	for _, char := range []byte(ransomNote) {
		charCount[char]--
		if charCount[char] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct("wc", "new bee, wo cao"))
}

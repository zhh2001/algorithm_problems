package main

func mergeAlternately(word1 string, word2 string) string {
	bytes := make([]byte, 0, len(word1)+len(word2))
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			bytes = append(bytes, word1[i])
		}
		if i < len(word2) {
			bytes = append(bytes, word2[i])
		}
	}
	return string(bytes)
}

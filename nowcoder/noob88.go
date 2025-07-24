package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串
 * @return bool布尔型
 */
func isValid(s string) bool {
	chars := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if len(chars) == 0 {
			chars = append(chars, s[i])
			continue
		}

		if s[i] == ']' && chars[len(chars)-1] == '[' {
			chars = chars[:len(chars)-1]
		} else if s[i] == ')' && chars[len(chars)-1] == '(' {
			chars = chars[:len(chars)-1]
		} else if s[i] == '}' && chars[len(chars)-1] == '{' {
			chars = chars[:len(chars)-1]
		} else {
			chars = append(chars, s[i])
		}
	}
	return len(chars) == 0
}

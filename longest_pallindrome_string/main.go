package main

import (
	"fmt"
)

func main() {
	s := "babad"
	fmt.Println("Subtring is", longestPalindrome(s))
}
func longestPalindrome(s string) string {
	start := 0
	end := 0

	maxLength := 0
	left := -1
	right := -1
	tempLength := maxLength
	length_str := len(s)
	for i, _ := range s {
		left, right = expandAroundCenter(s, i, i, length_str)
		tempLength = right - left + 1
		if tempLength > maxLength {

			maxLength = tempLength
			start = left
			end = right
		}

		if i == length_str-1 {
			break
		}

		left, right = expandAroundCenter(s, i, i+1, length_str)
		tempLength = right - left + 1
		if tempLength > maxLength && s[left] == s[right] {

			maxLength = tempLength
			start = left
			end = right
		}

	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int, length_str int) (int, int) {
	L := left
	R := right
	isLoop := false
	for L >= 0 && R < length_str && s[L] == s[R] {
		L--
		R++
		isLoop = true
	}
	if isLoop {
		L++
		R--
	}
	return L, R
}

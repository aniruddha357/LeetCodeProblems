package main

import "fmt"

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	global_length := 0

	char_map := map[byte]int{}

	left := 0
	right := 0

	for i := 0; i < len(s); i++ {
		character := s[i]

		if char_map[character] > left {
			left = char_map[character]
		}

		if right-left+1 > global_length {
			global_length = right - left + 1
		}

		char_map[character] = i + 1
		right += 1
	}

	return global_length
}

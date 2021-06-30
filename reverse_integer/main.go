package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(1534236469))
}
func reverse(x int) int {
	var rev int = 0
	var max int = 2147483647
	var min int = -2147483648
	for x != 0 {
		pop := x % 10

		rev *= 10
		rev += pop

		x /= 10
	}

	if rev > max || rev < min {
		return 0
	}
	return rev
}

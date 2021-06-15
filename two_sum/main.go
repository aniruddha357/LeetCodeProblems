package main

import "fmt"

func twoSum(nums []int, target int) []int {
	output := make([]int, 2)
	diffMap := make(map[int]int, len(nums))

	for i, val := range nums {
		diffMap[val] = i
	}
	for i, val := range nums {
		expectedVal := target - val
		if diffMap[expectedVal] != 0 && diffMap[expectedVal] != i {
			output[0] = i
			output[1] = diffMap[expectedVal]
			break
		}
	}
	if output[0] > output[1] {
		output[0], output[1] = output[1], output[0]
	}

	return output

}

func main() {
	nums := []int{1, 3, 4, 2}
	target := 6
	output := twoSum(nums, target)
	fmt.Println("Output is ", output)
}

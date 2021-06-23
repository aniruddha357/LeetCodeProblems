package main

import (
	"fmt"
	"math"
)

func main() {
	num1 := []int{2}
	num2 := []int{}
	fmt.Println(findMedianSortedArrays(num1, num2))

}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	middlePosition := float64((len1 + len2)) / 2.0

	if (len1 + len2) < 2 {
		middlePosition = 1
	}

	isEven := ((len1 + len2) % 2) == 0.0

	middle1 := 0
	middle2 := 0

	count1 := 0
	count2 := 0

	mP := int(math.Round(middlePosition)) - 1

	for i := 0; i <= mP || isEven; i++ {
		if count1 < len1 && (count2 >= len2 || nums1[count1] <= nums2[count2]) && i <= mP {
			middle1 = nums1[count1]
			count1++
		} else if count2 < len2 && i <= mP {
			middle1 = nums2[count2]
			count2++
		} else if count1 < len1 && (count2 >= len2 || nums1[count1] <= nums2[count2]) {
			middle2 = nums1[count1]
			break
		} else if count2 < len2 {
			middle2 = nums2[count2]
			break
		}
	}

	if !isEven {
		return float64(middle1)
	}

	return float64((middle1 + middle2)) / 2.0

}

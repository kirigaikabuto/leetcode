package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	find := 0
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] != nums[1] {
		return nums[0]
	}
	n := len(nums)
	if nums[n-1] != nums[n-2] {
		return nums[n-1]
	}
	for i := 1; i < n-1; i++ {
		if nums[i-1] != nums[i] && nums[i] != nums[i+1]{
			find = nums[i]
		}
	}
	return find
}

func main() {
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
}

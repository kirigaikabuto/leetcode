package main

func moveZeroes(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums = append(nums, 0)
			nums = append(nums[:i], nums[i+1:]...)
			i -= 1
		}
	}
}

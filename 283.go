package main

import "fmt"

func moveZeroes(nums []int) {
	n := len(nums)
	i := 0
	j := 0
	for true {
		if j == n-1 {
			break
		}
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
		} else {
			i += 1
		}
		j += 1
	}
	fmt.Println(nums)
}

func main() {
	test := []int{0,1,0,3,12}
	moveZeroes(test)
}
